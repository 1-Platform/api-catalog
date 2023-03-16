package dto

import "github.com/1-platform/api-catalog/internal/teams"

type CreateTeam struct {
	Name        string `json:"name" validate:"required,max=150"`
	Slug        string `json:"slug" validate:"required,is-pid"`
	Description string `json:"description" validate:"max=500"`
}

type UpdateTeam struct {
	Id string `param:"teamId" validate:"required"`
	*teams.UpdateTeamDTO
}

type DeleteTeam struct {
	Id string `param:"teamId" validate:"required"`
}

type GetTeamById struct {
	TeamId string `param:"teamId" validate:"required"`
}

type ListTeam struct {
	teams.ListTeamDTO
}

type ListTeamMembers struct {
	Id string `param:"teamId" validate:"required"`
	teams.ListTeamMembersDTO
}

type InviteTeamMember struct {
	TeamId       string `param:"teamId" validate:"required"`
	InviteeEmail string `json:"invitee_email" validate:"required,email"`
	InviteeRole  string `json:"role" validate:"required,oneof=admin member viewer"`
}

type RemoveTeamMember struct {
	TeamId       string `param:"teamId" validate:"required"`
	MembershipId string `param:"membershipId" validate:"required"`
}

type JoinTeam struct {
	TeamId string `param:"teamId" validate:"required"`
	Token  string `json:"token" validate:"required"`
}
