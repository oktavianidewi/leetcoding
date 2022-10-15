package main

func replaceElements(arr []int) []int {
	var max int
	left, right := 0, len(arr)-1
	for left <= right {
		if left == right {
			max = -1
		}
		max = getMax(arr[left+1:])
		arr[left] = max
		left++
	}
	return arr
}

func getMax(nums []int) int {
	pointer := 0
	moving := pointer + 1
	max_value := 0

	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		return nums[0]
	}

	for moving <= len(nums)-1 {
		if nums[pointer] > nums[moving] {
			max_value = nums[pointer]
		} else {
			max_value = nums[moving]
			pointer = moving
		}
		moving++
	}
	// fmt.Printf("nums, max %v, %v \n", nums, max_value)
	return max_value
}
