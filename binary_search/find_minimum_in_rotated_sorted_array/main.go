package main

func findMin(nums []int) int {
	left, mid, right := 0, 0, len(nums)-1
	if nums[left] < nums[right] {
		return nums[left]
	}

	for left < right {
		mid = (left + right) / 2

		// inflection point
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[left]
}
