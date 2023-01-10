package main

import "fmt"

func isSubsequence(s, t string) bool {
	if s == "" {
		return true
	}

	if t == "" {
		return false
	}

	for i := 0; i < len(s); {
		for j := 0; j < len(t) && i < len(s); {
			// compare
			if s[i] == t[j] {
				// fmt.Printf("compare %v and %v \n", s[i], t[j])
				i++
			}

			// stopper, karena i tidak increment dan j increment terus
			if i < len(s) && j == len(t)-1 {
				return false
			}
			j++
		}
	}
	return true
}

func main() {
	// s := "axc" // false
	// s := "abc" // true
	// s := "" // true
	// t := "ahbgdc"

	// s := "abc"
	// t := ""

	s := "b"
	t := "abc"
	result := isSubsequence(s, t)
	fmt.Printf("result %v\n", result)
}
