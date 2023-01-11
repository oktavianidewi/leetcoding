func missingNumber(nums []int) int {
	return factorial(len(nums)) - sumAll(nums)
}

func sumAll(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total = total + nums[i]
	}
	return total
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num + factorial(num-1)
}
