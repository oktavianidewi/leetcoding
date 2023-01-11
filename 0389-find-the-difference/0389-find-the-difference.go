func findTheDifference(s string, t string) byte {
	var sumT rune
	for _, char := range t {
		sumT += char
	}

	for _, keyword := range s {
		sumT -= keyword
	}
	// fmt.Printf("result %v \n", string(sumT))
	return byte(sumT)
}
