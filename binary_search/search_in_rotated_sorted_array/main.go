// https://www.geeksforgeeks.org/search-an-element-in-a-sorted-and-pivoted-array/
package main

func search_1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2 //find the median
		if nums[mid] == target {
			return mid //determine whether to look in the first half
		}
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right])) { //determine which part the target is
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// fungsi ini nggak bisa handle test-case jika angka tidak ditemukan diarray
func search(nums []int, target int) int {
	// index
	if len(nums) == 0 {
		return -1
	}

	left, mid, right := 0, 0, len(nums)-1
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else if target < nums[right] {
				right = mid - 1
			}
		} else {
			if nums[left] == nums[mid] {
				left++
			} else {
				right--
			}
		}
	}
	return -1

	// for left <= right {
	// 	mid = (left + right) / 2
	// 	if nums[mid] == target {
	// 		return mid
	// 	} else if nums[mid] < target {
	// 		left = mid + 1
	// 	} else if nums[mid] > target {
	// 		right = mid - 1
	// 	}
	// }
	// return -1
}
