package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	que := []*TreeNode{root}
	for root != nil {
		if root.Left != nil {
			que = append(que, root.Left)
		}

		if root.Right != nil {
			que = append(que, root.Right)
		}

		if root.Val == val {
			return root
		}

		if len(que) > 1 {
			que = que[1:]
			root = que[0]
		} else {
			root = nil
		}
	}
	return nil
}

func main() {
	tree := &TreeNode{Val: 4}
	tree.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 7}
	tree.Left.Left = &TreeNode{Val: 1}
	tree.Left.Right = &TreeNode{Val: 3}
	// val := 2
	val := 5 // harusnya return kosong
	result := searchBST(tree, val)
	fmt.Printf("result %v , val %v \n", result, result.Val)
}
