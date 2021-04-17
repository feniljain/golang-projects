package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	var n int
	fmt.Scanln(&s)
	fmt.Scanln(&n)

	if n > 26 {
		n %= 26
	}
	var newS string = ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			newS += cipher(r, n)
		} else {
			newS += string(r)
		}
	}
	fmt.Println(newS)
}

func cipher(r rune, n int) string {
	if unicode.IsUpper(r) {
		if int(r)+n <= 90 {
			return string(int(r) + n)
		} else {
			return string(64 + ((int(r) + n) - 90))
		}
	} else {
		if int(r)+n <= 122 {
			return string(int(r) + n)
		} else {
			return string(96 + ((int(r) + n) - 122))
		}
	}
}
