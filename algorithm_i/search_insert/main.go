package main

import (
	"fmt"
	"time"
)

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
		fmt.Printf("left %v, right %v, mid %v \n", left, right, mid)
		time.Sleep(1 * time.Second)
	}
	if val := nums[left]; target > val {
		return left + 1
	}
	return left
}

func main() {
	// nums := []int{1, 3, 5, 6}
	// target, expected := 5, 2
	// target, expected := 2, 1
	// target, expected := 7, 4

	nums := []int{1, 3}
	target, expected := 2, 1

	result := searchInsert(nums, target)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
