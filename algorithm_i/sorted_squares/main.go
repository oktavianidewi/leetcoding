package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
		if nums[i] < 0 {
			nums[i] = -1 * nums[i]
		}
	}
	sort.Ints(nums)
	return nums
}

func main() {
	// nums := []int{-4, -1, 0, 3, 10}
	// output := []int{0, 1, 9, 16, 100}

	nums := []int{-7, -3, 2, 3, 11}
	output := []int{4, 9, 9, 49, 121}

	result := sortedSquares(nums)
	fmt.Printf("expected %v, result %v \n", output, result)
}
