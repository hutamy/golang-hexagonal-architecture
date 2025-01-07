package datastore

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func NewRedisConn(ctx context.Context, host, password string, port int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       0,
	})

	logrus.Info("Redis Connection: ", client.Ping(ctx))

	if err := client.Ping(ctx).Err(); err != nil {
		logrus.Error("Redis Connection Error Log: ", err)
		return nil, err
	}

	return client, nil
}
