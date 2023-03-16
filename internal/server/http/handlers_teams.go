package http

import (
	"context"
	"net/http"
	"time"

	"github.com/1-platform/api-catalog/internal/auth"
	"github.com/1-platform/api-catalog/internal/server/http/contexthandler"
	"github.com/1-platform/api-catalog/internal/server/http/dto"
	"github.com/1-platform/api-catalog/internal/server/http/response"
	"github.com/1-platform/api-catalog/internal/teams"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) CreateTeam(c echo.Context) (err error) {
	ctDto := new(dto.CreateTeam)
	if err = c.Bind(ctDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(ctDto); err != nil {
		return err
	}

	userId := c.Get(contexthandler.USER_ID).(string)
	newTeam := teams.Team{Name: ctDto.Name, Desc: ctDto.Description, Slug: ctDto.Slug}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	team, err := h.api.CreateTeam(ctx, userId, &newTeam)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrInvalidRequest(err))
	}

	return c.JSON(http.StatusOK, response.Success(team, "successfully created team"))
}

func (h *Handlers) UpdateTeam(c echo.Context) (err error) {
	utDto := new(dto.UpdateTeam)
	if err = c.Bind(utDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(utDto); err != nil {
		return err
	}
	userId := c.Get(contexthandler.USER_ID).(string)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	team, err := h.api.UpdateTeam(ctx, userId, utDto.Id, utDto.UpdateTeamDTO)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrInvalidRequest(err))
	}

	return c.JSON(http.StatusOK, response.Success(team, "successfully updated team"))
}

func (h *Handlers) DeleteTeam(c echo.Context) (err error) {
	dtDto := new(dto.DeleteTeam)
	if err = c.Bind(dtDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(dtDto); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	team, err := h.api.DeleteTeam(ctx, dtDto.Id)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrInvalidRequest(err))
	}

	return c.JSON(http.StatusOK, response.Success(team, "successfully deleted team"))
}

func (h *Handlers) GetTeamById(c echo.Context) (err error) {
	opt := &dto.GetTeamById{}
	if err = c.Bind(opt); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(opt); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	team, err := h.api.GetTeamById(ctx, opt.TeamId)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrInvalidRequest(err))
	}

	return c.JSON(http.StatusOK, response.Success(team, "retrieved team by id"))
}

func (h *Handlers) ListTeams(c echo.Context) (err error) {
	listDto := &dto.ListTeam{
		ListTeamDTO: teams.ListTeamDTO{
			Limit:   30,
			Offset:  0,
			Sort:    "name",
			SortDir: "asc",
		},
	}

	if err = c.Bind(listDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(listDto); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tms, err := h.api.ListTeams(ctx, &listDto.ListTeamDTO)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrInvalidRequest(err))
	}

	return c.JSON(http.StatusOK, response.Success(tms, "list of teams"))
}

func (h *Handlers) ListTeamMembers(c echo.Context) (err error) {
	listDto := &dto.ListTeamMembers{
		ListTeamMembersDTO: teams.ListTeamMembersDTO{
			Limit:   30,
			Offset:  0,
			Sort:    "name",
			SortDir: "asc",
		},
	}

	if err = c.Bind(listDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(listDto); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tms, err := h.api.ListTeamMembers(ctx, listDto.Id, &listDto.ListTeamMembersDTO)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrInvalidRequest(err))
	}

	return c.JSON(http.StatusOK, response.Success(tms, "list of team members"))
}

func (h *Handlers) InviteTeamMember(c echo.Context) (err error) {
	inviteDto := &dto.InviteTeamMember{
		InviteeRole: teams.TeamMemberRole,
	}

	if err = c.Bind(inviteDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(inviteDto); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = h.api.InviteTeamMember(ctx, inviteDto.TeamId, inviteDto.InviteeEmail, teams.TeamMembershipRole(inviteDto.InviteeRole))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrInvalidRequest(err))
	}

	return c.JSON(http.StatusOK, response.Success(nil, "invited user to team"))
}

func (h *Handlers) JoinTeam(c echo.Context) (err error) {
	joinDto := &dto.JoinTeam{}
	if err = c.Bind(joinDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(joinDto); err != nil {
		return err
	}
	userId := c.Get(contexthandler.USER_ID).(string)
	userEmail := c.Get(contexthandler.USER_INFO).(*auth.User).Email

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = h.api.JoinTeamJoinTeamUsingToken(ctx, userId, userEmail, joinDto.TeamId, joinDto.Token)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrInvalidRequest(err))
	}

	return c.JSON(http.StatusOK, response.Success(nil, "user has joined the team"))
}

func (h *Handlers) RemoveTeamMember(c echo.Context) (err error) {
	rmDto := &dto.RemoveTeamMember{}
	if err = c.Bind(rmDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(rmDto); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = h.api.RemoveTeamMembership(ctx, rmDto.TeamId, rmDto.MembershipId)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrInvalidRequest(err))
	}

	return c.JSON(http.StatusOK, response.Success(nil, "removed user from team"))
}
