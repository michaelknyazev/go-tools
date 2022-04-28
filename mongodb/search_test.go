package mongodb_test

import (
	"encoding/json"
	"testing"

	"github.com/michaelknyazev/go-tools/helpers"
	"github.com/michaelknyazev/go-tools/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestSearchIndexModel struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title" bson:"title"`
	Content string             `json:"content" bson:"content"`
}

func TestCreateIndex(t *testing.T) {
	TestConnect(t)

	defer TestDatabaseDisconnect(t)

	var tmp []TestSearchIndexModel
	var index []interface{}

	jsonData, err := helpers.ParseJsonFile("../mock_data/TestIndex.json")

	if err != nil {
		t.Fatal("Missing the test index")
	}

	json.Unmarshal(jsonData, &tmp)

	for _, record := range tmp {
		record.ID = primitive.NewObjectID()

		index = append(index, record)
	}

	mongodb.CreateMany(getTestCollection(), index)

	if err := mongodb.CreateTextIndex(getTestCollection(), []string{"title", "content"}, false, "en"); err != nil {
		t.Fatal(err)
	}
}

func TestSearchIndex(t *testing.T) {
	TestConnect(t)

	defer TestDatabaseDisconnect(t)

	query := "Lady"

	_, err := mongodb.SearchTextIndex(getTestCollection(), query)

	if err != nil {
		t.Fatal(err)
	}
}
