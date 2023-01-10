package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	//
	if fast == nil {
		head = head.Next
		return head
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// n := 2

	nums := []int{1, 2, 3, 4, 5}
	n := 3

	// nums := []int{1}
	// n := 1

	// nums := []int{1, 2}
	// n := 1

	head := ConvertArrToLinkedList(nums)
	head = removeNthFromEnd(head, n)
	fmt.Printf("head %v \n", head)
	// traverse

	x := head
	// traverse
	for x != nil {
		fmt.Printf("value x %v > ", x)
		x = x.Next
	}
}
