package mongodb_test

import (
	"testing"

	db "github.com/michaelknyazev/go-tools/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type testModel struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func getTestCollection() *mongo.Collection {
	return db.GetCollection("go_tools_test")
}

func TestCreateItem(t *testing.T) {
	TestConnect(t)

	defer TestDatabaseDisconnect(t)

	var testItem testModel

	testItem.ID = primitive.NewObjectID()
	testItem.Name = "test"

	_, err := db.CreateItem(getTestCollection(), testItem)

	if err != nil {
		t.Fatal("Can't create new item in database", err)
	}
}
