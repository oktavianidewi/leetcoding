package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://shareablecode.com/snippets/golang-solution-for-leetcode-problem-add-binary-dxhd-FPGR

func addBinary(a string, b string) string {

	if len(b) > len(a) {
		// menyeragamkan a > b
		a, b = b, a
	}
	res := make([]string, len(a)+1)

	i, j, k, carry := len(a)-1, len(b)-1, len(a), 0

	for i >= 0 && j >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		bj, _ := strconv.Atoi(string(b[j]))
		res[k] = strconv.Itoa((ai + bj + carry) % 2)
		carry = (ai + bj + carry) / 2

		fmt.Printf("ai, bj, carry %v, %v %v %v \n", ai, bj, res, carry)

		i--
		j--
		k--
	}

	// memproses sisa i
	fmt.Printf("sisa i %v \n", i)
	for i >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		res[k] = strconv.Itoa((ai + carry) % 2)
		carry = (ai + carry) / 2
		i--
		k--
	}

	if carry > 0 {
		res[k] = strconv.Itoa(carry)
	}

	return strings.Join(res, "")
}
