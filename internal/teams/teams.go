package teams

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/1-platform/api-catalog/internal/pkg/config"
	"github.com/1-platform/api-catalog/internal/pkg/smtp"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrTeamSlugAlreadyExist = errors.New("Team slug already exist")
	ErrTeamNotFound         = errors.New("Team not found")
	ErrLastTeamMember       = errors.New("Cannot remove the last admin member")
	ErrCannotRemoveYourself = errors.New("Cannot remove yourself from team")
	ErrAlreadyTeamMember    = errors.New("User is already a member")
	ErrInvalidInviteeToken  = errors.New("Invalid invitee token provided")
)

type Teams struct {
	store        Store
	mainCfg      *config.Config
	emailService smtp.EmailService
}

type Store interface {
	GetATeamById(ctx context.Context, id primitive.ObjectID) (*Team, error)
	GetATeamBySlug(ctx context.Context, slug string) (*Team, error)
	InsertTeam(ctx context.Context, team *Team) (id primitive.ObjectID, err error)
	UpdateTeamById(ctx context.Context, userId primitive.ObjectID, teamId primitive.ObjectID, team *UpdateTeamDTO) (updatedTeam *Team, err error)
	DeleteTeamById(ctx context.Context, teamId primitive.ObjectID) (team *Team, err error)
	ListTeams(ctx context.Context, dto *ListTeamDTO) (team *Pagination[[]Team], err error)
	// membership store function
	InsertTeamMembership(ctx context.Context, membership *TeamMember) (id primitive.ObjectID, err error)
	RemoveTeamMembership(ctx context.Context, teamId primitive.ObjectID, membershipId primitive.ObjectID) error
	GetTeamMemberCount(ctx context.Context, teamId primitive.ObjectID) (int64, error)
	ListTeamMembers(ctx context.Context, teamId primitive.ObjectID, dto *ListTeamMembersDTO) (teamMembers *Pagination[[]TeamMemberDetailed], err error)
	GetTeamMembershipById(ctx context.Context, membershipId primitive.ObjectID) (membership *TeamMember, err error)
	GetTeamMembershipByEmail(ctx context.Context, teamId primitive.ObjectID, email string) (membership *TeamMemberDetailed, err error)
	UpdateTeamMembershipById(ctx context.Context, membershipId primitive.ObjectID, dto *UpdateTeamMembershipDTO) (err error)
	// to get membership of a user in a team, mainly for middleware
	GetTeamMembershipByUserId(ctx context.Context, teamId primitive.ObjectID, userId primitive.ObjectID) (membership *TeamMember, err error)
}

func New(store Store, emailService smtp.EmailService, mainCfg *config.Config) *Teams {
	return &Teams{store: store, mainCfg: mainCfg, emailService: emailService}
}

func (tms *Teams) CreateTeam(ctx context.Context, userId string, tm *Team) (*Team, error) {
	savedTm, _ := tms.store.GetATeamBySlug(ctx, tm.Slug)
	if savedTm != nil {
		return nil, ErrTeamSlugAlreadyExist
	}

	membershipMongoUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	tm.Id = primitive.NewObjectID()
	tm.CreatedBy = membershipMongoUserId
	tm.UpdatedBy = membershipMongoUserId
	tm.CreatedAt = time.Now()
	tm.UpdatedAt = time.Now()

	newTeamId, err := tms.store.InsertTeam(ctx, tm)
	if err != nil {
		return nil, err
	}

	membership := TeamMember{
		Id:     primitive.NewObjectID(),
		UserId: membershipMongoUserId,
		Role:   TeamAdminRole,
		TeamId: newTeamId,
		Status: MembershipAcceptedStatus,
	}
	_, err = tms.store.InsertTeamMembership(ctx, &membership)
	if err != nil {
		return nil, err
	}

	newTeam, err := tms.store.GetATeamById(ctx, newTeamId)
	if err != nil {
		return nil, err
	}
	return newTeam, nil
}

func (tms *Teams) UpdateTeam(ctx context.Context, userId string, teamId string, tmDTO *UpdateTeamDTO) (*Team, error) {
	teamMongoId, err := primitive.ObjectIDFromHex(teamId)
	if err != nil {
		return nil, err
	}

	membershipMongoUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	if tmDTO.Slug != nil {
		tmBySlug, err := tms.store.GetATeamBySlug(ctx, *tmDTO.Slug)
		if err == nil && tmBySlug.Id != teamMongoId {
			return nil, ErrTeamSlugAlreadyExist
		}
		if err != nil && err != mongo.ErrNoDocuments {
			return nil, err
		}
	}

	updatedTeam, err := tms.store.UpdateTeamById(ctx, membershipMongoUserId, teamMongoId, tmDTO)
	if err != nil {
		return nil, err
	}

	return updatedTeam, nil
}

func (tms *Teams) DeleteTeam(ctx context.Context, teamId string) (*Team, error) {
	teamMongoId, err := primitive.ObjectIDFromHex(teamId)
	if err != nil {
		return nil, err
	}

	deletedTeam, err := tms.store.DeleteTeamById(ctx, teamMongoId)
	if err != nil {
		return nil, err
	}

	return deletedTeam, nil
}

func (tms *Teams) ListTeams(ctx context.Context, listTmDto *ListTeamDTO) (*Pagination[[]Team], error) {
	teams, err := tms.store.ListTeams(ctx, listTmDto)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (tms *Teams) GetTeamById(ctx context.Context, teamId string) (*Team, error) {
	teamMongoId, err := primitive.ObjectIDFromHex(teamId)
	if err != nil {
		return nil, err
	}

	tm, err := tms.store.GetATeamById(ctx, teamMongoId)
	if err != nil {
		return nil, err
	}

	return tm, nil
}

func (tms *Teams) ListTeamMembers(ctx context.Context, teamId string, listTmMsDTO *ListTeamMembersDTO) (*Pagination[[]TeamMemberDetailed], error) {
	teamMongoId, err := primitive.ObjectIDFromHex(teamId)
	if err != nil {
		return nil, err
	}

	tm, err := tms.store.ListTeamMembers(ctx, teamMongoId, listTmMsDTO)
	if err != nil {
		return nil, err
	}

	return tm, nil
}

func (tms *Teams) InviteTeamMember(ctx context.Context, teamId string, inviteeEmail string, role TeamMembershipRole) error {
	teamMongoId, err := primitive.ObjectIDFromHex(teamId)
	if err != nil {
		return err
	}

	existingMember, err := tms.store.GetTeamMembershipByEmail(ctx, teamMongoId, inviteeEmail)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if err == nil {
		if existingMember.Status == MembershipAcceptedStatus {
			return ErrAlreadyTeamMember
		}

		// logic to resent invite url
		inviteeUrl, err := tms.getInviteeTokenURL(teamId, existingMember.Id.Hex())
		if err != nil {
			return err
		}
		body := getInviteeEmailTemplate("Team X", inviteeUrl)
		emailData := tms.emailService.GetMailData("Invited to join the team", body)
		if err = tms.emailService.SendMail([]string{inviteeEmail}, emailData); err != nil {
			return err
		}
		return nil
	}

	membership := TeamMember{
		Id:     primitive.NewObjectID(),
		Email:  &inviteeEmail,
		TeamId: teamMongoId,
		Status: MembershipInvitedStatus,
		Role:   role,
	}
	_, err = tms.store.InsertTeamMembership(ctx, &membership)
	if err != nil {
		return err
	}

	inviteeUrl, err := tms.getInviteeTokenURL(teamId, membership.Id.Hex())
	if err != nil {
		return err
	}
	body := getInviteeEmailTemplate("Team X", inviteeUrl)
	emailData := tms.emailService.GetMailData("Invited to join the team", body)
	if err = tms.emailService.SendMail([]string{inviteeEmail}, emailData); err != nil {
		return err
	}
	return nil
}

func (tms *Teams) JoinTeamUsingToken(ctx context.Context, userId string, userEmail string, teamId string, tokenStr string) error {
	userMongoId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	token, err := jwt.ParseWithClaims(tokenStr, &MembershipInviteJwt{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(tms.mainCfg.Jwt.SigningKey), nil
	})
	if !token.Valid {
		return ErrInvalidInviteeToken
	}
	claim, ok := token.Claims.(*MembershipInviteJwt)
	if !ok {
		return ErrInvalidInviteeToken
	}
	membershipMongoId, err := primitive.ObjectIDFromHex(claim.MembershipId)
	if err != nil {
		return err
	}

	teamMongoId, err := primitive.ObjectIDFromHex(claim.TeamId)
	if err != nil {
		return err
	}

	membership, err := tms.store.GetTeamMembershipById(ctx, membershipMongoId)
	if err != nil {
		return err
	}
	if membership.Status == MembershipAcceptedStatus {
		return ErrInvalidInviteeToken
	}
	if userEmail != *membership.Email || teamId != claim.TeamId || membership.TeamId != teamMongoId {
		return ErrInvalidInviteeToken
	}
	status := MembershipAcceptedStatus
	email := ""
	updateMembershipDto := &UpdateTeamMembershipDTO{
		Email:  &email,
		Status: &status,
		UserId: &userMongoId,
	}
	err = tms.store.UpdateTeamMembershipById(ctx, membership.Id, updateMembershipDto)
	if err != nil {
		return err
	}
	return nil
}

func (tms *Teams) RemoveTeamMember(ctx context.Context, teamId string, membershipId string) error {
	teamMongoId, err := primitive.ObjectIDFromHex(teamId)
	if err != nil {
		return err
	}

	membershipMongoId, err := primitive.ObjectIDFromHex(membershipId)
	if err != nil {
		return err
	}

	membership, err := tms.store.GetTeamMembershipById(ctx, membershipMongoId)
	if err != nil {
		return err
	}
	if membership.Id == membershipMongoId {
		return ErrCannotRemoveYourself
	}

	if err = tms.store.RemoveTeamMembership(ctx, teamMongoId, membershipMongoId); err != nil {
		return err
	}
	return nil
}

func (tms *Teams) GetUserTeamMembership(ctx context.Context, teamId, userId string) (*TeamMember, error) {
	membershipMongoUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	teamMongoId, err := primitive.ObjectIDFromHex(teamId)
	if err != nil {
		return nil, err
	}

	tmMembership, err := tms.store.GetTeamMembershipByUserId(ctx, teamMongoId, membershipMongoUserId)
	if err != nil {
		return nil, err
	}

	return tmMembership, nil
}

func (tms *Teams) getInviteeTokenURL(teamId string, membershipId string) (string, error) {
	// generate the unique url
	inviteeTkClaim := MembershipInviteJwt{
		TeamId:       teamId,
		MembershipId: membershipId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tms.mainCfg.TeamMembership.InviteeExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    tms.mainCfg.Jwt.IssuerName,
			Subject:   "membership invite",
			ID:        membershipId,
			Audience:  []string{tms.mainCfg.ApplicationURL},
		},
	}
	inviteeToken := jwt.NewWithClaims(jwt.SigningMethodHS256, inviteeTkClaim)
	signedToken, err := inviteeToken.SignedString([]byte(tms.mainCfg.Jwt.SigningKey))
	if err != nil {
		return "", err
	}
	inviteeUrl, err := url.Parse(tms.mainCfg.ApplicationURL)
	if err != nil {
		return "", err
	}
	query := inviteeUrl.Query()
	query.Set("invitee_token", signedToken)
	inviteeUrl.RawQuery = query.Encode()
	inviteeUrl.Path = "/user/team-join"
	return inviteeUrl.String(), nil
}
