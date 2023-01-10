// https://shareablecode.com/snippets/golang-solution-for-leetcode-problem-implement-strstr-iQfi-ijfR
package main

import (
	"fmt"
	"time"
)

func strStr(haystack string, needle string) int {
	found := -1

	// needle length > haystack length
	if len(haystack)-1 < len(needle)-1 {
		return found
	}

	i, j := 0, 0
	matchCount := 0
	for i <= len(haystack)-1 {
		if haystack[i] == needle[j] {
			fmt.Printf("compare %v, %v , %v, %v \n", string(haystack[i]), string(needle[j]), found, matchCount)

			// save the first index matching
			if found < 0 {
				found = i
			}
			matchCount++
			if matchCount == len(needle) {
				return found
			}
			j++
			i++
		} else if haystack[i] != needle[j] && found > 0 && matchCount > 0 {
			fmt.Printf("reset all\n")
			matchCount = 0
			i = found
			j = 0
			found = -1
			i++
		} else {
			i++
		}
		time.Sleep(1 * time.Second)
	}
	return found
}
