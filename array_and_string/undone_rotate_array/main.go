package main

func rotate(nums []int, k int) {

	if len(nums)%2 == 0 {
		// genap
		i, j := 0, k
		for i < k && j < len(nums) {
			nums[i], nums[j] = nums[j], nums[i]
			j++
			i++
		}

		// uncovered edge case:
		// if k > len(nums) {
		// }

	} else {
		// ganjil
		i, j := 0, k+1
		for i < k && j < len(nums) {
			nums[i], nums[j] = nums[j], nums[i]
			j++
			i++
		}
		// shift pivot to the end
		l := k + 1
		for k < len(nums) && l < len(nums) {
			nums[k], nums[l] = nums[l], nums[k]
			k++
			l++
		}
	}
}
