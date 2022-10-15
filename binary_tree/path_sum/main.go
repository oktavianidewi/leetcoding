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

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	res := isPathSum(root, targetSum)
	return res

	// subRes := isPathSum_dewi(root, targetSum)
	// fmt.Printf("value result %v \n", subRes)
	// return subRes == 0
}

func isPathSum(root *TreeNode, target int) bool {
	// do simpler problem
	if root == nil {
		fmt.Printf("target %v\n", target)
		return target == 0
	}
	if root.Right == nil {
		return isPathSum(root.Left, target-root.Val)
	}
	if root.Left == nil {
		return isPathSum(root.Right, target-root.Val)
	}
	return isPathSum(root.Left, target-root.Val) || isPathSum(root.Right, target-root.Val)
}

func isPathSum_bfs(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	level := 1
	return level
}

// kunci jawaban: https://leetcode.com/explore/learn/card/data-structure-tree/17/solve-problems-recursively/537/discuss/267991/golang

func main() {
	// nums := []interface{}{}
	// target := 0

	// nums := []interface{}{1, 2}
	// nums := []interface{}{1, nil, 3}
	// nums := []interface{}{1, 2, 3}
	// target := 5

	// nums := []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}
	// nums := []interface{}{5, 4, 8, 1, 13, 13, 4, 7, 2, nil, nil, nil, 1}
	target := 22
	tree := constructTreeNode(nums, 0, len(nums))
	res := hasPathSum(tree, target)
	fmt.Printf("result %v \n", res)
}
