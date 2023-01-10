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

// left,right, root
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	// func pass root and &ans parameter
	// &ans means : the address of the ans variable
	_postorder(root, &ans)
	return ans
}

// todo inline changes, ans menggunakan * bukan berarti untuk derefrence variable value, tapi untuk menunjukkan bahwa ans menggunakan data type pointer
func _postorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	// * untuk dereference variable. arr -> mengakses value dari pointer arr

	_postorder(root.Left, arr)
	_postorder(root.Right, arr)
	*arr = append(*arr, root.Val)
}

func main() {
	nums := []interface{}{1, nil, 2, 3}
	// nums := []interface{}{1, 2, 3, 4, 5, 6}
	// nums := []interface{}{1, 2, 3, 4, 5, 6, 7}
	// root := []interface{}{1, 2, 3, 4, 5, 6, 6, 6, 6, 6}
	tree := constructTreeNode(nums, 0, len(nums))

	postorder := postorderTraversal(tree)
	fmt.Printf("postorder %v \n", postorder)
}
