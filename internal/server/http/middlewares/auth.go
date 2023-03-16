package middlewares

import (
	"context"
	"errors"
	"github.com/akhilmhdh/authy"
	"github.com/labstack/echo/v4"

	"github.com/1-platform/api-catalog/internal/api"
	"github.com/1-platform/api-catalog/internal/auth"
	"github.com/1-platform/api-catalog/internal/server/http/contexthandler"
	"github.com/1-platform/api-catalog/internal/server/http/response"
	"github.com/1-platform/api-catalog/internal/teams"
)

type AuthMiddleware struct {
	api *api.API
}

func NewAuthMiddleware(api *api.API) *AuthMiddleware {
	return &AuthMiddleware{api}
}

func (am *AuthMiddleware) IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := authy.GetUserInfo(c.Request().Context())
		if err == nil {
			u := user.(*auth.User)
			c.Set(contexthandler.USER_ID, u.Id.Hex())
			c.Set(contexthandler.USER_INFO, u)
			return next(c)
		}
		c.Logger().Error(err)
		return response.ErrUnauthorizedAccess(errors.New("user not logged in"))
	}
}

func (am *AuthMiddleware) HasTeamMembershipAcess(roles []teams.TeamMembershipRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			teamId := c.Param("teamId")
			if teamId == "" {
				return response.ErrInvalidRequest(errors.New("team not found"))
			}
			userId := c.Get(contexthandler.USER_ID).(string)
			teamMembership, err := am.api.GetTeamMembershipByUserId(context.TODO(), teamId, userId)
			if err != nil {
				c.Logger().Error(err)
				return response.ErrUnauthorizedAccess(errors.New("user doesn't have access to team"))
			}
			for _, role := range roles {
				if role == teamMembership.Role {
					return next(c)
				}
			}

			return response.ErrUnauthorizedAccess(errors.New("user doesn't have access to team"))
		}
	}
}
