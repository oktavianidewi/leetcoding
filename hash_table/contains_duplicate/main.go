package main

import "fmt"

func containsDuplicate(nums []int) bool {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		duplicateCounter := 1
		if _, ok := hash[nums[i]]; ok {
			duplicateCounter = hash[nums[i]]
			duplicateCounter++
		}
		hash[nums[i]] = duplicateCounter
	}

	fmt.Printf("isi hash %v , %v \n", len(nums), hash)

	if len(nums) > len(hash) {
		return true
	}

	return false
}
