package main

import (
	"fmt"
	"time"
)

func searchRange(nums []int, target int) []int {

}
func searchRange_dewi(nums []int, target int) []int {
	// res := []int{}
	res := []int{}
	// left, mid, right := 0, 0, len(nums)-1

	left, mid, right := 0, 0, len(nums)-1
	for left <= right && len(res) <= 2 {
		mid = (left + right) / 2
		fmt.Printf("left, mid, right, res : %v, %v, %v , %v \n", left, mid, right, res)
		if target == nums[mid] {
			low, high := mid, mid

			for low >= 0 && nums[low] == nums[mid] {
				low--
			}
			for high <= len(nums)-1 && nums[high] == nums[mid] {
				high++
			}
			res = append(res, low+1)
			res = append(res, high)

		} else if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
		time.Sleep(1 * time.Second)
	}

	// res = append(res, -1, -1)
	return res

}
