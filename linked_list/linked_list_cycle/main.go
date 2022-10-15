package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity : O(n)
// Space complexity: O(n)
func hasCycleWithHash(head *ListNode) bool {
	if head == nil {
		return false
	}
	// approach with Hash
	var hashListNode []*ListNode
	for head.Next != nil {
		if ok := contain(hashListNode, head); ok {
			return true
		}
		hashListNode = append(hashListNode, head)
		head = head.Next
	}
	return false
}

func contain(hashListNode []*ListNode, search *ListNode) bool {
	for _, node := range hashListNode {
		// either Val and Next value is same
		if node == search {
			return true
		}
	}
	return false
}

// time complexity is O(N+K), which is O(n)
// Space complexity : O(1).
func hasCycle(head *ListNode) bool {
	// approach 1 : Two Pointer
	// currNode := l.node
	// for currNode {
	// 	fmt.Printf("val %v > ", currNode.val)
	// 	currNode = currNode.next
	// }
	return false
}

func ConvertArrToLinkedList(nums []int) *ListNode {
	listNode := &ListNode{}

	// temp sebagai moving pointer, supaya starting point tetap di head
	// apakah temp membuat variable Node baru? tidak, karena dia mengakses variable pointer yg sama seperti listNode
	temp := listNode
	for k := 0; k < len(nums); k++ {
		temp.Next = &ListNode{
			Val: nums[k],
		}
		// increment
		temp = temp.Next
	}
	return listNode.Next
}

func CreateCircularList(nums []int, pos int) *ListNode {
	list := ConvertArrToLinkedList(nums)

	if pos < 0 {
		return list
	}

	circular := list
	for j := 0; j < pos; j++ {
		circular = circular.Next
	}

	tail := circular
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = circular

	return list
}

func main() {
	// define a circular linked-list with [3,2,0,-4], pos = 1

	// #0
	// arr := []int{3, 2, 0, -4}
	// pos := 1

	// #1
	// arr := []int{}
	// pos := -1

	// #2
	arr := []int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}
	pos := -1
	head := CreateCircularList(arr, pos)

	fmt.Printf("head %v \n", head)
	// fmt.Printf("hasCycle %v \n", hasCycleWithHash(head))
	fmt.Printf("is cycle %v\n", hasCycleWithHash(head))

}
