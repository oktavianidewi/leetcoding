package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func convertArrToLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	moving := head
	for _, i := range nums {
		moving.Next = &ListNode{
			Val: i,
		}
		moving = moving.Next
	}
	return head.Next
}

func _traverse(head *ListNode) {
	if head != nil {
		_traverse(head.Next)
	} else {
		return
	}
	fmt.Printf("node val %v\n", head.Val)
}

func traverse(head *ListNode) {
	if head == nil {
		return
	}
	_traverse(head)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// preprocess
	if list1 == nil && list2 == nil {
		return nil
	}

	// temp := &ListNode{}
	temp := new(ListNode)
	pointer := temp

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pointer.Next = list1
			list1 = list1.Next
		} else {
			pointer.Next = list2
			list2 = list2.Next
		}
		pointer = pointer.Next
	}

	if list1 == nil {
		pointer.Next = list2
	} else if list2 == nil {
		pointer.Next = list1
	}

	return temp.Next
}
