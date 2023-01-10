package main

import "fmt"

func xorOperation(n int, start int) int {
	xorResult := start
	for i := 1; i < n; i++ {
		xorResult ^= (start + (2 * i))
	}
	return xorResult
}

func main() {
	// n, start, exp := 5, 0, 8
	n, start, exp := 4, 3, 8
	res := xorOperation(n, start)
	fmt.Printf("expected %v, res %v \n", exp, res)
}
