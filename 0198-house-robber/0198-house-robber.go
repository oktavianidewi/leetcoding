
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums)+1)
	dp[0], dp[1], dp[2] = 0, nums[0], nums[1]

	for i := 2; i < len(nums); i++ {
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
