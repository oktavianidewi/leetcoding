package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func convertArrToLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	temp := head

	for _, v := range nums {
		temp.Next = &ListNode{
			Val: v,
		}
		temp = temp.Next
	}
	return head.Next
}

func reverseList(list *ListNode) *ListNode {
	if list == nil {
		return nil
	}
	moving := list
	head := list.Next
	reversed := &ListNode{}
	counter := 0

	// get tail
	tail := list
	for tail.Next != nil {
		tail = tail.Next
	}

	for head != nil {
		// fmt.Printf("counter %v, moving %v, head %v \n", counter, moving, head)

		if reversed.Next == nil {
			moving.Next = nil
			reversed.Next = moving
		} else {
			// insert after head
			moving.Next = reversed.Next
			reversed.Next = moving
		}

		moving = head
		head = head.Next
		counter++
	}

	// fmt.Printf("tail %v \n", tail)
	// fmt.Printf("reversed.Next %v \n", reversed.Next)
	tail.Next = reversed.Next

	return tail
}
