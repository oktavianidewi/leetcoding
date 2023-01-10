package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var extraString string
	var smallestLength int

	if len(word2) < len(word1) {
		smallestLength = len(word2)
		extraString = word1[len(word2):]
	} else {
		smallestLength = len(word1)
		extraString = word2[len(word1):]
	}

	var res []string

	for i := 0; i < smallestLength; i++ {
		res = append(res, string(word1[i]), string(word2[i]))
	}

	return strings.Join(res, "") + extraString
}

func main() {
	// word1, word2, exp := "abc", "pqr", "apbqcr"
	// word1, word2, exp := "ab", "pqrs", "apbqrs"
	word1, word2, exp := "abcd", "pq", "apbqcd"
	res := mergeAlternately(word1, word2)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
