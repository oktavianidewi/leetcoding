package main

import (
	"fmt"
	"math"
)

func deleteAndEarn_dewi(nums []int) int {
	end := len(nums) - 1
	mid := end / 2

	if mid == 0 {
		return nums[0]
	}

	points := make([]int, mid+2)
	points[0], points[1] = 0, max(nums[0], nums[1])

	j := 2
	maxNums := math.MinInt32
	for i := 1; i <= mid; {
		fmt.Printf("i %v %v %v \n", i, len(nums), end)
		if 2*i == end {
			maxNums = nums[2*i]
		} else {
			maxNums = max(nums[2*i], nums[(2*i)+1])
		}

		points[j] = points[j-1] + maxNums

		fmt.Printf("i %v maxNums %v \n", i, maxNums)
		i++
		j++
	}
	fmt.Printf("points %v \n", points)

	return points[len(points)-1]
	// return max(points[len(nums)-1], points[len(nums)-2])
}
func maxFromArray(nums []int) int {
	return 0
}

// solutionku udah mirip dengan solution yang benar, tapi kenapa
// nums := []int{3, 1} // expect: 4, why???? harusnya kan 3. karena 1 dihapus.

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// why should make a 10001 length-size of array? lol
	size := len(nums) + 5
	dp := make([]int, size)
	for _, n := range nums {
		dp[n] += n
	}
	fmt.Printf("dp %v\n", dp)

	for i := 2; i < size; i++ {
		dp[i] = max(dp[i-1], dp[i]+dp[i-2])
	}

	fmt.Printf("dp %v\n", dp)

	return dp[size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// nums := []int{3, 4, 2}
	nums := []int{2, 2, 3, 3, 3, 4}
	// nums := []int{90, 2, 100, 89, 3, 4}
	// nums := []int{90}
	// nums := []int{3, 1} // expect: 4, why???? harusnya kan 3. karena 1 dihapus
	x := deleteAndEarn(nums)
	fmt.Printf("result %v \n", x)
}
