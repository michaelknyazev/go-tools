package mongodb_test

import (
	"testing"

	"github.com/spf13/viper"
)

func TestInit(t *testing.T) {
	viper.SetConfigFile("../.env")
	viper.ReadInConfig()

	t.Log("Loaded .env")
}
