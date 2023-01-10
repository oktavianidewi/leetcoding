package main

var pick = 6

func guess(num int) int {
	var res int
	if num > pick {
		res = -1
	} else if num < pick {
		res = 1
	}
	return res
}

func guessNumber(n int) int {
	// return guess(n)
	left, mid, right := 0, 0, n
	for left <= right {
		mid = (left + right) / 2
		if guess(mid) > 0 {
			left = mid + 1
		} else if guess(mid) < 0 {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
