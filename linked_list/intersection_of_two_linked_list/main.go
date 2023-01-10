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

// check bedanya dengan https://shareablecode.com/snippets/golang-solution-for-leetcode-problem-intersection-of-two-linked-lists-QWzp-ExPs

func ConvertAndIntersectArrToLinkedList(listA []int, skipA int, listB []int, skipB int) (*ListNode, *ListNode) {
	var headA, headB *ListNode
	headA = ConvertArrToLinkedList(listA)
	headB = ConvertArrToLinkedList(listB[:skipB])

	// pointing headB to headA
	pointB := headB
	for pointB.Next != nil {
		pointB = pointB.Next
	}
	// get intersection of headA
	pointA := headA
	for a := 0; a < skipA; a++ {
		pointA = pointA.Next
	}

	pointB.Next = pointA
	return headA, headB
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// use only O(1) memory -> gak bisa pake hash-node
	lenA := length(headA)
	lenB := length(headB)

	// swap, make headA the smallest
	if lenB < lenA {
		headA, headB = headB, headA
	}

	for headA != nil {
		fmt.Printf("compare %v to %v \n", headA, headB)
		if countains(headA, headB) {
			return headA
		}

		// incerement
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func length(node *ListNode) int {
	counter := 0
	for node != nil {
		counter++
		node = node.Next
	}
	return counter
}

func countains(keyword, dictionaries *ListNode) bool {
	for dictionaries != nil {
		if keyword == dictionaries {
			return true
		}
		dictionaries = dictionaries.Next
	}
	return false
}

func getIntersectionNode_key(headA, headB *ListNode) *ListNode {
	//boundary check
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	//if a & b have different len, then we will stop the loop after second iteration
	for a != b {
		//for the end of first iteration, we just reset the pointer to the head of another linkedlist
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		fmt.Printf("a = %v b = %v\n", a, b)
	}
	return a
}

func main() {
	// gimana cara bikin intersection array linked list?
	// intersection ini harus di pointer location yang sama
	// case 1
	// listA := []int{4, 1, 8, 4, 5}
	// skipA := 2
	// listB := []int{5, 6, 1, 8, 4, 5}
	// skipB := 3

	// case 2
	// listA := []int{1, 9, 1, 2, 4}
	// skipA := 3
	// listB := []int{3, 2, 4}
	// skipB := 1

	// case 3
	// listA := []int{2, 6, 4}
	// skipA := 3
	// listB := []int{1, 5}
	// skipB := 1

	listA := []int{2, 6, 4}
	skipA := 3
	listB := []int{1, 5}
	skipB := 2

	/*
		error at:

		Input:
		0
		[2,6,4]
		[1,5]
		3
		2
		Output:
		Intersected at '0'
		Expected:
		No intersection

	*/

	headA, headB := ConvertAndIntersectArrToLinkedList(listA, skipA, listB, skipB)

	intersection := getIntersectionNode(headA, headB)
	intersection_key := getIntersectionNode_key(headA, headB)

	fmt.Printf("intersection %v\n", intersection)
	fmt.Printf("intersection_key %v\n", intersection_key)

}
