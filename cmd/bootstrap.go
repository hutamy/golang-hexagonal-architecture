package cmd

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/hutamy/golang-hexagonal-architecture/internal/adapter/outbound/datastore"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mdb   *mongo.Database
	db    *sqlx.DB
	rdb   *sqlx.DB
	cache *redis.Client
)

func InitMongoModule(user, password, host, name string) *mongo.Database {
	db, err := datastore.NewMongoDBConn(user, password, host, name)
	if err != nil {
		panic(err)
	}

	return db.Database(name)
}

func InitPostgresModule(host string, port int, user, password, dbName string) *sqlx.DB {
	db, err := datastore.NewPostgresDBConn(host, port, user, password, dbName)
	if err != nil {
		panic(err)
	}

	return db
}

func InitRedisModule(ctx context.Context, host, password string, port int) *redis.Client {
	client, err := datastore.NewRedisConn(ctx, host, password, port)
	if err != nil {
		panic(err)
	}

	return client
}
