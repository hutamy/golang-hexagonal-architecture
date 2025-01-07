package registry

import "github.com/hutamy/golang-hexagonal-architecture/internal/port/outbound/registry"

type ServiceRegistry struct {
	repositoryRegistry registry.RepositoryRegistry
}

func NewServiceRegistry(repositoryRegistry registry.RepositoryRegistry) ServiceRegistry {
	return ServiceRegistry{
		repositoryRegistry: repositoryRegistry,
	}
}
