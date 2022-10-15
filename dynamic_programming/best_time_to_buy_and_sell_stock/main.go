package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {

	minPrice := math.MaxInt32
	maxPft := 0

	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if v-minPrice > maxPft {
			maxPft = v - minPrice
		}
	}

	return maxPft

}

func maxProfit_dewi(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	var delta, maxIdx int
	maxIdx = 1

	for minIdx := 0; minIdx < len(prices)-1; {
		// time.Sleep(1 * time.Second)
		// fmt.Printf("maxVal %v, minVal %v, >>> delta %v\n", prices[maxIdx], prices[minIdx], prices[maxIdx]-prices[minIdx])

		if prices[maxIdx]-prices[minIdx] <= 0 {
			minIdx++
		} else {
			delta = max(delta, prices[maxIdx]-prices[minIdx])
		}

		if minIdx < len(prices)-1 && maxIdx == len(prices)-1 {
			minIdx++
		} else {
			maxIdx++
		}
	}
	// fmt.Printf("delta %v \n", delta)
	return delta
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	// prices := []int{14, 21, 9, 1, 17}
	// prices := []int{25, 21, 9, 8, 5}
	// prices := []int{7, 1, 5, 3, 6, 4}
	// prices := []int{7, 6, 4, 3, 1}
	// prices := []int{1}
	// prices := []int{1, 7}
	// prices := []int{7, 2}
	// prices := []int{1, 2, 4}
	// prices := []int{4, 2, 1}
	// prices := []int{1, 4, 2}
	// prices := []int{6, 4, 2, 1}
	// prices := []int{4, 2, 6, 1}
	// prices := []int{4, 2, 1, 6}
	// prices := []int{4, 6, 2, 1}
	prices := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9}

	result := maxProfit(prices)
	fmt.Printf("result %v\n", result)
}
