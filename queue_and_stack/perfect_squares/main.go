package main

import (
	"fmt"
	"time"
)

// source: https://leetcode.com/explore/learn/card/queue-stack/231/practical-application-queue/1371/discuss/707927/golang-Recursive-and-BFS-(two-solutions)
func numSquares(n int) int {
	// Creating visited vector of size n + 1
	visited := []int{}

	// Queue of pair to store node and number of steps
	// Push starting node with 0; 0 indicate current number of step to reach n
	q := []int{n}

	// Initially ans variable is initialized with inf
	ans := 0

	// Mark starting node visited
	return ans
}

func numSquares_solution1(n int) int {
	q := []int{n}
	steps := 1
	for len(q) != 0 {
		popped := len(q)
		fmt.Printf("popped %v \n", popped)
		for _, n := range q {
			for i := 1; i*i <= n; i++ {
				if n-(i*i) == 0 {
					return steps
				}
				q = append(q, n-(i*i))
				fmt.Printf("queue %v \n", q)
				time.Sleep(1 * time.Second)
			}
		}
		// queue sudah habis dan belum memenuhi kondisi n-(i*i)
		steps++
		q = q[popped:]
		fmt.Printf("steps %v \n", steps)
		fmt.Printf("start queue %v \n", q)
	}
	return -1
}

// explanation
// https://www.youtube.com/watch?v=wZqW204FC6k
// https://www.youtube.com/watch?v=1xfx6M_GqFk

// https://www.geeksforgeeks.org/minimum-number-of-squares-whose-sum-equals-to-given-number-n/
// https://medium.com/nerd-for-tech/perfect-squares-python-program-b82fff87916f

func main() {
	n := 12
	expected := 3
	// ada 3 perfect square dari 12: 4, 4, 4
	answer := numSquares(n)

	// n := 13
	// expected := 2
	// ada 2 perfect square dari 13: 4, 9

	fmt.Printf("expected %v , answer %v \n", expected, answer)
}
