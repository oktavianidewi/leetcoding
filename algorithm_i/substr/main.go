package main

import (
	"fmt"
	"strings"
)

func solution(word, part string) bool {
	return strings.Contains(word, part)
}
func main() {
	// x := solution("abc", "bc") // returns true
	x := solution("abc", "d") // returns false
	fmt.Printf("solutions %v \n", x)
}
