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

func preorderTraversalIterative(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			res = append(res, cur.Val)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}

	return res
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	// func _preOrder pass root and &ans parameter
	// &ans means : the address of the ans variable
	_preOrder(root, &ans)
	return ans
}

// todo inline changes, ans menggunakan * bukan berarti untuk derefrence variable value, tapi untuk menunjukkan bahwa ans menggunakan data type pointer
func _preOrder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	// * untuk dereference variable. arr -> mengakses value dari pointer arr
	*arr = append(*arr, root.Val)

	_preOrder(root.Left, arr)
	_preOrder(root.Right, arr)
}

// node (center), left, right
func _preOrder_withoutPointer(root *TreeNode, ans []int) {
	if root == nil {
		return
	}

	ans = append(ans, root.Val)

	_preOrder_withoutPointer(root.Left, ans)
	_preOrder_withoutPointer(root.Right, ans)
}

func main() {
	nums := []interface{}{1, nil, 2, 3}
	// nums := []interface{}{10, 25, 2, 1, 14, 30, 5, 7}
	root := constructTreeNode(nums, 0, len(nums))
	preOrder := preorderTraversalIterative(root)
	fmt.Printf("preOrder %v \n", preOrder)

}
