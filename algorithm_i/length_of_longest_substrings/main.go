package main

import (
	"fmt"
	"time"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/1999/Golang-solution-(6ms)
func lengthOfLongestSubstring(s string) int {
	indexMap := make(map[byte]int)
	startIndex, subLen, maxSubLen := 0, 0, 0

	for i := 0; i < len(s); i++ {
		indx, exist := indexMap[s[i]]
		fmt.Printf("i %v, index %v, subLen %v, indexMap %v, startIndex %v \n", i, indx, subLen, indexMap, startIndex)
		if exist && indx >= startIndex {
			subLen = i - indx
			startIndex = indx + 1
		} else {
			subLen++
		}
		indexMap[s[i]] = i
		maxSubLen = max(subLen, maxSubLen)
		time.Sleep(1 * time.Second)
	}
	return maxSubLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring_dewi(s string) int {
	// sliding windows

	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	longestSubstring := 0
	hashStr := make(map[byte]bool)
	left, right := 0, len(s)-1
	i := 0
	j := 1

	for left < right || j == right {
		fmt.Printf("i %v, j %v, hash %v, left %v, right %v, longestSubstring %v \n", i, j, hashStr, left, right, longestSubstring)

		_, ok := hashStr[s[j]]
		if s[i] != s[j] && !ok {
			hashStr[s[i]] = true
			hashStr[s[j]] = true
		} else {
			hashStr = make(map[byte]bool)
			hashStr[s[i]] = true
		}
		left = j

		if longestSubstring < len(hashStr) {
			longestSubstring = len(hashStr)
		}
		i++
		j++
		time.Sleep(1 * time.Second)
	}
	return longestSubstring
}

func main() {
	// input, output := "abcabcbb", 3
	// input, output := "bbbbb", 1
	// input, output := "pwwkew", 3
	// input, output := "", 0
	// input, output := " ", 1
	// input, output := "au", 2
	// input, output := "dvdf", 3
	input, output := "anviaj", 5 // ini case aneh sih, kenapa expected output == 5, harusnya kan 4?? https://leetcode.com/submissions/detail/833610662/

	result := lengthOfLongestSubstring(input)
	fmt.Printf("expected %v, result %v \n", output, result)
}
