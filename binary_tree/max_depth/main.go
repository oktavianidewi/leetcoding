package main

import (
	"fmt"
	"time"
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

// with recursion -> better code read, but less time DFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TODO: implement using stack (BFS)
func maxDepth_bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	// jika root != nil, pasti depth=1
	depth := 1
	start, end := 0, 1
	// loop ke tiap level
	for {
		// count ada brp child per level
		numOfChildren := 0
		for i := start; i < end; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
				numOfChildren++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				numOfChildren++
			}
		}
		// jika root tidak punya anak
		if numOfChildren == 0 {
			break
		}

		depth++
		start = end
		end = end + numOfChildren
	}

	return depth
}

func maxDepth_dewi(root *TreeNode) int {
	var level int

	if root == nil {
		return level
	}

	var levelLeft int
	var levelRight int

	queue := []*TreeNode{root}
	level = 1

	for len(queue) > 0 {
		fmt.Printf("root values %v, left %v, right %v \n", root.Val, root.Left, root.Right)

		root = queue[0]
		fmt.Printf("init queue %v , len %v \n", queue, len(queue))

		if root.Left != nil {
			queue = append(queue, root.Left)
			levelLeft++
		}

		if root.Right != nil {
			queue = append(queue, root.Right)
			levelRight++
		}

		fmt.Printf("after added queue %v , len %v \n", queue, len(queue))

		queue = queue[1:]

		// 1 as root counter
		fmt.Printf("levelLeft %v, levelRight %v \n", levelLeft, levelRight)
		level = max(levelLeft, levelRight)

		time.Sleep(1 * time.Second)

	}

	return level

}

func main() {

	// nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	// expectedRes := 3

	// manual tree
	// tree := &TreeNode{
	// 	Val: 3,
	// }
	// tree.Left = &TreeNode{Val: 9}
	// tree.Right = &TreeNode{Val: 20}

	// tree.Right.Left = &TreeNode{Val: 15}
	// tree.Right.Right = &TreeNode{Val: 7}

	// programmatically tree (kurang valid!)
	// tree := constructTreeNode(nums, 0, len(nums))

	// nums := []interface{}{1, nil, 2} // implement manually
	// expectedRes := 2
	// tree := &TreeNode{
	// 	Val: 1,
	// }
	// tree.Right = &TreeNode{Val: 2}
	// tree.Left = &TreeNode{Val: 3}
	// tree.Right.Right = &TreeNode{Val: 9}

	// error: [0,2,4,1,null,3,-1,5,1,null,6,null,8]
	expectedRes := 4
	tree := &TreeNode{
		Val: 0,
	}
	tree.Left = &TreeNode{Val: 2}
	tree.Left.Left = &TreeNode{Val: 1}
	tree.Left.Left.Left = &TreeNode{Val: 5}
	tree.Left.Left.Right = &TreeNode{Val: 1}

	tree.Right = &TreeNode{Val: 4}
	tree.Right.Left = &TreeNode{Val: 3}
	tree.Right.Right = &TreeNode{Val: -1}
	tree.Right.Left.Right = &TreeNode{Val: 6}
	tree.Right.Right.Right = &TreeNode{Val: 8}

	// res := maxDepth_dfs(tree)
	res := maxDepth(tree)
	fmt.Printf("result %v, expected result %v \n", res, expectedRes)
}
