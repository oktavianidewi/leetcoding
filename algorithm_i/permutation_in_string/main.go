package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	window := make([]int, 26)
	for _, c := range s1 {
		window[c-'a']--
	}
	fmt.Printf("window1 %v \n", window)

	for l, r := 0, 0; r < len(s2); r++ {
		window[s2[r]-'a']++

		for ; window[s2[r]-'a'] > 0; l++ {
			window[s2[l]-'a']--
		}
		fmt.Printf("window2 %v \n", window)

		if r-l+1 == len(s1) {
			return true
		}
	}

	return false
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	output := true

	// s1 := "ab"
	// s2 := "eidboaoo"
	// output := false

	result := checkInclusion(s1, s2)
	fmt.Printf("expected %v, result %v \n", output, result)
}
