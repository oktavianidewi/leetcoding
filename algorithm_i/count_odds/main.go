package main

import "fmt"

func countOdds(low int, high int) int {
	var ans int
	if low%2 > 0 || high%2 > 0 {
		ans = 1 + (high-low)/2
	} else {
		ans = (high - low) / 2
	}
	return ans
}

func main() {
	a := countOdds(3, 20)
	b := countOdds(3, 21)
	c := countOdds(4, 20)
	d := countOdds(4, 21)
	fmt.Printf("a %v, b %v, c %v, d %v \n", a, b, c, d)
}
