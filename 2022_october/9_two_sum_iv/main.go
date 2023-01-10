package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func traverse_dfs(root *TreeNode) ([]int, map[int]bool) {}

// traverse bfs dengan condition que yang beda

// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/discuss/2680449/Golang-or-Hashmap

func traverse_bfs(root *TreeNode) ([]int, map[int]int) {
	var nums []int
	mapNums := make(map[int]int)
	que := []*TreeNode{root}

	for len(que) > 0 {
		if root.Left != nil {
			que = append(que, root.Left)
		}

		if root.Right != nil {
			que = append(que, root.Right)
		}

		nums = append(nums, que[0].Val)
		mapNums[que[0].Val]++
		if len(que) > 1 {
			root = que[1]
		}
		que = que[1:]
	}

	return nums, mapNums
}

func isTwoSum(k int, nums []int, mapNums map[int]int) bool {
	for key, _ := range mapNums {
		if mapNums[k-key] > 0 && k-key != key {
			return true
		}
		if k-key == key && mapNums[k-key] > 1 {
			return true
		}
	}
	return false
}

func findTarget(root *TreeNode, k int) bool {
	// traverse, move to array
	if root == nil {
		return false
	}

	nums, mapNums := traverse_bfs(root)

	res := isTwoSum(k, nums, mapNums)
	// two sum array with hashmap
	return res
}

func main() {
	// [5,3,6,2,4,null,7]
	// tree := &TreeNode{Val: 5}
	// tree.Left = &TreeNode{Val: 3}
	// tree.Left.Left = &TreeNode{Val: 2}
	// tree.Left.Right = &TreeNode{Val: 4}
	// tree.Right = &TreeNode{Val: 6}
	// tree.Right.Right = &TreeNode{Val: 7}
	// k := 9 // true
	// k := 28 // false
	// k := 11 // false

	tree := &TreeNode{Val: 1}
	k := 2

	twoSumExist := findTarget(tree, k)
	fmt.Printf("twoSumExist %v \n", twoSumExist)
}
