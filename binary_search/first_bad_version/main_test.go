package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Firstbadversion(t *testing.T) {
	input := 5
	expected := 4

	answer := firstBadVersion(input)
	assert.Equal(t, expected, answer)
}
