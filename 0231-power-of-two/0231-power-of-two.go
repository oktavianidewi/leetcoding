func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	var sisa int
	for n > 2 {
		if sisa < n%2 {
			sisa = n % 2
		}
		n = n >> 1
	}

	return sisa <= 0
}