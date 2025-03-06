package db

import (
	"coffeh/config"
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	client     *mongo.Client
	once       sync.Once
	connectErr error
)

type Store struct {
	client *mongo.Database
}

func NewStore() (*Store, func()) {
	client, disconnect := mustConnectToDb()
	store := &Store{
		client: client.Database("coffeh"),
	}

	return store, disconnect
}

func mustConnectToDb() (*mongo.Client, func()) {
	once.Do(func() {
		env := config.LoadEnv()
		clientOpts := options.Client().ApplyURI(env.DbURI)
		client, connectErr = mongo.Connect(clientOpts)
		if connectErr != nil {
			log.Panic(connectErr)
		} else {
			log.Println("Connected to MongoDB")
		}
	})

	if connectErr != nil {
		log.Panic(connectErr)
	}

	// Function to disconnect from DB
	disconnect := func() {
		if client != nil {
			if err := client.Disconnect(context.TODO()); err != nil {
				log.Println("❌ Error closing MongoDB connection:", err)
			} else {
				log.Println("✅ MongoDB connection closed")
			}
		}
	}

	return client, disconnect
}
