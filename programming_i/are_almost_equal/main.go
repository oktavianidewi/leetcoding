package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
	notEqualCounter := 0
	hashWord1 := make([]int, 26)
	hashWord2 := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			notEqualCounter++
		}
		if notEqualCounter > 2 {
			return false
		}

		hashWord1[s1[i]-'a']++
		hashWord2[s2[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if hashWord1[i] != hashWord2[i] {
			return false
		}
	}
	return true
}

func main() {
	// s1 := "bank"
	// s2 := "kanb"
	// exp := true

	// s1 := "kelb"
	// s2 := "kelb"
	// exp := true

	s1 := "abcd"
	s2 := "dcba"
	exp := false

	// s1 := "attack"
	// s2 := "defend"
	// exp := false
	res := areAlmostEqual(s1, s2)
	fmt.Printf("areAlmostEqual exp %v, res %v \n", exp, res)
}
