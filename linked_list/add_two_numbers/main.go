package main

import "fmt"

// Definition for singly-linked list.
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	listResult := new(ListNode)
	pListResult := listResult
	sum, carrier := 0, 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			sum = l2.Val + carrier
		} else if l2 == nil {
			sum = l1.Val + carrier
		} else {
			sum = l1.Val + l2.Val + carrier
		}

		if sum > 9 {
			carrier = 1
		} else {
			carrier = 0
		}

		// increment result
		pListResult.Next = &ListNode{Val: sum % 10}
		pListResult = pListResult.Next

		// increment l1 and l2
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carrier > 0 {
		pListResult.Next = &ListNode{Val: carrier}
	}
	return listResult.Next
}
