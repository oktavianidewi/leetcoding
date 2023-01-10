package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	n      int
	output int
}{
	{
		n:      0,
		output: 0,
	},
	{
		n:      1,
		output: 1,
	},
	{
		n:      2,
		output: 1,
	},
	{
		n:      3,
		output: 2,
	},
	{
		n:      4,
		output: 3,
	},
	{
		n:      5,
		output: 5,
	},
	{
		n:      6,
		output: 8,
	},
	{
		n:      7,
		output: 13,
	},
}

func TestFibonacci(t *testing.T) {
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		outputResult := fib(v.n)
		// fmt.Printf("isi output %v , expected %v > \n", outputResult, v.output)
		assert.Equal(t, outputResult, v.output)
	}
}
