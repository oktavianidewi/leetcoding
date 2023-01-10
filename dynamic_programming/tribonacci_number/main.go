package main

import "fmt"

func tribonacci(n int) int {
	hashMap := make(map[int]int)
	return _tribonacci(n, hashMap)
}

func _tribonacci(n int, hashMap map[int]int) int {
	if n <= 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	if val, ok := hashMap[n]; ok {
		return val
	}
	hashMap[n] = _tribonacci(n-1, hashMap) + _tribonacci(n-2, hashMap) + _tribonacci(n-3, hashMap)

	return hashMap[n]
}

// tambahi hashMap supaya gak lama prosesnya

func main() {
	fmt.Printf("hasil %v \n", tribonacci(35))
}
