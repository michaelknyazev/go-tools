package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connection *mongo.Client

func LoadConfigurationAndConnect() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configuration.URI))

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	connection = client
}

func Disconnect() {
	ctx := context.Background()
	if err := connection.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}
