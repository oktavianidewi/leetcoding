package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var result string
	hash := make(map[int]string)
	return result
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	output := "fl"
	// strs := []string{"dog", "racecar", "car"}
	// output := ""
	result := longestCommonPrefix(strs)
	fmt.Printf("expected %v, result %v \n", output, result)
}
