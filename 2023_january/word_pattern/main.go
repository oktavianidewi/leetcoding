package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	patternMap := make(map[rune]string)
	wordMap := make(map[string]rune)

	for i := 0; i < len(pattern); i++ {
		// fmt.Println(pattern[i])

		if found, ok := patternMap[rune(pattern[i])]; !ok {
			patternMap[rune(pattern[i])] = words[i]
		} else {
			if found != words[i] {
				return false
			}
		}

		if found, ok := wordMap[words[i]]; !ok {
			wordMap[words[i]] = rune(pattern[i])
		} else {
			if found != rune(pattern[i]) {
				return false
			}
		}

	}
	// fmt.Println(patternMap)
	// fmt.Println(wordMap)
	return true
}

func main() {
	// pattern := "abba"
	// s := "dog cat cat dog"
	// exp := true

	// pattern := "abba"
	// s := "dog cat cat fish"
	// exp := false

	pattern := "abba"
	s := "dog dog dog dog"
	exp := false

	// pattern := "abc"
	// s := "dog cat dog"
	// exp := false

	res := wordPattern(pattern, s)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
