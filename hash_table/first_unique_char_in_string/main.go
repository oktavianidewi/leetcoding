package main

func firstUniqChar(s string) int {
	// masalahnya map di golang nggak urut, harus pake apa ya?

	dupChar := make(map[string]int)
	oriIndexChar := make(map[string]int, len(s))
	for i := 0; i <= len(s)-1; i++ {
		// store minimum index
		oriIndexChar[string(s[i])] = i
		dupChar[string(s[i])] = dupChar[string(s[i])] + 1
	}
	// fmt.Printf("oriIndexChar %v \n", oriIndexChar)
	// fmt.Printf("dupChar %v \n", dupChar)

	minIndex := -1
	for k, v := range dupChar {
		if v == 1 {
			if minIndex == -1 {
				minIndex = oriIndexChar[k]
			} else if minIndex > -1 && oriIndexChar[k] < minIndex {
				minIndex = oriIndexChar[k]
			}
		}
	}
	return minIndex
}
