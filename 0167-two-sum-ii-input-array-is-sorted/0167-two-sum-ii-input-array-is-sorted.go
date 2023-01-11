func twoSum(numbers []int, target int) []int {
	var indexes []int
	for i := 0; i < len(numbers); i++ {
		delta := target - numbers[i]

		// binary search
		left := i + 1
		right := len(numbers)
		for left < right {
			// fmt.Printf("compare index i %v and j %v \n", i, left)

			mid := (left + right) / 2
			if delta < numbers[mid] {
				right = mid
			} else if delta > numbers[mid] {
				left = mid + 1
			} else if delta == numbers[mid] {
				return append(indexes, i+1, mid+1)
			}

			// time.Sleep(1 * time.Second)
		}
	}
	return indexes
}