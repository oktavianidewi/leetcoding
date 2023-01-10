package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return root.Val == root.Left.Val+root.Right.Val
}

func main() {
	// 1st tree
	tree := &TreeNode{Val: 10}
	tree.Left = &TreeNode{Val: 4}
	tree.Right = &TreeNode{Val: 6}
	expected := true

	// tree := &TreeNode{Val: 5}
	// tree.Left = &TreeNode{Val: 3}
	// tree.Right = &TreeNode{Val: 1}
	// expected := false

	result := checkTree(tree)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
