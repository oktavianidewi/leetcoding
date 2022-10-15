package main

import (
	"strings"
)

func reverseWords(s string) string {
	// not inline
	reversed := []string{}
	// mirip dengan split dengan " "
	sArr := strings.Fields(s)
	i := 0
	for i <= len(sArr)-1 {
		// swap i and j
		reversed = append(reversed, swapCharacter(sArr[i]))
		i++
	}
	return strings.Join(reversed, " ")
}

func swapCharacter(word string) string {
	i := 0
	j := len(word) - 1
	// proses mutable character, sehingga harus diubah ke array
	tempWord := make([]string, len(word))
	for i <= j {
		// swap i and j
		tempWord[i], tempWord[j] = string(word[j]), string(word[i])
		i++
		j--
	}
	// fmt.Printf("reversed word %v \n", tempWord)
	// fmt.Printf("reversed word %v \n", strings.Join(tempWord, ""))
	return strings.Join(tempWord, "")
}
