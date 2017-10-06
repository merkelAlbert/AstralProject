//Данный пакет предназначен для валидации полей конфига
package configValidator

import (
	"fmt"
	"unicode"
)

//Данная функция проверяет поле "Age"
func ValidateAge(age uint) bool {
	if age < 120 {

		return true
	}
	log.Println(fmt.Sprintf("Age Can not be more than 120. Received: %v", age))
	return false
}

//Данная функция проверяет поле "FirstName"
func ValidateFirstName(firstName string) bool {
	for _, c := range firstName {
		if !unicode.IsLetter(c) {
			log.Println("First name can bot contain non-letter symbol")
			return false
		}
	}
	return true
}
