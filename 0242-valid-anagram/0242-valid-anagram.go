func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	maps := make(map[rune]int)
	mapt := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		maps[rune(s[i])]++
		mapt[rune(t[i])]++
	}

	for key, val := range maps {
		valt, ok := mapt[key]
		if !ok || val != valt {
			return false
		}
	}
	return true
}