// https://leetcode.com/explore/learn/card/data-structure-tree/134/traverse-a-tree/931/discuss/406488/Golang-solution-with-queue

package main

import "fmt"

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

// inline queue with pointer
func enqueue(queue *[]TreeNode, root *TreeNode) {
	*queue = append(*queue, *root)
}

func dequeue(queue *[]TreeNode) {
	*queue = (*queue)[1:]
}

func peek(queue *[]TreeNode) *TreeNode {
	return &(*queue)[0]
}

// error problem : [1,2], output : [[1]], expected : [[1],[2]]
func levelOrder_dewi(root *TreeNode) [][]int {
	var ans [][]int
	// var level []int
	if root == nil {
		return ans
	}

	// queue
	var queue []TreeNode
	enqueue(&queue, root)
	// arr = append(arr, root.Val)
	// fmt.Printf("root.Left %v\n", root.Left)
	// fmt.Printf("root.Right %v\n", root.Right)

	for root != nil {
		var level []int
		if root.Left != nil {
			enqueue(&queue, root.Left)
		}
		if root.Right != nil {
			enqueue(&queue, root.Right)
		}

		// insert level value to array
		if len(ans) == 0 {
			val := peek(&queue)
			ans = append(ans, []int{val.Val})
		}

		if root.Left != nil && root.Right != nil {
			left := root.Left.Val
			right := root.Right.Val
			// fmt.Printf("left and right %v, %v \n", left, right)

			if left != 0 {
				level = append(level, left)
			}
			if right != 0 {
				level = append(level, right)
			}
		}

		if len(level) > 0 {
			ans = append(ans, level)
		}

		// fmt.Printf("ans %v\n", ans)

		dequeue(&queue)

		if len(queue) > 0 {
			root = peek(&queue)
		} else {
			root = nil
		}
	}
	return ans
}

// TODO: compare my solution with found solution
// solution
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)
		level := []int{}

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}

func main() {
	// nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	// nums := []interface{}{1}
	nums := []interface{}{1, 2}
	root := constructTreeNode(nums, 0, len(nums))
	answer := levelOrder(root)
	fmt.Print(answer)

	// var queue []int
	// enqueue(&queue, 10)
	// enqueue(&queue, 20)
	// enqueue(&queue, 30)
	// enqueue(&queue, 40)
	// fmt.Printf("queue %v \n", queue)
	// dequeue(&queue)
	// fmt.Printf("queue %v \n", queue)

}
