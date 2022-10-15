package main

func pivotIndex(nums []int) int {
	leftsum := 0
	totalSum := sumArr(nums)

	for i, v := range nums {
		if leftsum == totalSum-leftsum-v {
			return i
		}
		leftsum += v
	}
	return -1
}

func sumArr(nums []int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}
