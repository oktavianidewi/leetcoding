package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	mapNums := make(map[int]int)
	for _, num := range nums {
		mapNums[num] += 1
	}

	// count via combinations
	var res int
	for _, val := range mapNums {
		var temp int
		if val > 1 {
			temp = val * (val - 1) / 2
		}
		res += temp
	}
	return res
}

func main() {
	nums, exp := []int{2, 2, 1, 5, 1, 5, 5, 2, 3, 1, 1, 5, 3, 2, 3, 3, 3, 1, 3, 3, 4, 3, 1, 3, 1, 4, 5, 5, 2, 2, 1, 3, 5, 2, 2, 4, 3, 2, 5, 3, 1, 1, 3, 3, 2, 5, 2, 1, 2, 4, 3, 4, 4, 3, 2, 4, 4, 1, 3, 3, 3, 5, 5, 5, 4, 1, 1, 2, 3, 3, 2, 5, 3, 4, 5, 3, 1, 2, 5, 4, 5, 2, 3, 3, 1, 5, 2, 4, 2, 4, 4, 3, 1, 3}, 885
	// nums, exp := []int{1, 2, 3, 1, 1, 3}, 4
	// nums, exp := []int{1, 1, 1, 1}, 6
	// nums, exp := []int{1, 2, 3}, 0
	res := numIdenticalPairs(nums)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
