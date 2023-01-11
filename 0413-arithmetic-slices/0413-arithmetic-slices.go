func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	counter := 0
	consecutive := 1
	prevDiff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		currDiff := nums[i] - nums[i-1]
		if currDiff == prevDiff {
			counter += consecutive
			consecutive++
		} else {
			prevDiff = currDiff
			consecutive = 1
		}
	}
	return counter
}