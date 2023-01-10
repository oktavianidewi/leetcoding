package main

func groupAnagrams(strs []string) [][]string {
	// create map of array string, with key array of int with length of 26
	mp := map[[26]int][]string{}

	for _, s := range strs {
		// array of int with length of 26
		k := [26]int{}
		for i := 0; i < len(s); i++ {
			// fmt.Printf("ini apa %v, %v, %v \n", s[i], 'a', s[i]-'a')
			// put value 1 to its position in array of char, based on its delta int value of char
			k[s[i]-'a'] += 1
			// fmt.Printf("k %v \n", k)
			// time.Sleep(1 * time.Second)
		}
		// grouping the same pattern of array k
		mp[k] = append(mp[k], s)
		// fmt.Printf("isi mp after append %v \n", mp)
	}

	// finalize: move from an array with key of binary to key of integer
	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}
