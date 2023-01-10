package main

import "fmt"

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

// kok bisa reversed? karena recursion menyimpan semua proses print dalam stack, kemudian mengeluarkan sesuai urutan stack (FIFO)
func traverse(list *ListNode) {
	// traverse to check data recursively
	if list.Next != nil {
		traverse(list.Next)
	}
	fmt.Printf("traverse %v ", list.Val)
}

var top *ListNode

func isPalindrome(head *ListNode) bool {
	top = head
	return _isPalindrome(head)
}

func _isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	if !_isPalindrome(head.Next) {
		return false
	}

	if head.Val == top.Val {
		top = top.Next
		return true
	}

	return false
}
