// https://cloud.tencent.com/developer/article/1418136

package main

func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	max, sec := 0, 0
	p := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			sec = max
			max = nums[i]
			p = i
		} else if nums[i] > sec {
			sec = nums[i]
		}
	}
	if max > 2*sec {
		return p
	}
	return -1
}
