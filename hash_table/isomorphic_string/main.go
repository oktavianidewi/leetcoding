package main

func isIsomorphic(s string, t string) bool {
	// use hashkey
	if len(s) != len(t) {
		return false
	}

	// rune is an alias to the int32 data type
	hash_s := make(map[rune]int)
	hash_t := make(map[rune]int)
	for k, v := range s {
		if _, ok := hash_s[v]; ok {
			if hash_s[v] != int(t[k]) {
				return false
			}
		}
		hash_s[v] = int(t[k])

		// failed at use case: "kalo kebalikannya gimana?"
		// cek kebalikannya
		if _, ok := hash_t[rune(t[k])]; ok {
			if hash_t[rune(t[k])] != int(v) {
				return false
			}
		}
		hash_t[rune(t[k])] = int(v)

	}

	return true
}
