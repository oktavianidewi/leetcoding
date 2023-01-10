package main

import (
	"fmt"
	"sort"
)

func findWinners_dewi(matches [][]int) [][]int {
	var result [][]int
	winnerArr := make([]int, 100000001)
	loserArr := make([]int, 100000001)
	for _, match := range matches {
		winnerArr[match[0]] += 1
		loserArr[match[1]] += 1
	}

	wins := []int{}
	loses := []int{}
	for i := 1; i < len(winnerArr); i++ {
		if loserArr[i] == 0 && winnerArr[i] != 0 {
			wins = append(wins, i)
		}
		if loserArr[i] == 1 {
			loses = append(loses, i)
		}
	}
	result = append(result, wins, loses)
	return result
}

func findWinners(matches [][]int) [][]int {
	var result [][]int
	winnerMap := make(map[int]int)
	loserMap := make(map[int]int)
	for _, match := range matches {
		winnerMap[match[0]] += 1
		loserMap[match[1]] += 1
	}

	wins := []int{}
	for i, win := range winnerMap {
		_, ok := loserMap[i]
		if !ok && win > 0 {
			wins = append(wins, i)
		}
	}
	sort.Ints(wins)

	loses := []int{}
	for i, lose := range loserMap {
		if lose == 1 {
			loses = append(loses, i)
		}
	}
	sort.Ints(loses)

	result = append(result, wins, loses)
	return result
}

func main() {
	matches := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}
	expected := [][]int{{1, 2, 10}, {4, 5, 7, 8}}

	// matches := [][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}
	// expected := [][]int{{1, 2, 5, 6}, {}}

	result := findWinners(matches)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
