package registry

import (
	sRegistry "github.com/hutamy/golang-hexagonal-architecture/internal/port/inbound/registry"
	rRegistry "github.com/hutamy/golang-hexagonal-architecture/internal/port/outbound/registry"
)

type ServiceRegistry struct {
	repositoryRegistry rRegistry.RepositoryRegistry
}

func NewServiceRegistry(repositoryRegistry rRegistry.RepositoryRegistry) sRegistry.ServiceRegistry {
	return ServiceRegistry{
		repositoryRegistry: repositoryRegistry,
	}
}
