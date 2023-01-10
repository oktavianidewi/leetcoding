package main

import "fmt"

func canJump(nums []int) bool {
	goal := len(nums) - 1

	if goal == 0 {
		return true
	}

	steps := make([]int, goal)

	steps[0] = nums[0]
	i := 1

	for steps[i-1] < goal {
		// return false when steps is not leading to goal
		if i >= goal {
			return false
		}
		steps[i] = nums[steps[i-1]] + steps[i-1]

		fmt.Printf("steps %v, goal %v \n", steps, goal)

		i++
	}
	return true
}

func main() {
	// nums := []int{2, 3, 1, 1, 4}

	// test case bermasalah, untuk value ini harusnya return false, tapi yang benar return true?
	nums := []int{2, 5, 0, 0}

	// nums := []int{0}
	// nums := []int{2, 0}
	// nums := []int{1, 2}
	// nums := []int{3, 2, 1, 0, 4}
	res := canJump(nums)
	fmt.Printf("result %v \n", res)
}
