package mongodb_test

import (
	"os"
	"testing"

	db "github.com/michaelknyazev/go-tools/mongodb"
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

	defer db.Disconnect()
}
