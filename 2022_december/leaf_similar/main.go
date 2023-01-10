package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/leaf-similar-trees/solutions/301035/golang-solution/
// dfs
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leave1, leave2 := []int{}, []int{}
	dfs(root1, &leave1)
	dfs(root2, &leave2)
	// fmt.Printf("leave1 %v, leave2 %v \n", leave1, leave2)
	if len(leave1) == len(leave2) {
		for i := 0; i < len(leave1); i++ {
			if leave1[i] != leave2[i] {
				return false
			}
		}
		return true
	}
	return false
}

func dfs(root *TreeNode, leave *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leave = append(*leave, root.Val)
	}
	dfs(root.Left, leave)
	dfs(root.Right, leave)
}

// bfs
// bfs will append value of root in unordered manner, coba kalo pake pointer (?) -> sama aja
func leafSimilar_bfs(root1 *TreeNode, root2 *TreeNode) bool {
	leave1, leave2 := []int{}, []int{}
	getLeave(root1, &leave1)
	getLeave(root2, &leave2)

	// compare leave1 and leave2
	if len(leave1) != len(leave2) {
		return false
	}
	fmt.Printf("leave1 %v \n", leave1)
	fmt.Printf("leave2 %v \n", leave2)

	// for i := 0; i < len(leave1); i++ {
	// 	if leave1[i] == leave2[i] {
	// 		return true
	// 	}
	// }
	return true
}

func getLeave(root *TreeNode, leave *[]int) {
	que := []*TreeNode{root}
	for root != nil {
		if root.Left != nil {
			que = append(que, root.Left)
		}
		if root.Right != nil {
			que = append(que, root.Right)
		}
		if root.Left == nil && root.Right == nil {
			*leave = append(*leave, root.Val)
		}
		if len(que) > 1 {
			que = que[1:]
			root = que[0]
		} else {
			root = nil
		}
	}
}

func main() {
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 5}
	root1.Left.Left = &TreeNode{Val: 6}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Left.Right.Left = &TreeNode{Val: 7}
	root1.Left.Right.Right = &TreeNode{Val: 4}
	root1.Right = &TreeNode{Val: 1}
	root1.Right.Left = &TreeNode{Val: 9}
	root1.Right.Right = &TreeNode{Val: 8}

	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 5}
	root2.Left.Left = &TreeNode{Val: 6}
	root2.Left.Right = &TreeNode{Val: 7}
	root2.Right = &TreeNode{Val: 1}
	root2.Right.Left = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 2}
	root2.Right.Right.Left = &TreeNode{Val: 9}
	root2.Right.Right.Right = &TreeNode{Val: 8}

	expected := true

	result := leafSimilar(root1, root2)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
