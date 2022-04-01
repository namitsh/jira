package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var DB *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatalf("Error configuring the database: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatalf("Error pinging to the database: %s", err.Error())
	}

	log.Println("Successfully pinged the database")
	log.Println("Successfully connected to database")
	return client
}

func GetCollection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}
