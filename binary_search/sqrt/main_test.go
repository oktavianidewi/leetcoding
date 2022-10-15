package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MySqrt(t *testing.T) {
	input := 36
	expected := 6

	mySqrt := mySqrt(input)
	assert.Equal(t, expected, mySqrt)
}
