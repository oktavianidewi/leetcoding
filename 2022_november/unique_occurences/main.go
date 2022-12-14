package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	uniqueMap := make(map[int]int)
	for _, val := range arr {
		uniqueMap[val]++
	}
	// fmt.Printf("uniqueMap %v \n", uniqueMap)

	// compare
	uniqueCount := make(map[int]int)
	for _, count := range uniqueMap {
		if _, found := uniqueCount[count]; found {
			return false
		}
		uniqueCount[count] = 1
	}
	return true
}

func main() {
	// arr, expected := []int{1, 2, 2, 1, 1, 3}, true
	arr, expected := []int{1, 2}, false
	// arr, expected := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true

	// arr, expected := []int{-130, 21, -154, 159, -44, -126, 165, 68, -126, -126, -126, 128, -94, 165, -30, -44, -39, -94, 21, -130, 68, 68, 128, -130, -39, 181, 68, 68, 68, 139, 139, -39, 21, 21, -39, 68, 128, 131, -126, -154, -30, 165, 21, 159, 181, -39, -126, 131, -94, -44, 131, 128, 21, -44, 128, -94, 183, -94, 131, 139, -44, 128, 21, 181, -44, 131, 128, 131, 21, 68, 181, -44, -126, -130, 131, -190, 131, 181, 165, -94, 165, 165, -30, -154, 68, -39, -44, 165, -39, -126, 68, 68, -130, 68, -94, 181, -44, 131, 21, 183, -44, 21, -39, -130, -39, 131, 21, 165, 165, -126, 165, -44, -94, 68, 68, -94, -126, -126, -30, 181, 165, 68, -44, -39, -94, -126, -126, -30, 68, 181, -44, -94, -126, -44, -94, -30, 131, 165, -190, -130, -94, -94, 181, 128, 181, 181, 181, 139, -130, -94, -130, -130, 139, -130, -90, -154, 181, 165, -30, -154, 165, -190, 159, 165, 139, -126, -44, 131, -44, -190, -126, -130, -94, 128, -154, 68, -130, -130, 68, 21, -44, -30, -126, -126, 131, 159, -190, -126, 181, 139}, false

	res := uniqueOccurrences(arr)
	fmt.Printf("expected %v, res %v \n", expected, res)
}
