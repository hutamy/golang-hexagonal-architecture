package rest

import (
	"github.com/hutamy/golang-hexagonal-architecture/internal/port/inbound/registry"

	hv1 "github.com/hutamy/golang-hexagonal-architecture/internal/adapter/inbound/rest/handler/v1"
	v1 "github.com/hutamy/golang-hexagonal-architecture/internal/adapter/inbound/rest/router/v1"

	"github.com/labstack/echo/v4"
)

func Apply(e *echo.Echo, serviceRegistry registry.ServiceRegistry) {
	handlerV1 := hv1.New(serviceRegistry)
	v1.HealthCheckRouter(e, handlerV1)
}
