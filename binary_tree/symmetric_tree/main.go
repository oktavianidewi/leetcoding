package main

import (
	"fmt"
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

func isSymmetric_dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil && node2 != nil {
		return false
	}

	if node1 != nil && node2 == nil {
		return false
	}

	return node1.Val == node2.Val && isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left)
}

// bfs
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	que := []*TreeNode{root}
	levelVal := [][]int{{root.Val}}
	level := 1
	start, end := 0, 1

	for {
		numOfChild := 0
		childVal := []int{}

		for i := start; i < end; i++ {
			node := que[i]
			if node.Left != nil {
				que = append(que, node.Left)
				childVal = append(childVal, node.Left.Val)
				numOfChild++
			}
			if node.Right != nil {
				que = append(que, node.Right)
				childVal = append(childVal, node.Right.Val)
				numOfChild++
			}
		}

		if len(childVal) > 0 {
			levelVal = append(levelVal, childVal)
		}

		if numOfChild == 0 {
			break
		}
		start = end
		end = end + numOfChild
		level++
	}

	fmt.Printf("levelVal %v \n", levelVal)
	fmt.Printf("level %v \n", level)

	return false
}

func isPalindrome(nums []int) bool {
	return false
}

func main() {
	// nums := []interface{}{1, nil, 2}
	// nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	// nums := []interface{}{1, 2, 2, 3, 4, 4, 3}
	// tree := constructTreeNode(nums, 0, len(nums))

	// nums := []int{1, 2, 2, 3, 4, 4, 3}
	// 1st tree
	// tree := &TreeNode{Val: 1}
	// tree.Left = &TreeNode{Val: 2}
	// tree.Left.Left = &TreeNode{Val: 3}
	// tree.Left.Right = &TreeNode{Val: 4}
	// tree.Right = &TreeNode{Val: 2}
	// tree.Right.Left = &TreeNode{Val: 4}
	// tree.Right.Right = &TreeNode{Val: 3}

	// 2nd tree
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Left.Right = &TreeNode{Val: 3}
	tree.Right = &TreeNode{Val: 2}
	tree.Right.Right = &TreeNode{Val: 3}

	res := isSymmetric(tree)
	fmt.Printf("result %v \n", res)
}
