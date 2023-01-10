package main

import "fmt"

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	increasing_flag := false
	decreasing_flag := false

	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] < arr[left+1] {
			if decreasing_flag {
				return false
			}
			// initialize increase
			if left == 0 {
				increasing_flag = true
			}
			increasing_flag = increasing_flag && true
		} else if arr[left] > arr[left+1] {
			fmt.Printf("increasing_flag %v, compare %v , %v \n", increasing_flag, arr[left], arr[left+1])
			if !increasing_flag {
				return false
			}
			// initialize decrease
			if left != 0 && !decreasing_flag {
				decreasing_flag = true
			}
			decreasing_flag = decreasing_flag && true
		} else {
			// arr[left] equal to arr[right]
			increasing_flag = increasing_flag && false
			decreasing_flag = decreasing_flag && false
		}
		left++

	}

	if increasing_flag && decreasing_flag {
		return true
	} else {
		return false
	}
}
