package main

func findPeakElement(nums []int) int {
	left, mid, right := 0, 0, len(nums)-1
	for left < right {
		mid = (left + right) / 2
		// fmt.Printf("left, mid, right %v -- %v -- %v \n", left, mid, right)
		if nums[mid] > nums[mid+1] {
			if mid == right {
				return mid
			}
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
