package main

import (
	"fmt"
	"time"
)

func topKFrequent(nums []int, k int) []int {
	result := []int{}
	if len(nums) == 1 {
		return append(result, nums[0])
	}

	i, j := 0, 0
	counter := 0
	for i < len(nums)-1 {
		fmt.Printf("i = %v, j = %v, nums[i] = %v \n", i, j, nums[i])
		if nums[i] == nums[j] {
			counter++
		}

		if nums[i] != nums[j] && counter >= k {
			result = append(result, nums[i])
			counter = 0
			i = j
			fmt.Printf("result, counter, i : %v, %v, %v\n", result, counter, i)
		} else {
			counter = 0
			i = j
		}
		j++
		time.Sleep(1 * time.Second)
	}

	return result
}
