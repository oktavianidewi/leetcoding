package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q == nil || p == nil && q != nil {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	tree1 := &TreeNode{Val: 1}
	tree1.Left = &TreeNode{Val: 2}
	tree1.Right = &TreeNode{Val: 3}
	tree2 := &TreeNode{Val: 1}
	tree2.Left = &TreeNode{Val: 2}
	tree2.Right = &TreeNode{Val: 3}
	expected := true

	// tree1 := &TreeNode{Val: 1}
	// tree1.Left = &TreeNode{Val: 2}
	// tree2 := &TreeNode{Val: 1}
	// tree2.Right = &TreeNode{Val: 2}
	// expected := false

	result := isSameTree(tree1, tree2)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
