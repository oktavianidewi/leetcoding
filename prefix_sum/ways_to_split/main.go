package main

/*
explanation:
https://leetcode.com/problems/ways-to-split-array-into-three-subarrays/solutions/999113/java-scala-detailed-explanation-prefix-sum-binary-search/

golang
https://leetcode.com/problems/ways-to-split-array-into-three-subarrays/solutions/999383/go-3-pointers-o-n/
*/

import "fmt"

func waysToSplit(nums []int) int {
	var counter int
	len := len(nums)
	prefixSum := make([]int, len)
	prefixSum[0] = nums[0]
	for i := 1; i < len; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	right := len - 1
	for left := 0; left < len; {

		right += 2
		mid := (right - left) / 2

		fmt.Printf("left %v, mid %v, right %v \n", left, mid, right)

		left++
	}

	return counter
}

func main() {
	nums, expected := []int{1, 1, 1}, 1
	// nums, expected := []int{1, 2, 2, 2, 5, 0}, 3
	// nums, expected := []int{3, 2, 1}, 0
	result := waysToSplit(nums)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
