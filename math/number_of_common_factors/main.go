package main

import "fmt"

func commonFactors(a int, b int) int {
	if a < b {
		a, b = b, a
	}

	common := 0
	max := a % b
	if max == 0 {
		max = b
	}
	for n := 1; n <= max; n++ {
		if a%n == 0 && b%n == 0 {
			common += 1
		}
		// fmt.Println(n)
	}
	return common
}

func main() {
	// a, b, exp := 12, 6, 4
	// a, b, exp := 25, 30, 2
	a, b, exp := 12, 8, 3
	// 12, 8, 3
	// 12 mod 8 = 4
	// 12/4 and 8/4
	// 12/2 and 8/2
	// 12/1 and 8/1

	// 16, 5

	// 12, 3, 2
	// 12 mod 3 = 0
	// 12/3 and 3/3
	// 12/1 and 3/1
	res := commonFactors(a, b)
	fmt.Printf("expected %v, result %v", exp, res)
}
