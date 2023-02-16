package http

import (
	"fmt"
	"time"

	"github.com/1-platform/api-catalog/internal/api"
	"github.com/labstack/echo/v4"
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

func New(cfg *Config, a *api.API) (*HTTPServer, error) {
	h := &Handlers{api: a}

	e := echo.New()
	e.GET("/health", h.Health)

	return &HTTPServer{cfg: cfg, server: e}, nil
}
