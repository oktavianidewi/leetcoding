package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) > 1 {

		j := 1
		for i := 0; i < len(nums); {
			if nums[i] == 0 && nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			} else if nums[i] != 0 && nums[j] == 0 {
				i = j
			}
			j++
			fmt.Printf(">> nums %v \n", nums)
			if j == len(nums) {
				break
			}
		}
	}
}

func main() {
	// nums := []int{0, 1, 0, 3, 12}
	// 0, 1, 0, 3, 12
	// 1, 0, 0, 3, 12
	// 1, 3, 0, 0, 12
	// 1, 3, 12, 0, 0
	// result := []int{1, 3, 12, 0, 0}

	nums := []int{1, 0, 1}
	result := []int{1, 1, 0}

	// nums := []int{0}
	// result := []int{0}

	moveZeroes(nums)
	fmt.Printf("expected %v,result %v \n", result, nums)

}
