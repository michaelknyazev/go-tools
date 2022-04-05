package mongodb_test

import (
	"testing"

	db "github.com/michaelknyazev/go-tools/mongodb"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestInit(t *testing.T) {
	viper.SetConfigFile("../.env")
	viper.ReadInConfig()

	t.Log("Loaded .env")
}

func TestConnect(t *testing.T) {
	testUri := viper.Get("MONGO_URI").(string)
	err := db.Init(testUri)

	if err != nil {
		t.Fatal("Can't parse the mongo URI string")
	}

	err = db.Connect()

	if err != nil {
		t.Fatal("Can't connect to database.")
	}
}

type testModel struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

func getTestCollection() *mongo.Collection {
	return db.GetCollection("go_tools_test")
}

func TestCreateItem(t *testing.T) {
	var testItem testModel

	testItem.ID = primitive.NewObjectID()
	testItem.Name = "test"

	_, err := db.CreateItem(getTestCollection(), testItem)

	if err != nil {
		t.Fatal("Can't create new item in database", err)
	}
}

func TestDatabaseDisconnect(t *testing.T) {
	defer db.Disconnect()
}
