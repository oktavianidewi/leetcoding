package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	maps := make(map[rune]int)
	mapt := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		maps[rune(s[i])]++
		mapt[rune(t[i])]++
	}

	for key, val := range maps {
		valt, ok := mapt[key]
		if !ok || val != valt {
			return false
		}
	}
	return true
}

func main() {
	s, t, expected := "anagram", "nagaram", true
	// s, t, expected := "rat", "car", false
	// s, t, expected := "aacc", "ccac", false

	result := isAnagram(s, t)
	fmt.Printf("expected %v, result %v \n", expected, result)

}
