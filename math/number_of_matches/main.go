package main

import "fmt"

func numberOfMatches(n int) int {
	sumMatches := 0
	for n > 1 {
		match := n / 2
		sumMatches += (match)
		n = n - match
	}
	return sumMatches
}

func main() {
	// n, exp := 4, 3
	n, exp := 2, 1
	// n, exp := 7, 6
	// n, exp := 14, 13
	res := numberOfMatches(n)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
