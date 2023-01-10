package main

import "fmt"

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}

func main() {
	// n, exp := 5, 10
	n, exp := 6, 6
	res := smallestEvenMultiple(n)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
