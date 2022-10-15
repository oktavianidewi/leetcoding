package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func traverse(head *ListNode) {
	if head != nil {
		_traverse(head)
	}
	return
}

func _traverse(head *ListNode) {
	if head != nil {
		_traverse(head.Next)
	} else {
		return
	}
	fmt.Printf("node val %v \n", head.Val)
}

func convertArrToLinkedList(nums []int) *ListNode {
	head := new(ListNode)
	pointer := head
	for _, i := range nums {
		temp := &ListNode{
			Val: i,
		}
		pointer.Next = temp
		pointer = pointer.Next
	}
	return head.Next
}
func rotateRight(head *ListNode, k int) *ListNode {
	// preprocess
	if head == nil {
		return nil
	}

	lenList := 0
	pointer := head
	for pointer != nil {
		lenList++
		pointer = pointer.Next
	}

	k = k % lenList

	if k < 1 {
		return head
	}

	list, slow, fast := head, head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	newHead := slow.Next
	fast.Next = list
	slow.Next = nil

	return newHead
}
