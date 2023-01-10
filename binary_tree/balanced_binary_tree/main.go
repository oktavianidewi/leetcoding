package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {

}

func main() {
	tree := &TreeNode{Val: 3}
	tree.Left = &TreeNode{Val: 9}
	tree.Right = &TreeNode{Val: 20}
	tree.Right.Left = &TreeNode{Val: 15}
	tree.Right.Right = &TreeNode{Val: 7}
	isBalanced(tree)
}
