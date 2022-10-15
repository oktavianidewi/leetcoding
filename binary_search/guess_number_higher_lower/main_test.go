package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Guess(t *testing.T) {
	input := 4
	expected := 1

	answer := guess(input)
	fmt.Printf("test %v -- %v ", expected, answer)
	assert.Equal(t, expected, answer)
}

func Test_Guessnumber(t *testing.T) {
	input := 14
	expected := 6

	answer := guessNumber(input)
	assert.Equal(t, expected, answer)
}
