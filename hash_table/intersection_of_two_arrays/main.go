package main

func intersection(nums1 []int, nums2 []int) []int {
	// find the intersection of two arrays
	// use a map to store the elements in nums1
	// use a slice to store the intersection
	// iterate through nums2
	// if the element is in the map, append it to the slice
	// if the element is not in the map, add it to the map
	// return the slice

	// swap the arrays if nums1 is longer than nums2
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := make(map[int]bool)
	for _, v := range nums1 {
		m[v] = true
	}
	result := make(map[int]int)
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			// check if the element is in the map
			result[v] = 1
		}
	}

	// convert the map to a slice
	var resultSlice []int
	for k, _ := range result {
		resultSlice = append(resultSlice, k)
	}
	return resultSlice
}
