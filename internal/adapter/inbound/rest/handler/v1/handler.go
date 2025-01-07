package v1

import (
	"github.com/hutamy/golang-hexagonal-architecture/internal/port/inbound/registry"
)

type Handler struct {
	serviceRegistry registry.ServiceRegistry
}

func New(serviceRegistry registry.ServiceRegistry) *Handler {
	return &Handler{
		serviceRegistry: serviceRegistry,
	}
}

func (h *Handler) GetServiceRegistry() registry.ServiceRegistry {
	return h.serviceRegistry
}
