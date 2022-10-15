package main

import (
	"strings"
)

func reverseWords(s string) string {
	sArr := strings.Fields(s)
	// fmt.Printf("arr string %v \n", sArr)
	i := 0
	j := len(sArr) - 1
	for i < j {
		// swap i and j
		sArr[i], sArr[j] = sArr[j], sArr[i]
		i++
		j--
	}
	return strings.Join(sArr, " ")
}
