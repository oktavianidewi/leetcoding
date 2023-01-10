package main

import (
	"fmt"
)

func increasingTriplet_dewi(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	var tripletCounter int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			tripletCounter++
		} else {
			tripletCounter = 0
		}

		if tripletCounter >= 1 {
			return true
		}
	}
	return false
}

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// exp := true

	// nums := []int{2, 1, 5, 0, 4, 6}
	// exp := true

	// nums := []int{5, 4, 3, 2, 1}
	// exp := false

	// error
	nums := []int{20, 100, 10, 12, 5, 13}
	exp := true

	// error
	// nums := []int{2, 4, -2, -3}
	// exp := false

	ans := increasingTriplet(nums)
	fmt.Printf("expected %v, ans %v \n", exp, ans)
}
