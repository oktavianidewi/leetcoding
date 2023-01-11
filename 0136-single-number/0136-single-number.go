func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = hash[nums[i]] + 1
		if hash[nums[i]] > 1 {
			delete(hash, nums[i])
		}
	}
	for k := range hash {
		return k
	}
	return -1
}