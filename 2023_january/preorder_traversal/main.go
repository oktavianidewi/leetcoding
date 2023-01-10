package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// root, left, right
func preorder(root *TreeNode, result *[]int) {
	if root != nil {
		*result = append(*result, root.Val)
		preorder(root.Left, result)
		preorder(root.Right, result)
	}
}
func preorderTraversal(root *TreeNode) []int {
	var result []int
	preorder(root, &result)
	return result
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	exp := []int{1, 2, 3}
	res := preorderTraversal(root)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
