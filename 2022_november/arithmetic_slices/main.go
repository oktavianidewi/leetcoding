package main

import "fmt"

/*
sources:
https://dev.to/seanpgallivan/arithmetic-slices-4pla
https://leetcode.com/problems/arithmetic-slices/discuss/1816373/Golang-Solution-DP
*/

// func numberOfArithmeticSlices(nums []int) int {
// 	if len(nums) < 2 {
// 		return 0
// 	}

// 	var ans int
// 	diff := nums[1] - nums[0]
// 	count := 1

// 	for i := 2; i < len(nums); i++ {
// 		newDiff := nums[i] - nums[i-1]
// 		if newDiff != diff {
// 			count = 1
// 			diff = newDiff
// 		} else {
// 			ans += count
// 			count++
// 		}
// 	}

// 	return ans
// }

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	counter := 0
	consecutive := 1
	prevDiff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		currDiff := nums[i] - nums[i-1]
		if currDiff == prevDiff {
			counter += consecutive
			consecutive++
		} else {
			prevDiff = currDiff
			consecutive = 1
		}
	}
	return counter
}

func main() {
	// nums, expected := []int{1, 2, 3, 4}, 3
	nums, expected := []int{1}, 0
	result := numberOfArithmeticSlices(nums)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
