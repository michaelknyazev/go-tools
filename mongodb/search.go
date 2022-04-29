package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SearchTextIndex(collection *mongo.Collection, query string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"$text": bson.M{"$search": query}}

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}

	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func SearchTextField(collection *mongo.Collection, field string, query string) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	filter := bson.M{field: bson.M{"$regex": query, "$options": "im"}}

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}

	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
