package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func convertArrToLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	moving := head
	for _, v := range nums {
		temp := &ListNode{
			Val: v,
		}
		moving.Next = temp
		moving = moving.Next
	}
	return head.Next
}

func traverse(head *ListNode) {
	// traverse
	for head != nil {
		fmt.Printf("content %v \n", head.Val)
		head = head.Next
	}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	list, moving, worker := head, head, head

	events := &ListNode{}
	movingEvents := events

	// traverse to separate odd and even node
	counter := 1
	for moving != nil {
		if counter%2 == 0 {
			fmt.Printf("counter %v ", counter)
			movingEvents.Next = moving
			movingEvents = movingEvents.Next

			// set nil to the last events node
			if moving.Next == nil || moving.Next.Next == nil {
				movingEvents.Next = nil
			}

			moving = moving.Next

		} else {
			moving = moving.Next
			if worker.Next != nil {
				worker.Next = worker.Next.Next
				worker = worker.Next
			}
		}
		counter++
	}

	// access to the tail
	tail := list
	for tail.Next != nil {
		tail = tail.Next
	}

	// combine events and odds node
	tail.Next = events.Next

	return list
}
