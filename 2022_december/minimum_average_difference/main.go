package main

import (
	"fmt"
	"math"
)

func minimumAverageDifference_ans(nums []int) int {

	// k elements in the nums array (will be reused many times)
	var k int = len(nums)

	// create a prefix sum array
	var prefix []int = make([]int, k)
	prefix[0] = nums[0]
	for i := 1; i < k; i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	fmt.Printf("prefix %v \n", prefix)

	// for each element of prefix array,
	// calculate the average difference between the first i+1 elements
	// and the last k-i elements
	var answer int
	var minAverageDifference = math.MaxFloat64

	for i := 0; i < k; i++ {
		averageLeftHalf := float64(prefix[i] / (i + 1))
		// avoid division by 0 when averaging diferrence between all element and nothing
		var averageRightHalf float64
		if i+1 >= k {
			averageRightHalf = 0
		} else {
			fmt.Printf("avgRight comes from %v, minus %v , pembagi %v \n", prefix[k-1], prefix[i], k-i-1)
			averageRightHalf = float64((prefix[k-1] - prefix[i]) / (k - i - 1))

		}

		averageDifference := math.Abs(averageRightHalf - averageLeftHalf)
		fmt.Printf("averageLeftHalf %v, averageRightHalf %v, averageDifference %v \n", averageLeftHalf, averageRightHalf, averageDifference)

		// if we found the smallest one, save it and keep answer;
		if averageDifference < minAverageDifference {
			minAverageDifference = averageDifference
			answer = i
		}
	}

	return answer
}

func minimumAverageDifference(nums []int) int {
	// with prefix sum
	minAvg := math.MaxInt

	var minAvgIndex, avgDiff int
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	// loop through prefixSum
	for i := 0; i < len(nums); i++ {
		avgLeft := prefixSum[i] / (i + 1)

		var avgRight int
		// hitung avgRight
		if i+1 >= len(nums) {
			avgRight = 0
		} else {
			avgRight = (prefixSum[len(nums)-1] - prefixSum[i]) / (len(nums) - i - 1)
		}

		avgDiff = avgLeft - avgRight

		if avgDiff < 0 {
			avgDiff = -1 * avgDiff
		}

		if minAvg > avgDiff {
			minAvg = avgDiff
			minAvgIndex = i
		}
	}
	// fmt.Printf("prefixSum %v, avgDiff %v \n", prefixSum, avgDiff)

	return minAvgIndex
}

func main() {
	// nums, expected := []int{2, 5, 3, 9, 5, 3}, 3
	nums, expected := []int{0, 4, 3, 0, 0}, 0
	// nums, expected := []int{0}, 0
	result := minimumAverageDifference(nums)
	result_ans := minimumAverageDifference_ans(nums)
	fmt.Printf("expected %v, result %v \n", expected, result)
	fmt.Printf("expected_ans %v, result_ans %v \n", expected, result_ans)
}
