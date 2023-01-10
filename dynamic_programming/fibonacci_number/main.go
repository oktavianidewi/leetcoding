package main

import "fmt"

func fib(n int) int {
	return _fib(n)
}

func _fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return _fib(n-1) + _fib(n-2)
}

func main() {
	fmt.Printf("hasil %v ", fib(4))
}
