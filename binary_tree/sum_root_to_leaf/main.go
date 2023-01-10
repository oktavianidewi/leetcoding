package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	s := 0
	if root != nil {
		depth(root, 0*2+root.Val, &s)
	}
	return s
}

func depth(r *TreeNode, weight int, sum *int) {
	if r != nil {
		rl, rr := r.Left, r.Right
		if rl == nil && rr == nil {
			*sum = *sum + weight
		}
		if rl != nil {
			depth(rl, weight*2+rl.Val, sum)
			fmt.Println(rl, weight*2+rl.Val, *sum)
		}
		if rr != nil {
			depth(rr, weight*2+rr.Val, sum)
			fmt.Println(rr, weight*2+rr.Val, *sum)
		}
		// fmt.Println(rr.Val)
	}
}

func main() {
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 0}
	tree.Left.Left = &TreeNode{Val: 0}
	tree.Left.Right = &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 1}
	tree.Right.Left = &TreeNode{Val: 0}
	tree.Right.Right = &TreeNode{Val: 1}
	x := sumRootToLeaf(tree)
	fmt.Printf("hasil %v \n", x)
}
