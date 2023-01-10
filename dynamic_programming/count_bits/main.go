package main

import (
	"fmt"
	"math"
)

func countBits(n int) []int {
	if n <= 0 {
		return []int{0}
	}
	dp := make([]int, n+1)
	dp[1] = 1
	j := 1.0
	for i := 0; i <= n; i++ {
		if i == int(math.Pow(2, j)) {
			fmt.Printf("hasil %v dan %v \n", i, int(math.Pow(2, j)))
			dp[i] = 1
			j++
			i++
		}

		if i >= 3 && i <= n {
			powerVal := int(math.Pow(2, j-1))
			// fmt.Printf("proses %v,  %v \n", i, powerVal)
			dp[i] = dp[powerVal] + dp[i-powerVal]
		}

	}
	return dp
}

func main() {
	res := countBits(16)
	fmt.Printf("res %v \n", res)
}
