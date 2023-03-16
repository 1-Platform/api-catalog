package api

import (
	"context"

	"github.com/1-platform/api-catalog/internal/teams"
)

func (api *API) CreateTeam(ctx context.Context, userId string, tm *teams.Team) (*teams.Team, error) {
	t, err := api.teams.CreateTeam(ctx, userId, tm)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (api *API) UpdateTeam(ctx context.Context, userId string, teamId string, tm *teams.UpdateTeamDTO) (*teams.Team, error) {
	t, err := api.teams.UpdateTeam(ctx, userId, teamId, tm)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (api *API) DeleteTeam(ctx context.Context, teamId string) (*teams.Team, error) {
	t, err := api.teams.DeleteTeam(ctx, teamId)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (api *API) GetTeamById(ctx context.Context, teamId string) (*teams.Team, error) {
	t, err := api.teams.GetTeamById(ctx, teamId)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (api *API) ListTeams(ctx context.Context, dto *teams.ListTeamDTO) (*teams.Pagination[[]teams.Team], error) {
	t, err := api.teams.ListTeams(ctx, dto)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (api *API) ListTeamMembers(ctx context.Context, teamId string, dto *teams.ListTeamMembersDTO) (*teams.Pagination[[]teams.TeamMemberDetailed], error) {
	t, err := api.teams.ListTeamMembers(ctx, teamId, dto)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (api *API) InviteTeamMember(ctx context.Context, teamId string, inviteeEmail string, role teams.TeamMembershipRole) error {
	err := api.teams.InviteTeamMember(ctx, teamId, inviteeEmail, role)
	if err != nil {
		return err
	}
	return nil
}

func (api *API) JoinTeamJoinTeamUsingToken(ctx context.Context, userId string,
	userEmail string, teamId string, tokenStr string) error {
	err := api.teams.JoinTeamUsingToken(ctx, userId, userEmail, teamId, tokenStr)
	if err != nil {
		return err
	}
	return nil
}

func (api *API) RemoveTeamMembership(ctx context.Context, teamId string, membershipId string) error {
	if err := api.teams.RemoveTeamMember(ctx, teamId, membershipId); err != nil {
		return err
	}

	return nil
}

func (api *API) GetTeamMembershipByUserId(ctx context.Context, teamId string, userId string) (*teams.TeamMember, error) {
	tm, err := api.teams.GetUserTeamMembership(ctx, teamId, userId)
	if err != nil {
		return nil, err
	}

	return tm, nil
}
