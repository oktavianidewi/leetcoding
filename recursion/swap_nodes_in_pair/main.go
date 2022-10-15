package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func convertArrToLinkedList(nums []int) *ListNode {
	head := new(ListNode)
	pointer := head

	for _, i := range nums {
		temp := &ListNode{
			Val:  i,
			Next: nil,
		}
		pointer.Next = temp
		pointer = pointer.Next
	}
	return head.Next
}

func traverse(head *ListNode) {
	for head != nil {
		fmt.Printf("node %v > \n", head.Val)
		head = head.Next
	}
}

// recursive
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	_swapPairs(head)
	return head
}

func _swapPairs(head *ListNode) {
	node1 := head
	if node1 == nil {
		return
	}

	node2 := node1.Next
	if node2 != nil {
		// swap process
		node1.Val, node2.Val = node2.Val, node1.Val
		_swapPairs(node2.Next)
	} else {
		_swapPairs(node2)
	}
}
