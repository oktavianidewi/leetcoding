package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	end := len(nums) - 1

	if end == 0 {
		return nums[0]
	}

	sums := make([]int, end+1)
	sums[0] = nums[end]
	j := 0
	// fmt.Printf("isi initial sums %v \n", sums)

	var operand int
	for i := end - 1; i >= 0; {
		// fmt.Printf("tambah sums dan nums %v %v \n", sums[j], nums[i])
		if sums[j] < 0 {
			operand = 0
		} else {
			operand = sums[j]
		}
		sums[j+1] = operand + nums[i]
		i--
		j++
	}
	// fmt.Printf("isi sums %v \n", sums)
	max := getMaxArr(sums)
	return max
}

func getMaxArr(nums []int) int {
	maxVal := math.MinInt32
	for _, val := range nums {
		maxVal = max(val, maxVal)
	}
	return maxVal
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums := []int{1}
	// nums := []int{-1}
	// nums := []int{-2, -1}
	// nums := []int{-1, -2}
	nums := []int{5, 4, -1, 7, 8}
	maxSum := maxSubArray(nums)
	fmt.Printf("maxSum %v \n", maxSum)
}
