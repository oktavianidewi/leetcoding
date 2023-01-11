func mergeAlternately(word1 string, word2 string) string {
	var extraString string
	var smallestLength int

	if len(word2) < len(word1) {
		smallestLength = len(word2)
		extraString = word1[len(word2):]
	} else {
		smallestLength = len(word1)
		extraString = word2[len(word1):]
	}

	var res []string

	for i := 0; i < smallestLength; i++ {
		res = append(res, string(word1[i]), string(word2[i]))
	}

	return strings.Join(res, "") + extraString
}