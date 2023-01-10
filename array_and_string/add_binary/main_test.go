package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addBinary_1(t *testing.T) {
	// t.Skip()
	a := "11"
	b := "1"
	expected := "100"
	answer := addBinary(a, b)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_addBinary_2(t *testing.T) {
	// t.Skip()
	a := "1010"
	b := "1011"
	expected := "10101"
	answer := addBinary(a, b)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}
