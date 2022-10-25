func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
		// fmt.Printf("left %v, right %v, mid %v \n", left, right, mid)
		// time.Sleep(1 * time.Second)
	}
	if val := nums[left]; target > val {
		return left + 1
	}
	return left
}