package http

import (
	"net/http"

	"github.com/1-platform/api-catalog/internal/api"
	"github.com/1-platform/api-catalog/pkg/logger"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	logger *logger.Logger
	api    *api.API
}

func (h *Handlers) Health(c echo.Context) error {
	res, err := h.api.Health()

	if err != nil {
		h.logger.Errorw("Health", err)
		return echo.NewHTTPError(http.StatusExpectationFailed, "Failed to get health")
	}

	return c.JSON(http.StatusOK, res)
}
