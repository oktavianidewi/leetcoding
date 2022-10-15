package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CaseTrue(t *testing.T) {

	input := 16
	expected := true

	answer := isPerfectSquare(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseFalse(t *testing.T) {

	input := 14
	expected := false

	answer := isPerfectSquare(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseTrue_Edge1(t *testing.T) {

	input := 1
	expected := true

	answer := isPerfectSquare(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseTrue_Edge9(t *testing.T) {
	input := 9
	expected := true

	answer := isPerfectSquare(input)
	assert.Equal(t, expected, answer)
}
