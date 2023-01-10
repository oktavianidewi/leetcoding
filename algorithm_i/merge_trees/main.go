package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root := &TreeNode{Val: root1.Val + root2.Val}
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}

func main() {
	// root1, root2 := []interface{}{1, 3, 2, 5}, []interface{}{2, 1, 3, nil, 4, nil, 7}
	// result := []interface{}{3, 4, 5, 5, 4, nil, 7}

	// root1, root2 := []interface{}{1}, []interface{}{1, 2}
	// result := []interface{}{2, 2}

	// root1
	root1 := &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}

	x := mergeTrees(root1, root2)
	fmt.Printf("value x %v \n", x)
}
