package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {

	return []int{}
}

func main() {
	tree := Node{Val: 1}
	tree.Children = append(tree.Children, &Node{Val: 3})
	tree.Children[0].Children = append(tree.Children[0].Children, &Node{Val: 5}, &Node{Val: 6})
	tree.Children = append(tree.Children, &Node{Val: 2})
	tree.Children = append(tree.Children, &Node{Val: 4})

	res := preorder(&tree)
	fmt.Printf("result %v \n", res)
}
