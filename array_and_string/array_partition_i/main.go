package main

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	pairSum := 0
	sort.Ints(nums)
	i := 0
	for i < len(nums)-1 {
		pairSum += nums[i]
		i += 2
	}
	return pairSum
}

// 1, 2, 2, 5, 6, 6
// 1, 2, 3, 4
