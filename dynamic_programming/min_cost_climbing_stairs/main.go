package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	var min int
	dp := make([]int, len(cost)+1)

	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		fmt.Printf("index %v, %v \n", i, len(cost))

		min = minVal(dp[i-2], dp[i-1])
		// fmt.Printf("isi min %v \n", min)

		// }
		dp[i] = min + cost[i]
	}
	// fmt.Printf("isi dp %v \n", dp)
	return minVal(dp[len(cost)-2], dp[len(cost)-1])
}

func minVal(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	// nums := []int{10, 15, 20}
	nums := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	minCost := minCostClimbingStairs(nums)
	fmt.Printf("minimum cost %v \n", minCost)
}
