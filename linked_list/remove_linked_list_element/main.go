package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func convertArrToLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	insert := head
	for _, v := range nums {
		temp := &ListNode{
			Val: v,
		}
		insert.Next = temp
		insert = insert.Next
	}
	return head.Next
}

func traverse(list *ListNode) {
	// traverse to check data
	for list != nil {
		fmt.Printf("input %v > ", list.Val)
		list = list.Next
	}
}
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	list, moving, worker := head, head, head

	for moving != nil {
		// fmt.Printf("moving.Val %v \n", moving.Val)

		if moving.Val == val {
			if worker.Val == val {
				worker = worker.Next
				list = worker
			} else {
				for worker.Next.Val != val {
					worker = worker.Next
				}
				// fmt.Printf("worker %v \n", worker)
				worker.Next = moving.Next
			}
		}
		moving = moving.Next
	}
	// fmt.Printf("list %v \n", list)
	return list
}
