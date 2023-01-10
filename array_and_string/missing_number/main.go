package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	return factorial(len(nums)) - sumAll(nums)
}

func sumAll(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total = total + nums[i]
	}
	return total
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num + factorial(num-1)
}

func main() {
	// 0, 1, 3
	// nums, expected := []int{3, 0, 1}, 2
	// nums, expected := []int{0, 1}, 2
	nums, expected := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8
	result := missingNumber(nums)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
