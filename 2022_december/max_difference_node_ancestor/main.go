package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func maxAncestorDiff(root *TreeNode) int {
// 	maxDiff := math.MinInt
// 	var nums []int
// 	// tree to array
// 	dfs(root, &nums)
// 	fmt.Printf("nums %v \n", nums)
// 	// then cari diff value nya
// 	for i:=
// 	return maxDiff
// }

// func dfs(root *TreeNode, nums *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	*nums = append(*nums, root.Val)
// 	dfs(root.Left, nums)
// 	dfs(root.Right, nums)
// }

func maxAncestorDiff(root *TreeNode) int {
	res := 0
	fmt.Println(root, res, root.Val, root.Val)
	helper(root, &res, root.Val, root.Val)
	return res
}

func helper(cur *TreeNode, res *int, mn, mx int) {
	if cur == nil {
		return
	}
	m := max(abs(cur.Val-mn), abs(cur.Val-mx))
	*res = max(*res, m)

	fmt.Println(cur.Left, *res, mn, cur.Val, mx, cur.Val)
	fmt.Println(cur.Right, *res, mn, cur.Val, mx, cur.Val)
	helper(cur.Left, res, min(mn, cur.Val), max(mx, cur.Val))
	helper(cur.Right, res, min(mn, cur.Val), max(mx, cur.Val))
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func abs(x int) int {
	if x > 0 {
		return x
	}

	return -x
}

func main() {
	// tree1
	tree := &TreeNode{Val: 8}
	tree.Left = &TreeNode{Val: 3}
	tree.Left.Left = &TreeNode{Val: 1}
	tree.Left.Right = &TreeNode{Val: 6}
	tree.Left.Right.Left = &TreeNode{Val: 4}
	tree.Left.Right.Right = &TreeNode{Val: 7}
	tree.Right = &TreeNode{Val: 10}
	tree.Right.Right = &TreeNode{Val: 14}
	tree.Right.Right.Left = &TreeNode{Val: 13}
	expected := 3

	// tree2
	// tree := &TreeNode{Val: 1}
	// tree.Right = &TreeNode{Val: 2}
	// tree.Right.Right = &TreeNode{Val: 0}
	// tree.Right.Right.Left = &TreeNode{Val: 3}
	// expected := 3

	output := maxAncestorDiff(tree)
	fmt.Printf("expected %v, result %v \n", expected, output)
}
