package main

import (
	"fmt"
	"unicode"
)

func main() {
	var word string
	fmt.Scanln(&word)
	var cnt int = 1
	for _, r := range word {
		if unicode.IsUpper(r) {
			cnt += 1
		}
	}
	fmt.Println(cnt)
}
