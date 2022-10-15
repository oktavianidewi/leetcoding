package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return x >= 0
	}
	// with 2 pointer
	byteX := []byte(strconv.Itoa(x))
	left := 0
	right := len(byteX) - 1
	mid := (left + right) / 2
	for left <= mid && right >= mid {
		if byteX[left] == byteX[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func main() {
	// x := 12345954321
	x := 121
	// x := -121
	res := isPalindrome(x)
	fmt.Printf("res %v", res)

}
