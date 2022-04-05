package validation_test

import (
	"github.com/michaelknyazev/go-tools/validation"
	"testing"
)

type TestValidationStruct struct {
	Name string `validation:"required"`
}

func TestValidationInitializing(t *testing.T) {
	validation.Init()

	testStruct := TestValidationStruct{"Michael"}

	if err := validation.Validate.Struct(&testStruct); err != nil {
		t.Fatal("Validation is not working")
	}
}
