package main

import "fmt"

func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	for n > 0 {
		sum = sum + n%10
		product = product * (n % 10)
		// n % 10 to get the least significant bit on each loop
		// decreament least significant bit
		n = n / 10
	}
	return product - sum
}

func main() {
	n := 234
	expAns := 15

	// n := 4221
	// expAns := 21

	ans := subtractProductAndSum(n)
	fmt.Printf("expected %v, answer %v \n", expAns, ans)
}
