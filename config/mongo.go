package config

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClientInstance *mongo.Client
	mongoOnce           sync.Once
	mongoErr            error
)

func GetMongoClient(ctx context.Context, uri string) (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)

		ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		mongoClientInstance, mongoErr = mongo.Connect(ctxWithTimeout, clientOptions)
		if mongoErr != nil {
			log.Fatal(mongoErr)
		}

		mongoErr = mongoClientInstance.Ping(ctxWithTimeout, nil)
		if mongoErr != nil {
			log.Fatal(mongoErr)
		}

		log.Println("Connected to MongoDB!")
	})
	return mongoClientInstance, mongoErr
}
