/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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