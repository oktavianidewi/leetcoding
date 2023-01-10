package main

func twoSum(numbers []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i <= len(numbers)-1; i++ {
		hash[numbers[i]] = i
	}

	result := []int{}

	for i := 0; i <= len(numbers)-1; i++ {
		operand := target - numbers[i]
		if _, ok := hash[operand]; ok && hash[operand] != i {
			result = append(result, i, hash[operand])
			return result
		}
	}
	return result
}
