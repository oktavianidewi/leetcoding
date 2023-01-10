package main

import (
	"fmt"
	"time"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for k := m + n - 1; k >= 0; k-- {
		// fmt.Printf("decrement k: %v, i: %v, j: %v, nums1: %v, nums2: %v \n", k, i, j, nums1[i], nums2[j])
		if i >= 0 && j >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else if j >= 0 {
			nums1[k] = nums2[j]
			j--
		}
		// time.Sleep(1 * time.Second)
	}
}

func merge_dewi(nums1 []int, m int, nums2 []int, n int) {
	for len(nums2) > 0 {
		left2, right2 := 0, len(nums2)-1
		for left2 <= right2 {

			left1, right1 := 0, len(nums1)-1
			for left1 <= right1 {
				fmt.Printf("nums2 vs nums1 %v, %v, %v \n", nums2[left2], nums1[left1], nums2[left2] < nums1[left1])
				fmt.Printf("nums1 %v \n", nums1)

				if nums2[left2] < nums1[left1] || (nums1[left1] == 0 && left1 > m-n-1) {
					insertAt(nums1, left1, nums2[left2])
					left2 += 1
				}

				/*else if nums2[left2] > nums1[left1] && nums1[left1] == 0 {
					insertAt(nums1, left1, nums2[left2])
					left2 += 1
				}*/

				if left2 > len(nums2)-1 {
					return
				}

				left1 += 1
				time.Sleep(1 * time.Second)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func insertAt(nums []int, index, num int) []int {
	copy(nums[index+1:], nums[index:])
	nums[index] = num
	// fmt.Printf("copied nums %v, %v, %v \n", nums[index+1:], nums[index:], nums)
	return nums
}
