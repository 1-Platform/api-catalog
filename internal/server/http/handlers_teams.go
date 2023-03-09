package http

import (
	"net/http"

	"github.com/1-platform/api-catalog/internal/server/http/dto"
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

	return c.JSON(http.StatusOK, SuccessRes(ctDto, "team creation success"))
}
