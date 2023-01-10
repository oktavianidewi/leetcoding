package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	found := 0
	hash_jewels := make(map[rune]bool)
	for _, jewel := range jewels {
		hash_jewels[rune(jewel)] = true
	}
	fmt.Printf("hash_jewels %v \n", hash_jewels)

	for _, stone := range stones {
		if _, ok := hash_jewels[rune(stone)]; ok {
			found++
		}
	}
	return found
}
