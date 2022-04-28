package mongodb_test

import (
	"testing"

	db "github.com/michaelknyazev/go-tools/mongodb"
	"github.com/spf13/viper"
)

func TestConnect(t *testing.T) {
	uri := viper.Get("MONGO_URI").(string)

	if err := db.Init(uri); err != nil {
		t.Fatal("Can't parse the mongo URI string")
	}

	if err := db.Connect(); err != nil {
		t.Fatal("Can't connect to database.")
	}
}

func TestDatabaseDisconnect(t *testing.T) {
	if err := db.Disconnect(); err != nil {
		t.Fatal("Can't disconnect from database")
	}
}
