package registry

import "context"

type InTransaction func(registry RepositoryRegistry) error
type RepositoryRegistry interface {
	DoInTransaction(ctx context.Context, txFunc InTransaction) error
}
