func uniqueOccurrences(arr []int) bool {
	uniqueMap := make(map[int]int)
	for _, val := range arr {
		uniqueMap[val]++
	}
	// fmt.Printf("uniqueMap %v \n", uniqueMap)

	// compare
	uniqueCount := make(map[int]int)
	for _, count := range uniqueMap {
		if _, found := uniqueCount[count]; found {
			return false
		}
		uniqueCount[count] = 1
	}
	return true
}