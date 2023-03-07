package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/1-platform/api-catalog/internal/api"
	"github.com/1-platform/api-catalog/internal/auth"
	"github.com/1-platform/api-catalog/pkg/logger"
	"github.com/akhilmhdh/authy"
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

func (srv *HTTPServer) Listen() error {
	return srv.server.Start(fmt.Sprintf("%s:%d", srv.cfg.Host, srv.cfg.Port))
}

func New(cfg *Config, log *logger.Logger, authMod *auth.Auth, a *api.API) (*HTTPServer, error) {
	h := &Handlers{api: a, logger: log}

	plogger := log.Desugar()
	e := echo.New()

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
			c.String(http.StatusOK, fmt.Sprintf("Welcome: %s - %s", user.GetPID(), u.Name))
			return nil
		}
		c.String(http.StatusOK, "welcome")
		return nil
	})

	return &HTTPServer{cfg: cfg, server: e}, nil
}
