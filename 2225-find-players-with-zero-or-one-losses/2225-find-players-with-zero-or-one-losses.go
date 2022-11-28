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