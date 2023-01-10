package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth_x(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCounter := 0
	rightCounter := 0
	leftCount := dfsx(root.Left)
	rightCount := dfsx(root.Right)

	dfs(root.Left, &leftCounter)
	dfs(root.Right, &rightCounter)
	fmt.Printf("leftCount %v, rightCount %v \n", leftCount, rightCount)
	fmt.Printf("leftCounter %v, rightCounter %v \n", leftCounter, rightCounter)

	if leftCounter == 0 || rightCounter == 0 {
		return 1 + leftCounter + rightCounter
	}

	if leftCounter < rightCounter {
		return 1 + leftCounter
	}

	return 1 + rightCounter
}

func dfsx(root *TreeNode) (counter int) {
	if root == nil {
		return counter
	}
	counter++
	if root.Left != nil {
		return dfsx(root.Left)
	}

	if root.Right != nil {
		return dfsx(root.Right)
	}
	return counter
}

func dfs(root *TreeNode, counter *int) {
	if root == nil {
		return
	}

	*counter = *counter + 1
	if root.Left != nil {
		dfs(root.Left, counter)
	}

	if root.Right != nil {
		dfs(root.Right, counter)
	}
}

func main() {
	// tree1
	// root := &TreeNode{Val: 3}
	// root.Left = &TreeNode{Val: 9}
	// root.Right = &TreeNode{Val: 20}
	// root.Right.Left = &TreeNode{Val: 15}
	// root.Right.Right = &TreeNode{Val: 7}

	// kenapa masih error disini? [0,2,4,1,null,3,-1,5,1,null,6,null,8]
	// solution https://leetcode.com/problems/minimum-depth-of-binary-tree/solutions/264076/golang-4ms-recursive/

	// tree err
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Left.Left = &TreeNode{Val: 5}
	root.Left.Left.Right = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Left.Right = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: -1}
	root.Right.Right.Right = &TreeNode{Val: 8}

	// tree2
	// root := &TreeNode{Val: 2}
	// root.Right = &TreeNode{Val: 3}
	// root.Right.Right = &TreeNode{Val: 4}
	// root.Right.Right.Right = &TreeNode{Val: 5}
	// root.Right.Right.Right.Right = &TreeNode{Val: 6}

	// tree3
	// root := &TreeNode{Val: 0}
	// root.Left = &TreeNode{Val: 2}
	// root.Right = &TreeNode{Val: 4}
	// root.Left.Left = &TreeNode{Val: 1}
	// root.Left.Left.Left = &TreeNode{Val: 5}
	// root.Left.Left.Right = &TreeNode{Val: 1}
	// root.Right.Left = &TreeNode{Val: 3}
	// root.Right.Left.Right = &TreeNode{Val: 6}
	// root.Right.Right = &TreeNode{Val: -1}
	// root.Right.Right.Right = &TreeNode{Val: 8}

	x := minDepth(root)
	fmt.Println(x)

}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCounter := minDepth(root.Left)
	rightCounter := minDepth(root.Right)

	if leftCounter == 0 {
		return rightCounter + 1
	}
	if rightCounter == 0 {
		return leftCounter + 1
	}

	return min(leftCounter, rightCounter) + 1
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
