package mongodb_test

import (
	"os"
	"testing"

	db "github.com/michaelknyazev/go-tools/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDatabaseInitializing(t *testing.T) {
	testUri := os.Getenv("MONGO_URI")

	err := db.Init(&testUri)

	if err != nil {
		t.Fatal("Can't parse the mongo URI string")
	}

	err = db.Connect()

	if err != nil {
		t.Fatal("Can't connect to database.")
	}
}

type TestModel struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func GetTestCollection() *mongo.Collection {
	return db.GetCollection("go_tools_test")
}

func TestCreateItem(t *testing.T) {
	var testItem TestModel

	testItem.ID = primitive.NewObjectID()
	testItem.Name = "test"

	_, err := db.CreateItem(GetTestCollection(), testItem)

	if err != nil {
		t.Fatal("Can't create new item in database", err)
	}
}

func TestDatabaseDisconnect(t *testing.T) {
	defer db.Disconnect()
}
