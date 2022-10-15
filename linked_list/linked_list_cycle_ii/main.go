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
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var slow, fast *ListNode = head, head

	for {
		// There is no cycle if we ever reach a nil value
		if fast == nil || fast.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	// masih nggak ngerti bagian ini untuk apa, tadi aku sampe sini
	// bagian ini untuk menentukan dimana irisannya di cycle berikutnya
	slow = head
	for slow != fast {
		fmt.Printf("slow %v \n", slow)
		slow = slow.Next

		fmt.Printf("fast %v \n", fast)
		fast = fast.Next
	}

	return fast
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
	// pos := -1

	// #1 err
	// arr := []int{}
	// pos := -1

	// #2 err
	// arr := []int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5}
	// pos := -1

	// #3
	arr := []int{3, 2, 0, -4}
	pos := 1

	// #4
	// arr := []int{1, 2}
	// pos := 0

	// arr := []int{1}
	// pos := -1

	head := CreateCircularList(arr, pos)

	fmt.Printf("detect cycle %v\n", detectCycle(head))
}
