package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	n      int
	x      float64
	output float64
}{
	// {
	// 	x:      2.0,
	// 	n:      10,
	// 	output: 1024.0,
	// },
	{
		x:      2.10,
		n:      3,
		output: 9.2610,
	},
	{
		x:      2.0,
		n:      -2,
		output: 0.25,
	},
}

func TestClimbStairs(t *testing.T) {
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		outputResult := myPow(v.x, v.n)
		// fmt.Printf("isi output %v , expected %v > \n", outputResult, v.output)
		assert.Equal(t, v.output, outputResult)
	}
}
