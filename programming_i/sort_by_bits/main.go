package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	// rearrange from hash to array
	sort.Slice(arr, func(a, b int) bool {
		c := countOneBits(arr[a])
		d := countOneBits(arr[b])
		if c == d {
			return arr[a] < arr[b]
		}
		return c < d
	})
	return arr
}
func countOneBits(num int) int {
	if num == 0 {
		return 0
	}
	remainder, counter := 0, 0
	for num > 0 {
		remainder = num % 2
		num = num / 2
		if remainder == 1 {
			counter++
		}
	}
	return counter
}

func main() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	output := []int{0, 1, 2, 4, 8, 3, 5, 6, 7}

	// input := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	// output := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	result := sortByBits(input)
	fmt.Printf("expected %v, result %v \n", output, result)
}
