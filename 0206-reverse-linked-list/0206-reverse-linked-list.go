/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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