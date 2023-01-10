package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// fix added setelah lihat solusi ini: https://leetcode.com/problems/sum-of-left-leaves/discuss/492825/golang-bfs-0ms-100-2mb-100

func sumOfLeftLeaves(root *TreeNode) int {
	var total int
	var que []*TreeNode
	que = append(que, root)
	for root != nil {
		if root.Left != nil {
			if root.Left.Left == nil && root.Left.Right == nil {
				total = total + root.Left.Val
			}
			// fmt.Printf("total %v \n", total)
			que = append(que, root.Left)
		}
		if root.Right != nil {
			que = append(que, root.Right)
		}

		if len(que) > 1 {
			que = que[1:]
			root = que[0]
		} else {
			root = nil
		}
	}
	return total
}

func main() {
	// nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	// output := 24

	// tree := &TreeNode{Val: 3}
	// tree.Left = &TreeNode{Val: 9}
	// tree.Right = &TreeNode{Val: 20}
	// tree.Right.Left = &TreeNode{Val: 15}
	// tree.Right.Right = &TreeNode{Val: 7}

	// [1,2,3,4,5]
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	tree.Left.Left = &TreeNode{Val: 4}
	tree.Left.Right = &TreeNode{Val: 5}
	tree.Right = &TreeNode{Val: 3}
	output := 4

	result := sumOfLeftLeaves(tree)
	fmt.Printf("expected %v result %v \n", output, result)
}
