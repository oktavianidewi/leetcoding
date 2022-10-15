package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	i := 0
	j := 0
	right := len(numbers) - 1
	res := make([]int, 2)
	delta := -1
	for i < right {
		if target > 0 {
			delta = target - numbers[i]
		} else {
			delta = -1*(target) + numbers[i]
		}

		for j <= right {
			if target > 0 {
				fmt.Printf("compute + %v, %v \n", delta, numbers[j])
				if numbers[j] == delta {
					delta -= numbers[j]
					res[0] = i + 1
					res[1] = j + 1
				}
			} else {
				fmt.Printf("compute - %v, %v \n", delta, numbers[j])
				if numbers[j] == delta {
					delta += numbers[j]
					res[0] = i + 1
					res[1] = j + 1
				}
			}

			j++

			if delta < 0 {
				break
			}
		}
		i++
	}
	return res
}
