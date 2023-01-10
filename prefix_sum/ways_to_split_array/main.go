package main

import "fmt"

func waysToSplitArray(nums []int) int {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	var counter int
	// fmt.Printf("prefixSum %v \n", prefixSum)

	len := len(nums) - 1
	for i := 0; i < len; i++ {
		if prefixSum[i] >= (prefixSum[len] - prefixSum[i]) {
			counter++
		}
	}

	return counter
}

func main() {
	// nums, expected := []int{10, 4, -8, 7}, 2
	// nums, expected := []int{2, 3, 1, 0}, 2
	nums, expected := []int{0, 0}, 1
	output := waysToSplitArray(nums)
	fmt.Printf("expected %v, output %v \n", expected, output)
}
