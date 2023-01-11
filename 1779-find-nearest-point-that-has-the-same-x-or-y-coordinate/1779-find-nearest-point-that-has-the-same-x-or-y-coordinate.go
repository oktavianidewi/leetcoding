
func nearestValidPoint(x int, y int, points [][]int) int {
	// return index of the smallest
	minVal := math.MaxInt32
	idx := -1

	for i, point := range points {
		x1, y1 := point[0], point[1]
		if x == x1 && manhattanDistance(y, y1) < minVal {
			minVal = manhattanDistance(y, y1)
			idx = i
		} else if y == y1 && manhattanDistance(x, x1) < minVal {
			minVal = manhattanDistance(x, x1)
			idx = i
		}
	}
	return idx
}

func manhattanDistance(val, val1 int) int {
	return absVal(val - val1)
}

func absVal(n int) int {
	if n < 0 {
		return -n
	}
	return n
}