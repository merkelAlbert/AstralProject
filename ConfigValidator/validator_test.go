package configValidator

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

	age = 1000
	result = ValidateAge(age)
	if result != false {
		t.Error(fmt.Sprintf("Expected: false, got: %v", result))
	}
}

func TestFirstNameValidator(t *testing.T) {
	var name string
	name = "Альберт"
	var result = ValidateFirstName(name)
	if result != true {
		t.Error(fmt.Sprintf("Expected: true, got: %v", result))
	}

	name = "12345"
	result = ValidateFirstName(name)
	if result != false {
		t.Error(fmt.Sprintf("Expected: false, got: %v", result))
	}
}
