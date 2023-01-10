package main

import "testing"

// create test cases
func TestIntersection(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	expected := []int{2, 2}
	// karena 2 ada di index 1, 2

	actual := intersection(nums1, nums2)
	if len(actual) != len(expected) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestIntersection_2(t *testing.T) {
	// t.Skip()
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	expected := []int{4, 9}
	actual := intersection(nums1, nums2)
	if len(actual) != len(expected) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestIntersection_3(t *testing.T) {
	t.Skip()
	nums1 := []int{1, 2}
	nums2 := []int{1, 1}
	expected := []int{1}
	// karena 1 hanya ada di index 0
	actual := intersection(nums1, nums2)
	if len(actual) != len(expected) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
