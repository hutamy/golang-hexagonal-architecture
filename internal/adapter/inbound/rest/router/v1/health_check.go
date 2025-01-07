package v1

import (
	hv1 "github.com/hutamy/golang-hexagonal-architecture/internal/adapter/inbound/rest/handler/v1"

	"github.com/labstack/echo/v4"
	echoDatadog "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

func HealthCheckRouter(e *echo.Echo, handler *hv1.Handler) {
	e.GET("/api/v1/health-check", handler.HealthCheckHandler(), echoDatadog.Middleware())
	e.GET("/", handler.HealthCheckHandler(), echoDatadog.Middleware())
}
