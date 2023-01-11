func sortByBits(arr []int) []int {
	// rearrange from hash to array
	sort.Slice(arr, func(a, b int) bool {
		c := countOneBits(arr[a])
		d := countOneBits(arr[b])
		if c == d {
			return arr[a] < arr[b]
		}
		return c < d
	})
	return arr
}
func countOneBits(num int) int {
	if num == 0 {
		return 0
	}
	remainder, counter := 0, 0
	for num > 0 {
		remainder = num % 2
		num = num / 2
		if remainder == 1 {
			counter++
		}
	}
	return counter
}
