/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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