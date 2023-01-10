package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/largest-perimeter-triangle/discuss/1081040/Golang-solution-faster-than-95.65

func largestPerimeter(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	// 2, 3, 4, 4, 9, 15
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

// sebenernya bisa pakai cara ini, tapi kurang efisien, karena: harus menyimpan result sum di memory lain dan banyak step lain
func largestPerimeter_dewi(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	// 2, 3, 3, 4

	sort.Ints(nums)

	left := 0
	right := 2
	sum := 0
	var sums []int
	for i := left; i <= right; i++ {
		sum = sum + nums[i]
		if nums[left]+nums[left+1] <= nums[right] {
			if len(sums) == 0 {
				return 0
			}
			return sums[len(sums)-1]
		}

		fmt.Printf("sum %v \n", sum)

		if i == right {
			sums = append(sums, sum)
			sum = 0

			if len(nums)-1 == right {
				break
			}

			i = left
			left++
			right++
		}
	}
	return sums[len(sums)-1]
}

func main() {
	// sliding window

	// nums := []int{2, 1, 2}
	// exp := 5

	// nums := []int{1, 2, 1, 2}
	// nums := []int{1, 2, 1}
	// exp := 0

	// nums := []int{3, 2, 3, 4}
	// exp := 10

	// 2, 3, 3, 6
	// nums := []int{3, 6, 2, 3}
	// exp := 8

	// 2, 3, 4, 4, 9, 15
	nums := []int{3, 4, 15, 2, 9, 4}
	exp := 11

	ans := largestPerimeter(nums)
	fmt.Printf("expected %v, ans %v \n", exp, ans)
}
