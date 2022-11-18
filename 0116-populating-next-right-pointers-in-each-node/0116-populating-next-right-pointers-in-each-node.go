/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {
        return root
    }
    queue := []*Node{root}
    for len(queue) > 0 {
        levelSize := len(queue)
        for i := 0; i < levelSize; i++ {
            currentNode := queue[0]
            queue = queue[1:]
            
            if i < levelSize - 1 {
                currentNode.Next = queue[0]
            } else if i == levelSize - 1 {
                currentNode.Next = nil
            }
            
            if currentNode.Left != nil {
                queue = append(queue, currentNode.Left)
            }
            
            if currentNode.Right != nil {
                queue = append(queue, currentNode.Right)
            }
        }
    }
    return root
}