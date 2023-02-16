package http

import (
	"net/http"

	"github.com/1-platform/api-catalog/internal/api"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	api *api.API
}

func (h *Handlers) Health(c echo.Context) error {
	res, err := h.api.Health()

	if err != nil {
		return echo.NewHTTPError(http.StatusExpectationFailed, "Failed to get health")
	}

	return c.JSON(http.StatusOK, res)
}
