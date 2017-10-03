package ConfigValidator

import (
	"fmt"
	"unicode"
)

func ValidateAge(age uint) bool {
	if age < 120 {

		return true
	}
	fmt.Println(fmt.Sprintf("Age Can not be more than 120. Received: %v", age))
	return false
}

func ValidatefirstName(firstName string) bool {
	for _, c := range firstName {
		if !unicode.IsLetter(c) {
			fmt.Println("First name can bot contain non-letter symbol")
			return false
		}
	}
	return true
}
