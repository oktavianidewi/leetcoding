package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	n := 5
	res := climbStairs(n)
	fmt.Printf("num of steps %v \n", res)
}
