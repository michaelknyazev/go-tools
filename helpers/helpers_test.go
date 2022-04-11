package helpers_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/michaelknyazev/go-tools/helpers"
)

func TestParseJson(t *testing.T) {
	path := "../domains.json"

	var target []string

	content, err := helpers.ParseJsonFile(path)

	if err != nil {
		panic(err)
	}

	json.Unmarshal(content, &target)

	fmt.Println(target)
}
