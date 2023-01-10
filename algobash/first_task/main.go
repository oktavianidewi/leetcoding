package main

import "fmt"

// lumayan lah 11 out of 9 pass
// Solution is your solution code.
func dSolution(a int, b int) int {
	// Your code starts here.

	var bgcd func(a, b, res int) int

	bgcd = func(a, b, res int) int {
		switch {
		case a == b:
			return res * a
		case a%2 == 0 && b%2 == 0:
			return bgcd(a/2, b/2, 2*res)
		case a%2 == 0:
			return bgcd(a/2, b, res)
		case b%2 == 0:
			return bgcd(a, b/2, res)
		case a > b:
			return bgcd(a-b, b, res)
		default:
			return bgcd(a, b-a, res)
		}
	}
	return bgcd(a, b, 1)

	// return -1
}

// Solution is your solution code.
func Solution(n int, m int) int {
	if n == 0 || m == 0 {
		return 0
	}
	// Your code starts here.
	var result int
	min, max := minMax(n, m)
	// min, max := n, m

	if min%2 == 0 {
		mid := min / 2
		fmt.Printf("mid %v \n", mid)
		fmt.Printf("midss %v \n", max%min)

		if max%min == 0 {
			result = 3
		} else {
			result = 2
		}
		for i := 2; i < mid; i++ {
			if max%min == 0 || min%i == 0 {
				fmt.Printf("habis dibagi %v \n", i)
				result++
			}
		}
	} else {
		if max%min != 0 {
			result = 1
		} else {
			result = 2
		}
	}

	return result
}

func minMax(n, m int) (int, int) {
	if n < m {
		return n, m
	}
	return m, n
}

func main() {
	// n, m, expected := 8, 12, 3
	// n, m, expected := 7, 5, 1
	// n, m, expected := 7, 35, 2
	// n, m, expected := 10, 10, 4
	// n, m, expected := 14, 21, 3
	n, m, expected := 14, 100, 2
	// n, m, expected := 24, 12, 6
	// n, m, expected := 0, 12, 0
	result := Solution(n, m)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
