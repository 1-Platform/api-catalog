package teams

import (
	"fmt"
	"time"

	"github.com/1-platform/api-catalog/internal/auth"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name"`
	Desc      string             `json:"description" bson:"description"`
	Slug      string             `json:"slug"`
	CreatedBy primitive.ObjectID `bson:"created_by" json:"created_by"`
	UpdatedBy primitive.ObjectID `bson:"updated_by" json:"updated_by"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// Membership is not embedded to separate the membership
// Checking admin etc can be kept it in this collection
type TeamMember struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Email  *string            `json:"-" bson:"invitee_email,omitempty"`
	TeamId primitive.ObjectID `json:"team_id" bson:"team_id"`
	UserId primitive.ObjectID `json:"user_id" bson:"user_id"`
	Role   TeamMembershipRole `json:"role" bson:"role"`
	Status MembershipStatus   `json:"status" bson:"status"`
}

// for showing to users etc
type TeamMemberDetailed struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	Email  string             `json:"-" bson:"invitee_email,omitempty"`
	Role   TeamMembershipRole `json:"role" bson:"role"`
	User   auth.User          `json:"user" bson:"user"`
	Team   *Team              `json:"team,omitempty" bson:"team,omitempty"`
	Status MembershipStatus   `json:"status" bson:"status"`
}

type TeamMembershipRole string

const (
	TeamAdminRole  = "admin"
	TeamMemberRole = "member"
	TeamViewerRole = "viewer"
)

func (tmr TeamMembershipRole) IsValidRole() bool {
	return tmr == TeamAdminRole || tmr == TeamMemberRole || tmr == TeamViewerRole
}

func (tmr TeamMembershipRole) IsAdmin() bool {
	return tmr == TeamAdminRole
}

type MembershipStatus string

const (
	MembershipInvitedStatus  = "invited"
	MembershipAcceptedStatus = "accepted"
)

type MembershipInviteJwt struct {
	TeamId       string `json:"team_id"`
	MembershipId string `json:"membership_id"`
	jwt.RegisteredClaims
}

type Pagination[T any] struct {
	Total   int64 `json:"count"`
	Records T     `json:"records"`
}

type AggregatedPagination[T any] struct {
	Total []struct {
		CountDocuments int64 `bson:"count"`
	}
	Records T `json:"records"`
}

type AggregratePagination[T any] []struct {
	Total   []int64 `json:"count"`
	Records T       `json:"records"`
}

// DTO
type UpdateTeamDTO struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,max=150"`
	Slug        *string `json:"slug,omitempty" validate:"omitempty,is-pid"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
}

type ListTeamDTO struct {
	Limit   int64  `query:"limit" validate:"omitempty,number"`
	Offset  int64  `query:"offset" validate:"omitempty,number"`
	Sort    string `query:"sort_by" validate:"omitempty,oneof=name created_at updated_at"`
	SortDir string `query:"sort_dir" validate:"omitempty,oneof=asc desc"`
}

type ListTeamMembersDTO struct {
	Limit   int64  `query:"limit" validate:"omitempty,number"`
	Offset  int64  `query:"offset" validate:"omitempty,number"`
	Sort    string `query:"sort_by" validate:"omitempty,oneof=name"`
	SortDir string `query:"sort_dir" validate:"omitempty,oneof=asc desc"`
}

type UpdateTeamMembershipDTO struct {
	Role   *string             `json:"role,omitempty" validate:"omitempty,oneof=admin member viewer"`
	UserId *primitive.ObjectID `json:"-"`
	Email  *string             `json:"-"`
	Status *string             `json:"-"`
}

// email templates
func getInviteeEmailTemplate(teamName, tokenURL string) string {
	return fmt.Sprintf(`
Hey there,
    You have been invited to join %s team.

	<a href="%s">
		Click here to join
	</a>
		`, teamName, tokenURL)
}
