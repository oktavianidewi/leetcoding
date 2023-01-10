package main

import (
	"fmt"
	"time"
)

// with 2 pointer and swap method
func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1
	index_swap := 0

	for left <= right {
		if nums[left] != val {
			temp := nums[left]
			nums[left] = nums[index_swap]
			nums[index_swap] = temp
			index_swap++
		}
		fmt.Printf("after swap %v, %v\n", nums, index_swap)

		left++
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("removeElement %v \n", nums)
	return index_swap
}
