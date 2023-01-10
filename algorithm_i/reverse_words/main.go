package main

// https://leetcode.com/problems/reverse-words-in-a-string-iii/?envType=study-plan&id=algorithm-i

func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
