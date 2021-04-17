package cmd

import (
	"fmt"
	"unicode"
)

//Execute function acts as a starting point for the application
func Execute() {
	Init()
	//AddNumber("1234567890")
	//AddNumber("123 456 7891")
	//AddNumber("(123) 456 7892")
	//AddNumber("(123) 456-7893")
	//AddNumber("123-456-7894")
	//AddNumber("123-456-7890")
	//AddNumber("1234567892")
	//AddNumber("(123)456-7892")
	numbers := GetAllNumbers()
	fmt.Println(numbers)
	for _, number := range numbers {
		newNumber := normalizeNumber(number.number)
		UpdateNumber(number.id, newNumber)
	}
}

func normalizeNumber(number string) string {
	var newNumber string = ""
	for _, r := range number {
		if unicode.IsDigit(r) {
			newNumber += string(r)
		}
	}
	return newNumber
}
