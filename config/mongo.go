package config

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClientInstance *mongo.Client
	mongoOnce           sync.Once
	mongoErr            error
)

func GetMongoClient(uri string) (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)
		ctx, cancel := context.WithTimeout(context.Background(), 10)
		defer cancel()

		mongoClientInstance, mongoErr = mongo.Connect(ctx, clientOptions)
		if mongoErr != nil {
			log.Fatal(mongoErr)
		}

		mongoErr = mongoClientInstance.Ping(ctx, nil)
		if mongoErr != nil {
			log.Fatal(mongoErr)
		}

		log.Println("Connected to MongoDB!")
	})
	return mongoClientInstance, mongoErr
}
