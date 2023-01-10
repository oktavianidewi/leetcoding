package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findRestaurant_0(t *testing.T) {
	// t.Skip()
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	expected := []string{"Shogun"}
	answer := findRestaurant(list1, list2)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_findRestaurant_1(t *testing.T) {
	// t.Skip()
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"KFC", "Shogun", "Burger King"}
	expected := []string{"Shogun"}
	answer := findRestaurant(list1, list2)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_findRestaurant_2(t *testing.T) {
	// t.Skip()
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Subway"}
	expected := []string{}
	answer := findRestaurant(list1, list2)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}
