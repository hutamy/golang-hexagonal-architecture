package registry

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/hutamy/golang-hexagonal-architecture/internal/adapter/outbound/repository/postgresdb"
	"github.com/hutamy/golang-hexagonal-architecture/internal/port/outbound/registry"
	"github.com/jmoiron/sqlx"

	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryRegistry struct {
	mongo      *mongo.Database
	db         *sqlx.DB
	dbReplica  *sqlx.DB
	dbExecutor postgresdb.DBExecutor
	cache      *redis.Client
}

func NewRepositoryRegistry(mongo *mongo.Database, db, rdb *sqlx.DB, cache *redis.Client) registry.RepositoryRegistry {
	return &RepositoryRegistry{
		mongo:     mongo,
		db:        db,
		dbReplica: rdb,
		cache:     cache,
	}
}

func (r *RepositoryRegistry) DoInTransaction(ctx context.Context, txFunc registry.InTransaction) (err error) {
	var tx *sqlx.Tx

	registry := r

	if r.dbExecutor == nil {
		tx, err = r.db.BeginTxx(ctx, nil)

		if err != nil {
			return
		}

		defer func() {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p) // re-throw panic after Rollback
			} else if err != nil {
				rErr := tx.Rollback() // err is non-nil; don't change it
				if rErr != nil {
					err = rErr
				}
			} else {
				err = tx.Commit() // err is nil; if Commit returns error update err
			}
		}()
		registry = &RepositoryRegistry{
			db:         r.db,
			dbReplica:  r.dbReplica,
			dbExecutor: tx,
			cache:      r.cache,
		}
	}

	err = txFunc(registry)
	return
}
