package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {

	temp := head
	reverseIndex := 0
	for head != nil {
		reverseIndex++
		head = head.Next
	}
	var decimal int

	head = temp
	for head != nil {
		decimal = decimal + head.Val*int(math.Pow(2, float64(reverseIndex-1)))
		reverseIndex--
		head = head.Next
	}

	fmt.Printf("decimal %v \n", decimal)
	return decimal
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

func main() {
	nums := []int{1, 0, 1}
	output := 5

	// nums := []int{0}
	// output := 0

	head := convertArrToLinkedList(nums)

	result := getDecimalValue(head)
	fmt.Printf("expected %v result %v \n", output, result)
}
