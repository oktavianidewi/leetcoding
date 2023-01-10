package main

import "fmt"

// https://leetcode.com/problems/sum-of-all-odd-length-subarrays/discuss/1262789/Golang-with-Slice
func sum(arr []int) int {
	ret := 0
	for _, elem := range arr {
		ret += elem
	}
	return ret
}

func sumOddLengthSubarrays(arr []int) int {
	ret := 0
	for currentLen := 1; currentLen <= len(arr); currentLen += 2 {
		for i := range arr {
			if i+currentLen > len(arr) {
				continue
			}
			ret += sum(arr[i : i+currentLen])
		}
	}
	return ret
}

func main() {
	arr := []int{1, 4, 2, 5, 3}
	output := 58

	// arr := []int{1, 2}
	// output := 3

	// arr := []int{10, 11, 12}
	// output := 66
	// 10+11+12+(10+11+12)

	result := sumOddLengthSubarrays(arr)
	fmt.Printf("expected %v, result %v \n", output, result)
}
