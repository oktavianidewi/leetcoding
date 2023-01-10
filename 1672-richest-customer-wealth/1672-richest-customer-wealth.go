func maximumWealth(accounts [][]int) int {
	max := 0
	li := len(accounts)
	for i := 0; i < li; i++ {
		curr := 0
		lj := len(accounts[i])
		for j := 0; j < lj; j++ {
			curr += accounts[i][j]
		}
		if curr > max {
			max = curr
		}
	}
	return max
}