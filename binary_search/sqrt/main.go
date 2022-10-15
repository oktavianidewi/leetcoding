package main

import (
	"fmt"
	"time"
)

// contekan jawaban
func mySqrt(x int) int {
	left := 0
	right := x
	res := 0
	for left <= right {
		mid := (left + right) / 2
		fmt.Printf("%v-%v-%v-%v \n", left, mid, right, mid*mid)
		if mid*mid <= x {
			left = mid + 1
			res = mid
		} else {

			right = mid - 1
		}
		time.Sleep(1 * time.Second)
	}
	return res
}

func mySqrt_dewi(x int) int {
	// var nums []int
	// i := 0
	// for i < x {
	// 	nums = append(nums, i)
	// 	i += 1
	// }
	var validPowerX, startX, midX, endX float64
	// if x%2 == 1 {
	// 	x = x - 1
	// }

	startX = float64(0)
	endX = float64(x)

	for startX <= endX {
		// midX = startX + (endX / 2) // salah milih mid jadi berabe
		midX = (startX + endX) / 2
		validPowerX = midX * midX
		fmt.Printf("%v-%v-%v-%v \n", startX, midX, endX, validPowerX)
		if int(validPowerX) == x {
			return int(midX)
		} else if validPowerX > endX {
			// lower than mid search space
			endX = midX
		} else if validPowerX < endX {
			// greater than mid search space
			startX = midX
		}
		startX += 0.1
		time.Sleep(1 * time.Second)
	}
	return int(midX)
}
