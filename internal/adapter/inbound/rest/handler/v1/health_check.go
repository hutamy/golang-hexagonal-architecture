package v1

import (
	"net/http"

	"github.com/hutamy/golang-hexagonal-architecture/shared/util"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HealthCheckHandler() func(echo.Context) error {
	return func(c echo.Context) error {
		return util.SetResponse(c, http.StatusOK, "ASIK github.com/hutamy/golang-hexagonal-architecture Service", nil)
	}
}
