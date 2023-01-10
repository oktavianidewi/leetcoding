package main

import (
	"fmt"
	"time"
)

func findMaxConsecutiveOnes_external_solutions(nums []int) int {
	c, tmp := 0, 0
	for _, v := range nums {
		if v == 1 {
			tmp++
			c = max(c, tmp)

		} else {
			tmp = 0
			fmt.Printf("nilai max %v \n", c)
			time.Sleep(1 * time.Second)
		}
	}
	return max(c, tmp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// hati2 dengan proses reset counter
// nilai max bisa dibikin function seperti di contoh solution
func findMaxConsecutiveOnes(nums []int) int {
	left, right := 0, len(nums)-1
	counter := 0
	max := 0
	counter_max := 0
	for left <= right {
		if nums[left] == 1 {
			counter++
		} else {
			if max < counter {
				max = counter
			}
			counter = 0
			fmt.Printf("nilai max: %v %v \n", max, counter_max)
		}
		left += 1
	}
	if counter > max {
		return counter
	}
	return max
}
