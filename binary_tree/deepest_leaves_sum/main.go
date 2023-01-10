package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	leftLeaveVal := dfs_left(root.Left)
	rightLeaveVal := dfs_right(root.Right)
	fmt.Printf("leftLeaveVal %v, rightLeaveVal %v \n", leftLeaveVal, rightLeaveVal)
	return leftLeaveVal + rightLeaveVal
}

func dfs_left(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	if root.Left != nil && root.Right == nil {
		return dfs_left(root.Left)
	}
	if root.Left == nil && root.Right != nil {
		return dfs_left(root.Right)
	}
	return dfs_left(root.Left)
}

func dfs_right(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	if root.Left != nil && root.Right == nil {
		return dfs_right(root.Left)
	}
	if root.Left == nil && root.Right != nil {
		return dfs_right(root.Right)
	}
	return dfs_right(root.Right)
}

func main() {
	// create 1st tree
	// tree := TreeNode{Val: 1}
	// tree.Left = &TreeNode{Val: 2}
	// tree.Left.Left = &TreeNode{Val: 4}
	// tree.Left.Right = &TreeNode{Val: 5}
	// tree.Left.Left.Left = &TreeNode{Val: 7}
	// tree.Right = &TreeNode{Val: 3}
	// tree.Right.Right = &TreeNode{Val: 6}
	// tree.Right.Right.Right = &TreeNode{Val: 8}
	// expected := 15

	// create 2nd tree
	tree := TreeNode{Val: 6}
	tree.Left = &TreeNode{Val: 7}
	tree.Left.Left = &TreeNode{Val: 2}
	tree.Left.Right = &TreeNode{Val: 7}
	tree.Left.Left.Left = &TreeNode{Val: 9}
	tree.Left.Right.Left = &TreeNode{Val: 1}
	tree.Left.Right.Right = &TreeNode{Val: 4}
	tree.Right = &TreeNode{Val: 8}
	tree.Right.Left = &TreeNode{Val: 1}
	tree.Right.Right = &TreeNode{Val: 3}
	tree.Right.Right.Right = &TreeNode{Val: 5}
	expected := 19

	result := deepestLeavesSum(&tree)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
