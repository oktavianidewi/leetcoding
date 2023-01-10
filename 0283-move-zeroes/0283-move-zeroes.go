func moveZeroes(nums []int) {
	if len(nums) > 1 {

		j := 1
		for i := 0; i < len(nums); {
			if nums[i] == 0 && nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			} else if nums[i] != 0 && nums[j] == 0 {
				i = j
			}
			j++
			// fmt.Printf(">> nums %v \n", nums)
			if j == len(nums) {
				break
			}
		}
	}
}