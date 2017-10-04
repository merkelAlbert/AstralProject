package ConfigValidator

import (
	"fmt"
	"unicode"
)

func ValidateAge(age uint) bool {
	if age < 120 {

		return true
	}
	log.Println(fmt.Sprintf("Age Can not be more than 120. Received: %v", age))
	return false
}

func ValidateFirstName(firstName string) bool {
	for _, c := range firstName {
		if !unicode.IsLetter(c) {
			log.Println("First name can bot contain non-letter symbol")
			return false
		}
	}
	return true
}
