package ConfigValidator

import (
	"fmt"
	"testing"
)

func TestAgeValidator(t *testing.T) {
	var age uint
	age = 10
	var result = ValidateAge(age)
	if result != true {
		t.Error(fmt.Sprintf("Expected: true, got: %v", result))
	}
}

func TestFirstNameValidator(t *testing.T) {
	var name string
	name = "Альберт"
	var result = ValidateFirstName(name)
	if result != true {
		t.Error(fmt.Sprintf("Expected: true, got: %v", result))
	}
}
