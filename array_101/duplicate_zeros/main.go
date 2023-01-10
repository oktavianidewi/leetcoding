package main

import (
	"fmt"
)

func duplicateZeros(arr []int) {
	// left, right := 0, len(arr)-1
	var zero_index int
	left, right := 0, len(arr)-1
	for left <= right {
		if arr[left] == 0 {
			// edge case
			if left == right {
				zero_index = left
			} else {
				zero_index = left + 1
			}
			arr = insertIntoArr(arr, zero_index, 0)
			// skip additional 0 index
			left += 1
		}
		left += 1
		// fmt.Printf("array %v \n", arr)
		// time.Sleep(1 * time.Second)
	}
}

func insertIntoArr(arr []int, zero_index, num int) []int {
	// shift array dengan command copy
	fmt.Printf("destination, src %v, %v \n", arr[zero_index+1:], arr[zero_index:len(arr)-1])
	// copy(dest, src)
	copy(arr[zero_index+1:], arr[zero_index:len(arr)-1])
	arr[zero_index] = 0
	return arr
}
