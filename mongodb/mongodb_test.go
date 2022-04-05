package mongodb_test

import (
	"os"
	"testing"

	db "github.com/michaelknyazev/go-tools/mongodb"
)

func TestDatabaseInitializing(t *testing.T) {
	testUri := os.Getenv("MONGO_URI")

	db.Init(&testUri)
}
