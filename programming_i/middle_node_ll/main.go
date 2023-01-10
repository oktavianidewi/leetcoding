package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func middleNode(head *ListNode) *ListNode {
	lengthNode := 0
	temp := head
	for head != nil {
		lengthNode++
		head = head.Next
	}

	midNode := lengthNode >> 1
	fmt.Printf("midNode %v \n", midNode)

	head = temp
	for i := 0; i < midNode; i++ {
		head = head.Next
	}

	return head
}

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	// output := []int{3, 4, 5}

	nums := []int{1, 2, 3, 4, 5, 6}
	output := []int{4, 5, 6}

	head := convertArrToLinkedList(nums)
	res := middleNode(head)

	fmt.Printf("expected %v result %v \n", output, res.Val)
}
