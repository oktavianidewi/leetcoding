package main

import "fmt"

// https://leetcode.com/problems/reverse-bits/discuss/426492/Golang-100-super-simple

// num&1 -> bisa untuk mendeteksi apakah sebuah angka ganjil atau genap
// num 65, hasil and 1
// num 32, hasil and 0
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 + num&1
		num >>= 1
	}
	return res
}

func main() {
	var num uint32
	num = 00000000000000000000000000001011
	// num = 00000010100101000001111010011100
	x := reverseBits(num)
	fmt.Printf("result %v \n", x)
}
