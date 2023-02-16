package http

import (
	"fmt"
	"time"

	"github.com/1-platform/api-catalog/internal/api"
	"github.com/1-platform/api-catalog/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Config struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DialTimeout  time.Duration
}

type HTTPServer struct {
	server *echo.Echo
	cfg    *Config
}

func (srv *HTTPServer) Listen() error {
	return srv.server.Start(fmt.Sprintf("%s:%s", srv.cfg.Host, srv.cfg.Port))
}

func New(cfg *Config, log *logger.Logger, a *api.API) (*HTTPServer, error) {
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

	e.GET("/health", h.Health)

	return &HTTPServer{cfg: cfg, server: e}, nil
}
