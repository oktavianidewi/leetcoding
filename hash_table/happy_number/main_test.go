package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// create test cases
func Test_isHappy(t *testing.T) {
	n := 19
	expected := true
	answer := isHappy(n)
	assert.Equal(t, expected, answer)
}

func Test_isHappy_2(t *testing.T) {
	n := 2
	expected := false
	answer := isHappy(n)
	assert.Equal(t, expected, answer)
}
