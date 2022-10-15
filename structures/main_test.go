package main

import (
	"fmt"
	"testing"
)

type params struct {
	param []int
}

type answers struct {
	answer bool
}

type positions struct {
	position int
}

type question struct {
	params
	positions
	answers
}

func Test_HashCycle(t testing.T) {
	qs := []question{
		{
			params{[]int{3, 4, 0, -4}},
			positions{1},
			answers{true},
		},
		{
			params{[]int{1, 2}},
			positions{0},
			answers{true},
		},
		{
			params{[]int{1}},
			positions{-1},
			answers{false},
		},
	}

	for _, q := range qs {
		fmt.Printf("input %v, output %v \n", q.answer, ConvertArrToLinkedList(q.param, q.position))
	}
}
