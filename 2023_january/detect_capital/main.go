package main

import (
	"fmt"
)

func detectCapitalUse(word string) bool {
	firstChar := rune(word[0])
	uppercaseCounter := 0
	for i := 1; i < len(word); i++ {
		char := rune(word[i])
		if char >= rune('A') && char <= rune('Z') {
			uppercaseCounter++
		}
	}
	if firstChar >= rune('A') && firstChar <= rune('Z') {
		// first char uppercase
		if uppercaseCounter == 0 || uppercaseCounter == len(word)-1 {
			return true
		}
	} else if firstChar >= rune('a') && firstChar <= rune('z') {
		// first char lowercase
		if uppercaseCounter == 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func main() {
	var word string
	var exp bool

	// word, exp = "USA", true
	// word, exp = "mL", false
	// word, exp = "leetcode", true
	// word, exp = "Google", true
	word, exp = "FlaG", false
	res := detectCapitalUse(word)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
