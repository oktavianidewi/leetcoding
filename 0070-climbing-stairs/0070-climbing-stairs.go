func climbStairs(n int) int {
	steps := _climbStairs(n)
	return steps
}

var hashSteps = make(map[int]int)

func _climbStairs(n int) int {
	if n == 0 || n == 1 {
		hashSteps[n] = 1
		return hashSteps[n]
	}
	if _, found := hashSteps[n]; !found {
		hashSteps[n] = _climbStairs(n-1) + _climbStairs(n-2)
	}
	return hashSteps[n]
}
