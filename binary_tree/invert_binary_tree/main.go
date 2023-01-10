package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func main() {
	// tree := &TreeNode{Val: 4}
	// tree.Left = &TreeNode{Val: 2}
	// tree.Right = &TreeNode{Val: 7}
	// tree.Left.Left = &TreeNode{Val: 1}
	// tree.Left.Right = &TreeNode{Val: 3}
	// tree.Right.Left = &TreeNode{Val: 6}
	// tree.Right.Right = &TreeNode{Val: 9}

	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}

	// tree := &TreeNode{Val: 2}
	// tree.Left = &TreeNode{Val: 1}
	// tree.Right = &TreeNode{Val: 3}

	result := invertTree(tree)
	fmt.Println(result)
}
