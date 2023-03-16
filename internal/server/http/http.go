package http

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/1-platform/api-catalog/internal/api"
	"github.com/1-platform/api-catalog/internal/auth"
	"github.com/1-platform/api-catalog/internal/server/http/middlewares"
	"github.com/1-platform/api-catalog/internal/server/http/response"
	"github.com/1-platform/api-catalog/internal/teams"
	"github.com/1-platform/api-catalog/pkg/logger"
	"github.com/akhilmhdh/authy"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Config struct {
	Host         string
	Port         uint16
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DialTimeout  time.Duration
}

type HTTPServer struct {
	server *echo.Echo
	cfg    *Config
}

type RequestValidator struct {
	validator *validator.Validate
}

func (rv *RequestValidator) Validate(i any) error {
	if err := rv.validator.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return response.ErrValidationFailed(validationErrors)
	}
	return nil
}

func customErrorHandler(err error, c echo.Context) {
	if res, ok := err.(*response.Response); ok {
		c.Logger().Error(res.Err)
		c.JSON(res.HTTPStatusCode, res)
		return
	}
	c.Logger().Error(err)
	c.JSON(http.StatusInternalServerError, &response.Response{
		Success: false,
		Message: "Something went wrong",
	})
}

func (srv *HTTPServer) Listen() error {
	return srv.server.Start(fmt.Sprintf("%s:%d", srv.cfg.Host, srv.cfg.Port))
}

func New(cfg *Config, log *logger.Logger,
	authMod *auth.Auth, a *api.API) (*HTTPServer, error) {
	h := &Handlers{api: a, logger: log}

	plogger := log.Desugar()
	e := echo.New()

	// echo override
	vl := validator.New()
	vl.RegisterValidation("is-pid", validatePid)
	e.Validator = &RequestValidator{validator: vl}
	e.HTTPErrorHandler = customErrorHandler

	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			plogger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))

	// custom middlewares
	e.Use(echo.WrapMiddleware(authMod.InjectMiddleware()))
	authMiddleware := middlewares.NewAuthMiddleware(a)

	e.GET("/health", h.Health)

	// team route
	teamRoute := e.Group("/teams", authMiddleware.IsAuthenticated)
	teamRoute.POST("/", h.CreateTeam)
	teamRoute.PATCH("/:teamId/", h.UpdateTeam, authMiddleware.HasTeamMembershipAcess([]teams.TeamMembershipRole{
		teams.TeamAdminRole, teams.TeamMemberRole,
	}))
	teamRoute.DELETE("/:teamId/", h.DeleteTeam, authMiddleware.HasTeamMembershipAcess([]teams.TeamMembershipRole{
		teams.TeamAdminRole,
	}))
	teamRoute.GET("/", h.ListTeams)
	teamRoute.GET("/:teamId/", h.GetTeamById)

	// team members route
	teamMemberRoute := e.Group("/teams/:teamId/members", authMiddleware.IsAuthenticated)
	teamMemberRoute.GET("/", h.ListTeamMembers)
	teamMemberRoute.POST("/", h.InviteTeamMember, authMiddleware.HasTeamMembershipAcess([]teams.TeamMembershipRole{
		teams.TeamAdminRole,
	}))
	teamMemberRoute.POST("/join/", h.JoinTeam)
	teamRoute.DELETE("/:membershipId/", h.RemoveTeamMember, authMiddleware.HasTeamMembershipAcess([]teams.TeamMembershipRole{
		teams.TeamAdminRole,
	}))

	e.GET("/", func(c echo.Context) error {
		user, err := authy.GetUserInfo(c.Request().Context())
		if err == nil {
			u := user.(*auth.User)
			c.JSON(http.StatusOK, response.Success(map[string]any{
				"email": user.GetPID(),
				"name":  u.Name,
			}, "user logged in"))
			return nil
		}
		c.JSON(http.StatusOK, response.Success(nil, "user not logged in"))
		return nil
	})

	return &HTTPServer{cfg: cfg, server: e}, nil
}

// custom validaton
var pidRegex, _ = regexp.Compile("^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$")

func validatePid(fl validator.FieldLevel) bool {
	return pidRegex.MatchString(fl.Field().String())
}
