package main

import (
	"fmt"
	"sort"
)

func maxIceCream(costs []int, coins int) int {
	if len(costs) == 1 && coins < costs[0] {
		return 0
	} else if len(costs) == 1 && coins > costs[0] {
		return 1
	}

	// bisa di-simplekan dg cara:
	/*
		    var curr_cost, count int

		    for _ , cost := range costs {
		        curr_cost += cost
		        if curr_cost > coins {
		            return count
		        }
		        count++
		    }
		    return count
		}

	*/
	sort.Ints(costs)
	sum := 0
	for i := 0; i < len(costs)-1; i++ {
		// fmt.Println(sum, coins, i)
		sum += costs[i]
		if coins < sum {
			return 0
		} else if sum == coins || sum+costs[i+1] > coins {
			return i + 1
		}
	}
	// fmt.Printf("coins %v, sum %v \n", coins, sum)
	if coins >= sum {
		return len(costs)
	}
	return sum
}

func main() {
	// costs, coins, exp := []int{1, 3, 2, 4, 1}, 7, 4
	// costs, coins, exp := []int{10, 6, 8, 7, 7, 8}, 5, 0
	// costs, coins, exp := []int{1, 6, 3, 1, 2, 5}, 20, 6
	// costs, coins, exp := []int{7, 3, 3, 6, 6, 6, 10, 5, 9, 2}, 56, 9
	costs, coins, exp := []int{11}, 1, 0
	res := maxIceCream(costs, coins)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
