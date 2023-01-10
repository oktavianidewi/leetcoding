package main

import "testing"

// create test cases
func TestIntersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	expected := []int{2}
	actual := intersection(nums1, nums2)
	if len(actual) != len(expected) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestIntersection_2(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	expected := []int{4, 9}
	actual := intersection(nums1, nums2)
	if len(actual) != len(expected) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
