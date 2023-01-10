package main

import (
	"fmt"
)

func permute_ori(nums []int, r int) [][]int {
	if r == 1 {
		//Convert every item in L to List and
		//Append it to List of List
		temp := make([][]int, 0)
		for _, rr := range nums {
			t := make([]int, 0)
			t = append(t, rr)
			temp = append(temp, [][]int{t}...)
		}
		return temp
	} else {
		res := make([][]int, 0)
		for i := 0; i < len(nums); i++ {
			//Create List Without L[i] element
			perms := make([]int, 0)

			// fmt.Printf("i %v \n", i)

			perms = append(perms, nums[:i]...)
			// fmt.Printf("first perms %v \n", perms)

			perms = append(perms, nums[i+1:]...)
			// fmt.Printf("second perms %v \n", perms)

			//Call recursively to Permutations
			for _, x := range permute_ori(perms, r-1) {
				t := append(x, nums[i])
				res = append(res, [][]int{t}...)
			}
		}
		return res
	}
}

func permute(nums []int) [][]int {
	r := len(nums)
	if r == 1 {
		//Convert every item in L to List and
		//Append it to List of List
		temp := make([][]int, 0)
		for _, rr := range nums {
			t := make([]int, 0)
			t = append(t, rr)
			temp = append(temp, [][]int{t}...)
		}
		return temp
	} else {
		res := make([][]int, 0)
		for i := 0; i < len(nums); i++ {
			//Create List Without L[i] element
			perms := make([]int, 0)

			fmt.Printf("i %v \n", i)

			perms = append(perms, nums[:i]...)
			fmt.Printf("first perms %v \n", perms)

			perms = append(perms, nums[i+1:]...)
			fmt.Printf("second perms %v \n", perms)

			//Call recursively to Permutations
			for _, x := range permute(perms) {
				fmt.Printf("x %v \n", x)
				t := append(x, nums[i])
				fmt.Printf("i %v, nums[i] %v \n", i, nums[i])

				res = append(res, [][]int{t}...)
				fmt.Printf("res %v \n", res)
			}
		}
		return res
	}
}
func main() {
	// permutation with no repetition
	nums := []int{1, 2, 3}
	output := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}

	res := permute(nums)
	fmt.Printf("expected %v, result %v \n", output, res)
}
