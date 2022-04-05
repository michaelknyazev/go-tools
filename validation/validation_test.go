package validation_test

import (
	"testing"

	"github.com/michaelknyazev/go-tools/validation"
)

type TestValidationStruct struct {
	Name string `validation:"required"`
}

func TestInit(t *testing.T) {
	validation.Init()

	testStruct := TestValidationStruct{"Michael"}

	if err := validation.Validate.Struct(&testStruct); err != nil {
		t.Fatal("Validation is not working")
	}
}
