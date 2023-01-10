/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    var queue []*Node
    var res []int
    queue = append(queue, root)  
    
    for len(queue) != 0 {
        tmp := queue[len(queue) - 1]
        queue = queue[:len(queue) - 1]
        if tmp != nil {
            res = append(res, tmp.Val)
            for i := len(tmp.Children) - 1; i >= 0; i-- {
                queue = append(queue, tmp.Children[i])
            }
        }
    }
    return res
}