package main

import (
	"fmt"
	"time"
)

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	left, mid, right := 0, 0, num

	for left <= right {
		mid = left + (right-left)/2
		// fmt.Printf("left, mid, right %v, %v, %v \n", left, mid, right)

		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid
			if (left == mid) && (left == right) {
				return false
			}
		} else {
			left = mid + 1
			if (left == mid) && (left == right) {
				return false
			}
		}
		// time.Sleep(1 * time.Second)
	}
	return false
}

func isPerfectSquare_2(num int) bool {
	left, mid, right := 0, 0, num
	for left <= right {
		if num == 1 {
			mid = num
		} else {
			mid = (left + right) / 2
		}
		fmt.Printf("left, mid, right %v, %v, %v \n", left, mid, right)
		// rumus square = sisi mid * sisi mid
		// fmt.Printf("mid %v, %v, %v \n", mid*mid, mid*mid == num, num)

		if mid*mid == num {
			return true
		} else if mid*mid > num {
			right = mid
		} else {
			return false
		}
		time.Sleep(1 * time.Second)
	}
	return false
}
