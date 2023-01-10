package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var sumValidNum int
	que := []*TreeNode{root}
	for root != nil {
		if root.Left != nil {
			que = append(que, root.Left)
		}
		if root.Right != nil {
			que = append(que, root.Right)
		}
		if root.Val >= low && root.Val <= high {
			sumValidNum += root.Val
		}

		if len(que) > 1 {
			que = que[1:]
			root = que[0]
		} else {
			root = nil
		}
	}
	return sumValidNum
}

func main() {
	// 1st tree
	// tree := &TreeNode{Val: 10}
	// tree.Left = &TreeNode{Val: 5}
	// tree.Left.Left = &TreeNode{Val: 3}
	// tree.Left.Right = &TreeNode{Val: 7}
	// tree.Right = &TreeNode{Val: 15}
	// tree.Right.Right = &TreeNode{Val: 18}
	// low, high, expected := 7, 15, 32

	// 2nd tree
	tree := &TreeNode{Val: 10}
	tree.Left = &TreeNode{Val: 5}
	tree.Left.Left = &TreeNode{Val: 3}
	tree.Left.Left.Left = &TreeNode{Val: 1}
	tree.Left.Right = &TreeNode{Val: 7}
	tree.Left.Right.Left = &TreeNode{Val: 6}
	tree.Right = &TreeNode{Val: 15}
	tree.Right.Left = &TreeNode{Val: 13}
	tree.Right.Right = &TreeNode{Val: 18}
	low, high, expected := 6, 10, 23

	result := rangeSumBST(tree, low, high)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
