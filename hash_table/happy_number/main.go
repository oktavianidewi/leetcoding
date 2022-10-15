package main

import (
	"fmt"
	"strconv"
)

func isHappy(n int) bool {
	hash := make(map[int]int)

	sum := 0
	for sum != 1 {
		// restart sum
		sum = 0

		// convert int to string
		nStr := strconv.Itoa(n)

		for i := 0; i < len(nStr); i++ {
			// convert single string to int
			nSingle, _ := strconv.Atoi(string(nStr[i]))
			fmt.Printf("number %v \n", nSingle)
			sum = sum + (nSingle * nSingle)
		}
		n = sum
		// fmt.Printf("sum %v \n", sum)
		hash[n] = hash[n] + 1

		// detect circular value, if sum value is repeating, return false
		if hash[n] > 1 {
			return false
		}

		// time.Sleep(3 * time.Second)
	}
	return true
}
