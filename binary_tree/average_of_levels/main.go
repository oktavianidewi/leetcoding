package main

import (
	"fmt"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructTreeNode(arr []interface{}, i, n int) *TreeNode {
	root := &TreeNode{}

	if i >= n {
		return root
	}
	if arr[i] != nil {
		root.Val = arr[i].(int)
	}
	root.Left = constructTreeNode(arr, 2*i+1, n)
	root.Right = constructTreeNode(arr, 2*i+2, n)
	return root
}

func averageOfLevels(root *TreeNode) []float64 {
	var answer []float64
	var queue []*TreeNode
	queue = append(queue, root)
	fmt.Printf("queue %v \n", queue)

	// queue = (queue)[1:]
	// gimana cara tau level nya ya????

	// for root != nil {
	for root != nil {

		// if root.Left.Val != 0 {
		queue = append(queue, root.Left)
		// }

		// if root.Right.Val != 0 {
		queue = append(queue, root.Right)
		// }

		answer = append(answer, float64(root.Val))
		fmt.Printf("answer %v \n", answer)

		// stop loop: bisa set queue jadi 0 atau root = nil
		if len(queue) > 1 {
			// dequeue
			queue = (queue)[1:]
			// set new root value
			root = queue[0]
		} else {
			root = nil
		}

		fmt.Printf("after dequeue %v , last val %v, len %v \n", queue, queue[0].Val, len(queue))
		time.Sleep(1 * time.Second)
	}

	return answer
}

func main() {
	nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	// root := []interface{}{3, 9, 20, 15, 7}
	output := []float64{3.00000, 14.50000, 11.00000}

	root := constructTreeNode(nums, 0, len(nums))
	answer := averageOfLevels(root)

	fmt.Printf("expected %v, answer %v \n", output, answer)
}
