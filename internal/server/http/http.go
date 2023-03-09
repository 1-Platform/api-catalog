package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/1-platform/api-catalog/internal/api"
	"github.com/1-platform/api-catalog/internal/auth"
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
		return ErrValidationFailed(validationErrors)
	}
	return nil
}

func customErrorHandler(err error, c echo.Context) {
	if res, ok := err.(*Response); ok {
		c.Logger().Error(res.Err)
		c.JSON(res.HTTPStatusCode, res)
		return
	}
	c.Logger().Error(err)
	c.JSON(http.StatusInternalServerError, &Response{
		Success: false,
		Message: "Something went wrong",
	})
}

func (srv *HTTPServer) Listen() error {
	return srv.server.Start(fmt.Sprintf("%s:%d", srv.cfg.Host, srv.cfg.Port))
}

func New(cfg *Config, log *logger.Logger, authMod *auth.Auth, a *api.API) (*HTTPServer, error) {
	h := &Handlers{api: a, logger: log}

	plogger := log.Desugar()
	e := echo.New()
	e.Validator = &RequestValidator{validator: validator.New()}
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

	e.Use(echo.WrapMiddleware(authMod.InjectMiddleware()))
	e.GET("/health", h.Health)

	e.GET("/", func(c echo.Context) error {
		user, err := authy.GetUserInfo(c.Request().Context())
		if err == nil {
			u := user.(*auth.User)
			c.JSON(http.StatusOK, SuccessRes(map[string]any{
				"email": user.GetPID(),
				"name":  u.Name,
			}, "user logged in"))
			return nil
		}
		c.JSON(http.StatusOK, SuccessRes(nil, "user not logged in"))
		return nil
	})

	teamRoute := e.Group("/teams")
	teamRoute.POST("/", h.CreateTeam)

	return &HTTPServer{cfg: cfg, server: e}, nil
}
