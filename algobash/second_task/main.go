package main

import (
	"fmt"
)

// Findminimum is your solution code.
func Findminimum(n int, a []int, b []int) int {
	// Your code starts here.

	minVal := -1
	var boys []int
	var girls []int
	for i, gender := range a {
		if gender == 0 {
			boys = append(boys, b[i])
		} else {
			girls = append(girls, b[i])
		}
	}

	// fmt.Println(boys, girls, minVal)

	for _, boyWeight := range boys {
		for _, girlWeight := range girls {
			diff := abs(boyWeight - girlWeight)
			// fmt.Println(diff)
			if minVal < 0 || minVal > diff {
				minVal = diff
			}
		}
	}
	return minVal
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	// n, a, b, expected := 5, []int{0, 0, 1, 0, 1}, []int{1, 13, 9, 7, 15}, 2
	n, a, b, expected := 3, []int{0, 0, 1}, []int{1, 2, 2}, 0
	// n, a, b, expected := 3, []int{0, 0, 0}, []int{1, 2, 2}, 0
	result := Findminimum(n, a, b)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
