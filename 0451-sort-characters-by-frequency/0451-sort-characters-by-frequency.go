func frequencySort(s string) string {
	// create a slice of integers with length of 255
	freq := make([]int, 255)
	// store the number of occourance of each char. the char represented by its integer value.
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
		// fmt.Printf("byte of char %v, char %v \n", s[i], string(s[i]))
	}
	// fmt.Printf("freq %v \n", freq)
	b := []byte(s)
	// fmt.Printf("b before %v \n", b)

	// we want to sort b, based on the freq slice
	sort.Slice(b, func(i, j int) bool {

		if freq[b[i]] == freq[b[j]] {
			return b[i] > b[j]
		}
		return freq[b[i]] > freq[b[j]]
	})
	return string(b)
}