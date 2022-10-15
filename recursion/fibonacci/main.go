package main

func fib(n int) int {
	return _fib(n)
}

var hashFib = make(map[int]int)

func _fib(n int) int {
	if n == 0 || n == 1 {
		hashFib[n] = n
		return n
	}
	// hash
	if _, found := hashFib[n]; !found {
		hashFib[n] = _fib(n-1) + _fib(n-2)
	}
	return hashFib[n]
}

// recursion dengan tail dan non-tail
