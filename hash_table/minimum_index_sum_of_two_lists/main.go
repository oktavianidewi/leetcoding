package main

func findRestaurant_dewi(list1 []string, list2 []string) []string {
	// swap bigger list to list1
	if len(list2) > len(list1) {
		list1, list2 = list2, list1
	}

	// create hash for the biggest list
	hash1 := make(map[string]int)
	for i := 0; i <= len(list1)-1; i++ {
		// increment 1 to i (to avoid 0 value. unmatched pair results in 0)
		hash1[list1[i]] = i + 1
	}
	// fmt.Printf("all hash %v \n", hash1)
	// fmt.Printf("find hash %v \n", hash1["Shogun"])

	// match list to the hash
	minIndexSum := -1
	matches := []string{}
	for j := 0; j <= len(list2)-1; j++ {
		if _, ok := hash1[list2[j]]; ok {
			// fmt.Printf("found! %v \n", list2[j])
			// if there are matched index more than 1
			if minIndexSum != -1 && minIndexSum > j+(hash1[list2[j]]-1) {
				return []string{list2[j]}
			}
			minIndexSum = j + (hash1[list2[j]] - 1)
			matches = append(matches, list2[j])
		}
	}
	return matches
}

func findRestaurant(list1 []string, list2 []string) []string {
	// convert list to hash
	hash1 := make(map[string]int)
	for key, val := range list1 {
		hash1[val] = key
	}
	hash2 := make(map[string]int)
	for key, val := range list2 {
		hash2[val] = key
	}

	// fmt.Printf("hash1 %v \n", hash1)
	// fmt.Printf("hash2 %v \n", hash2)

	minIndexSum := -1
	result := []string{}
	for key, val := range hash1 {
		// if val hash1 matched with value in hash2
		// fmt.Printf("isi key %v \n", key)
		if _, ok := hash2[key]; ok {
			if minIndexSum < 0 {
				minIndexSum = val + hash2[key]
				result = append(result, key)
				// fmt.Printf("hasil result %v \n", result)
			} else {
				sumIndex := val + hash2[key]
				if minIndexSum > sumIndex {
					// remove prev value
					result = nil

					// add min index to the end
					result = append(result, key)
				}
			}
		}
	}
	// fmt.Printf("isi result %v \n", result)
	return result
}
