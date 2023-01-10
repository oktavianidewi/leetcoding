package main

import "fmt"

func rotate(nums []int, k int) {
	end := len(nums)
	if k > end {
		k = k % end
	}

	// fmt.Printf("lebar array %v, rotate pivot after mod %v \n", end, k)

	j := 0
	for i := end - k; i < end; i++ {
		// fmt.Printf("nums i %v \n", nums[i])
		nums[j], nums[i] = nums[i], nums[j]
		j++
	}

	// sort the unrotated number
	var unsorted int
	if k == end {
		unsorted = 0
	} else if k > (end - k) {
		unsorted = (end - k) - (k % (end - k))
	} else {
		unsorted = end - (2 * k)
	}
	fmt.Printf("hasil nums %v , last j %v , unsorted %v \n", nums, j, unsorted)

	sortLoop := 0
	temp := j
	for sortLoop < unsorted && temp < end-1 {
		nums[temp], nums[temp+1] = nums[temp+1], nums[temp]
		temp++

		if temp == end-1 {
			temp = j
			sortLoop++
		}
	}
	// fmt.Printf("hasil nums %v , last j %v , unsorted %v \n", nums, j, unsorted)
}

func main() {
	nums, k, output := []int{1}, 1, []int{1}
	// nums, k, output := []int{1, 2}, 3, []int{2, 1}

	// nums, k, output := []int{1, 2, 3, 4, 5, 6, 7}, 1, []int{5, 6, 7, 1, 2, 3, 4}

	// nums, k, output := []int{-1}, 2, []int{-1}
	// nums, k, output := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53}, 82, []int{25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}

	// nums, k, output := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, 15, []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 1, 2, 3, 4}
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19
	// 5, 6, 7, 8, 1, 2, 3, 4, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19
	// 5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 3, 4, 13, 14, 15, 16, 17, 18, 19
	// 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 4, 1, 2, 3

	// 1, 2, 3, 4, 5, 6, 7
	// 5, 2, 3, 4, 1, 6, 7
	// 5, 6, 7, 4, 1, 2, 3
	// 5, 6, 7, 1, 4, 2, 3
	// 5, 6, 7, 1, 2, 4, 3
	// 5, 6, 7, 1, 2, 3, 4

	// rotate factor 2
	// 6, 7, 3, 4, 5, 1, 2
	// 6, 7, 4, 3, 5, 1, 2
	// 6, 7, 4, 5, 3, 1, 2
	// 6, 7, 4, 5, 1, 3, 2
	// 6, 7, 4, 5, 1, 2, 3

	rotate(nums, k)
	fmt.Printf("expected %v, in-place result %v \n", output, nums)
}
