package main

import (
	"fmt"
	"math"
)

func findNumbers(nums []int) int {
	counter := 0
	var i float64
	left, right := 0, len(nums)-1
	for left <= right {
		i = 0
		for {
			if (int32(i)+1)%2 == 0 {
				base := nums[left] / int(math.Pow(10, i))
				fmt.Printf("division %v dibagi %v = %v \n", nums[left], math.Pow(10, i), base)

				if base > 0 && base < 10 {
					counter++
					fmt.Printf("counter %v \n", counter)
					break
				} else if base <= 0 {
					break
				}
			}
			i += 1
			// time.Sleep(1 * time.Second)
		}
		left += 1
	}
	return counter
}
