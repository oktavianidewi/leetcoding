package main

import (
	"fmt"
	"time"
)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	if len(nums) == 1 {
		return nums[left]
	}

	// if already sorted
	if len(nums) > 2 && nums[left] < nums[right] {
		return nums[left]
	}

	for left <= right {
		mid := left + (right-left)/2
		fmt.Printf("left, mid, right %v, %v, %v \n", left, mid, right)

		if left == mid && nums[left] != nums[right] {
			if nums[left] < nums[right] {
				return nums[left]
			} else {
				return nums[right]
			}
		} else if left == mid && nums[left] == nums[right] {
			if len(nums) > 3 {
				left = 0
				right = left + (len(nums)-left)/2
				mid = left + (right-left)/2
			} else {
				return nums[left]
			}
		}

		// increment search s
		if nums[mid] > nums[right] {
			// search space ke kanan
			left = mid
		} else if nums[mid] == nums[right] {
			if nums[mid] <= nums[right] {
				left = mid
			} else if nums[mid] > nums[right] {
				right = mid
			}
		} else {
			// search space ke kiri
			right = mid
		}

		time.Sleep(1 * time.Second)
	}
	return -1
}
