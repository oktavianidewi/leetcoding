package main

import "fmt"

var vowelsMap = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func constructHalfMap(s string, start, end int) map[string]int {
	stringMap := make(map[string]int)
	for i := start; i < end; i++ {
		if _, isVowel := vowelsMap[rune(s[i])]; isVowel {
			stringMap["vowel"]++
		} else {
			stringMap["consonant"]++
		}
	}
	return stringMap
}

func halvesAreAlike(s string) bool {
	if len(s)%2 > 0 {
		return false
	}
	mid := len(s) >> 1
	leftMap := constructHalfMap(s, 0, mid)
	rightMap := constructHalfMap(s, mid, len(s))

	// fmt.Printf("leftMap %v \n", leftMap)
	// fmt.Printf("rightMap %v \n", rightMap)

	// compare leftMap and rightMap
	if leftMap["vowel"] == rightMap["vowel"] && leftMap["consonant"] == rightMap["consonant"] {
		return true
	}
	return false
}

func main() {
	// s, expected := "book", true
	s, expected := "textbook", false
	result := halvesAreAlike(s)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
