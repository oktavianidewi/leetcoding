package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	var sisa int
	for n > 2 {
		if sisa < n%2 {
			sisa = n % 2
		}
		n = n >> 1
	}

	return sisa <= 0
}

func main() {
	// n, output := 1, true
	// n, output := 5, false
	// n, output := 30, false
	// n, output := 16, true
	// n, output := 512, true
	// n, output := 514, false
	n, output := 0, false
	result := isPowerOfTwo(n)
	fmt.Printf("expected %v, result %v \n", output, result)
}
