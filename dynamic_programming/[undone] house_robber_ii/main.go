package main

import "fmt"

// https://leetcode.com/problems/house-robber-ii/discuss/343816/Golang-DP-100-solution

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	// compare head and tail value
	maxHead := isMaxHead(nums[0], nums[len(nums)-1])
	var res int
	if maxHead {
		fmt.Printf("RotateFront\n")
		res = rotateFront(nums)
	} else {
		fmt.Printf("RotateEnd\n")
		res = rotateEnd(nums)
	}
	// fmt.Printf("maxHead %v \n", maxHead)
	return res
}

func rotateEnd(nums []int) int {
	dp := make([]int, len(nums))
	maxNums := len(nums) - 1

	dp[0], dp[1], dp[2] = 0, nums[maxNums], nums[maxNums-1]
	// fmt.Printf("dp %v \n", dp)
	j := 0
	for i := maxNums - 2; i > 0; i-- {
		fmt.Printf("index %v, dp %v, maxVal compare %v %v \n", i, maxNums+1-i, dp[j], dp[j+1])
		// dp[len(nums)-i] = 88
		dp[maxNums+1-i] = max(dp[j], dp[j+1]) + nums[i]
		j++
	}
	// fmt.Printf("dp %v \n", dp)
	return max(dp[len(dp)-1], dp[len(dp)-2])
}
func rotateFront(nums []int) int {
	dp := make([]int, len(nums))

	dp[0], dp[1], dp[2] = 0, nums[0], nums[1]
	// fmt.Printf("dp %v \n", dp)
	for i := 2; i < len(nums)-1; i++ {
		dp[i+1] = max(dp[i-2], dp[i-1]) + nums[i]
	}
	// fmt.Printf("dp %v \n", dp)
	return max(dp[len(dp)-1], dp[len(dp)-2])
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isMaxHead(head, tail int) bool {
	return head >= tail
}

func main() {
	// arranged in a circle
	houses := []int{1, 2, 3, 1}
	// houses := []int{2, 7, 9, 3, 1}
	// houses := []int{0}
	// houses := []int{0, 0}
	// houses := []int{1, 2, 1, 1}
	// houses := []int{6, 3, 2, 10}
	// houses := []int{2, 1, 1, 2}
	// houses := []int{1, 2, 3}
	// houses := []int{2, 3, 2}
	// houses := []int{6, 3, 1, 7, 1, 1, 1}
	// houses := []int{1, 3, 1, 7, 1, 1, 6}
	// houses := []int{6, 3, 140, 20, 200}
	// houses := []int{200, 3, 140, 20, 10}

	maxMoney := rob(houses)
	fmt.Printf("money %v \n", maxMoney)
}
