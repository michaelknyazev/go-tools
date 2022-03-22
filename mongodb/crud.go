package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateItem(collection *mongo.Collection, body interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, body)

	if err != nil {
		return result, err
	}

	return body, nil
}

func ReadItemByID(collection *mongo.Collection, itemId string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(itemId)
	result := collection.FindOne(ctx, bson.M{"_id": objectId})

	return result
}

func ReplaceItemById(collection *mongo.Collection, itemId string, update interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(itemId)

	filter := bson.M{"_id": objectId}
	updateResult, err := collection.ReplaceOne(ctx, filter, update)

	if err != nil {
		return updateResult, err
	}

	return update, nil
}

func DeleteItemById(collection *mongo.Collection, itemId string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectId, _ := primitive.ObjectIDFromHex(itemId)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectId})

	return result, err
}
