package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums)+1)
	dp[0], dp[1], dp[2] = 0, nums[0], nums[1]

	for i := 2; i < len(nums); i++ {
		dp[i+1] = max(dp[i-2], dp[i-1]) + nums[i]
	}
	fmt.Printf("dp %v \n", dp)
	return max(dp[len(dp)-1], dp[len(dp)-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// houses := []int{1, 2, 3, 1}
	houses := []int{2, 7, 9, 3, 1}
	// houses := []int{0}
	// houses := []int{2, 1, 1, 2}

	maxMoney := rob(houses)
	fmt.Printf("money %v \n", maxMoney)
}
