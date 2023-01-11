/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	var total int
	var que []*TreeNode
	que = append(que, root)
	for root != nil {
		if root.Left != nil {
			if root.Left.Left == nil && root.Left.Right == nil {
				total = total + root.Left.Val
			}
			// fmt.Printf("total %v \n", total)
			que = append(que, root.Left)
		}
		if root.Right != nil {
			que = append(que, root.Right)
		}

		if len(que) > 1 {
			que = que[1:]
			root = que[0]
		} else {
			root = nil
		}
	}
	return total
}