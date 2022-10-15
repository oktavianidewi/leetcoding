package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	prev := arr[1] - arr[0]

	for i := 1; i < len(arr)-1; i++ {
		if prev != arr[i+1]-arr[i] {
			return false
		}

		prev = arr[i+1] - arr[i]
	}
	return true
}

func canMakeArithmeticProgression_dewi(arr []int) bool {
	sort.Ints(arr)
	divisions := make(map[int]int)

	for i := 0; i < len(arr)-1; i++ {
		div := arr[i] - arr[i+1]
		divisions[div]++
	}

	// fmt.Printf("array %v, divisions %v \n", arr, divisions)

	return len(divisions) == 1
}

func main() {
	// nums := []int{3, 5, 1}
	// exp := true

	nums := []int{1, 2, 4}
	exp := false

	ans := canMakeArithmeticProgression(nums)
	fmt.Printf("expected %v , ans %v \n", exp, ans)
}
