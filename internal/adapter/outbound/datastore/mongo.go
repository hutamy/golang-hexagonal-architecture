package datastore

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoDbOnce sync.Once
)

func NewMongoDBConn(user, password, host, name string) (client *mongo.Client, err error) {
	mongoDbOnce.Do(func() {
		connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s", user, password, host, name)
		commandMonitor := &event.CommandMonitor{
			Started: func(_ context.Context, evt *event.CommandStartedEvent) {
				log.Printf("Started command: %s - %s\n", evt.CommandName, evt.Command)
			},
			Succeeded: func(_ context.Context, evt *event.CommandSucceededEvent) {
				log.Printf("Succeeded command: %s - %d ms\n", evt.CommandName, evt.DurationNanos/1e6)
			},
			Failed: func(_ context.Context, evt *event.CommandFailedEvent) {
				log.Printf("Failed command: %s - %d ms\n", evt.CommandName, evt.DurationNanos/1e6)
			},
		}

		clientOptions := options.Client().ApplyURI(connectionString).
			SetMonitor(commandMonitor).
			SetConnectTimeout(300 * time.Second).
			SetMaxConnIdleTime(30 * time.Second).
			SetRetryWrites(true).
			SetRetryReads(true)
		mgoClient, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatalf("Error connecting to MongoDB: %v", err)
		}

		err = mgoClient.Ping(context.TODO(), nil)
		if err != nil {
			panic(fmt.Sprintf("Failed to ping MongoDB: %v", err))
		}

		client = mgoClient
	})

	return client, err
}

func CloseMongo(db *mongo.Database) {
	log.Println("close mongo client connection ...")
	if db != nil {
		client := db.Client()
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Println(errors.Wrap(err, "close mongo client").Error())
		}
	}

	log.Println("mongo client connection closed")
}
