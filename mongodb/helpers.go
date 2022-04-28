package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(name string) *mongo.Collection {
	collection := connection.Database(configuration.Name).Collection(name)

	return collection
}

func CreateTextIndex(collection *mongo.Collection, fields []string, unique bool, language string) error {
	weights := map[string]interface{}{}

	for i, field := range fields {
		weights[field] = len(fields) - i
	}

	model := mongo.IndexModel{
		Keys:    bson.M{"$**": "text"},
		Options: options.Index().SetDefaultLanguage(language).SetWeights(weights),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, model)

	return err
}
