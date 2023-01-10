package main

import "fmt"

func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}
	prefix, suffix := make([]int, n), make([]int, n)
	prefix[0] = 1
	suffix[n-1] = n
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + i + 1
		suffix[n-1-i] = suffix[n-i] + n - i

		// pivot
		if prefix[i] == suffix[i] {
			return i + 1
		}
	}
	return -1
}

func main() {
	// n, exp := 8, 6
	// n, exp := 1, 1
	n, exp := 4, -1
	res := pivotInteger(n)
	fmt.Printf("expected %v, result %v ", exp, res)
}
