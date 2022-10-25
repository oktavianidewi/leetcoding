// func rotate(nums []int, k int) {

// 	if len(nums)%2 == 0 {
// 		// genap
// 		i, j := 0, k
// 		for i < k && j < len(nums) {
// 			nums[i], nums[j] = nums[j], nums[i]
// 			j++
// 			i++
// 		}
// 	} else {
// 		// ganjil
// 		i, j := 0, k+1
// 		for i < k && j < len(nums) {
// 			nums[i], nums[j] = nums[j], nums[i]
// 			j++
// 			i++
// 		}
// 		// shift pivot to the end
// 		l := k + 1
// 		for k < len(nums) && l < len(nums) {
// 			nums[k], nums[l] = nums[l], nums[k]
// 			k++
// 			l++
// 		}
// 	}
// }

func rotate(nums []int, k int) {
	end := len(nums)
	if k > end {
		k = k % end
	}

	// fmt.Printf("lebar array %v, rotate pivot after mod %v \n", end, k)

	j := 0
	for i := end - k; i < end; i++ {
		// fmt.Printf("nums i %v \n", nums[i])
		nums[j], nums[i] = nums[i], nums[j]
		j++
	}

	// sort the unrotated number
	var unsorted int
	if k == end {
		unsorted = 0
	} else if k > (end - k) {
		unsorted = (end - k) - (k % (end - k))
	} else {
		unsorted = end - (2 * k)
	}
	// fmt.Printf("hasil nums %v , last j %v , unsorted %v \n", nums, j, unsorted)

	sortLoop := 0
	temp := j
	for sortLoop < unsorted && temp < end-1 {
		nums[temp], nums[temp+1] = nums[temp+1], nums[temp]
		temp++

		if temp == end-1 {
			temp = j
			sortLoop++
		}
	}
	// fmt.Printf("hasil nums %v , last j %v , unsorted %v \n", nums, j, unsorted)
}
