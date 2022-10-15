package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// create test cases
func TestNumJewelsInStones(t *testing.T) {
	jewels := "aA"
	stones := "aAAbbbb"
	expected := 3

	actual := numJewelsInStones(jewels, stones)
	assert.Equal(t, expected, actual)
}

func TestNumJewelsInStones_2(t *testing.T) {
	jewels := "z"
	stones := "ZZ"
	expected := 0

	actual := numJewelsInStones(jewels, stones)
	assert.Equal(t, expected, actual)
}
