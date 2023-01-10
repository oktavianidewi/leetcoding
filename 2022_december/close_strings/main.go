package main

import (
	"fmt"
	"reflect"
	"sort"
)

func strToMap(word string) map[rune]int {
	result := make(map[rune]int)
	for _, y := range word {
		result[y]++
	}
	return result
}
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	map1 := strToMap(word1)
	map2 := strToMap(word2)

	// fmt.Printf("map1 %v \n", map1)
	// fmt.Printf("map2 %v \n", map2)

	if len(map1) != len(map2) {
		return false
	}

	var vals1, vals2 []int
	for key1, val1 := range map1 {
		val2, matchKey := map2[key1]
		if !matchKey {
			return false
		}
		vals1 = append(vals1, val1)
		vals2 = append(vals2, val2)
	}
	sort.Ints(vals1)
	sort.Ints(vals2)
	return reflect.DeepEqual(vals1, vals2)

	// return true
}

func main() {
	// word1, word2 := "abc", "bca"
	// expected := true

	// word1, word2 := "a", "aa"
	// expected := false

	word1, word2 := "cabbba", "abbccc"
	// word1, word2 := "cabbba", "abbggg"
	expected := true

	// word1, word2 := "abbzzca", "babzzcz"
	// expected := false

	result := closeStrings(word1, word2)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
