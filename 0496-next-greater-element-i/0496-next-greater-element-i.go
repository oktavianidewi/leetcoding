func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var result []int

	for i := 0; i < len(nums1); i++ {
		var found bool
		var greaterExists bool
		for j := 0; j < len(nums2); j++ {
			if found == true && nums2[j] > nums1[i] {
				greaterExists = true
				result = append(result, nums2[j])
				break
			}
			if nums1[i] == nums2[j] {
				found = true
			}
		}
		if greaterExists == false {
			result = append(result, -1)
		}
	}

	return result
}