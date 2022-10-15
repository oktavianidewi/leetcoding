package main

var badVersion = 4

func isBadVersion(n int) bool {
	if n == badVersion {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	left, mid, right := 1, 0, n
	for left <= right {
		mid = (left + right) / 2
		if isBadVersion(mid) {
			if !isBadVersion(mid - 1) {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
