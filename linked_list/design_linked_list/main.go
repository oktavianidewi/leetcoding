package structure

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	node *Node
}

func NewLinkedList() *LinkedList {
	// initialize 0 index linkedlist
	return &LinkedList{
		// node: &Node{},
	}
}
func (l *LinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}
	// TODO loop by node
	currNode := l.node
	for i := 0; i < index; i++ {
		currNode = currNode.next
	}
	fmt.Printf("isi %v ", l.node)
	return currNode.val
}

func (l *LinkedList) Traverse() {
	currNode := l.node
	for currNode != nil {
		fmt.Printf("val %v > ", currNode.val)
		currNode = currNode.next
	}
}

func (l *LinkedList) Length() int {
	i := 0
	currNode := l.node
	fmt.Printf("fnode %v \n ", currNode)
	if currNode == nil {
		return i
	}
	for currNode != nil {
		currNode = currNode.next
		i++
	}
	return i
}

func (l *LinkedList) AddAtTail(value int) {
	length := l.Length()
	fmt.Printf("insert at position %v", length)
	l.AddAtIndex(length, value)
}
func (l *LinkedList) AddAtHead(value int) {
	// head = index 0
	l.AddAtIndex(0, value)
}

func (l *LinkedList) DeleteAtIndex(index int) {
	lenList := l.Length()

	// pre-process index
	if index < 0 {
		index = 0
	} else if index >= lenList {
		index = lenList - 1
	}

	// delete process

	if lenList <= 0 {
		return
	} else {
		currNode := l.node
		for i := 0; i < index-1; i++ {
			currNode = currNode.next
		}
		currNode.next = currNode.next.next
	}
}

func (l *LinkedList) AddAtIndex(index, value int) {
	temp := &Node{
		val:  value,
		next: nil,
	}

	// pre-process index
	lenList := l.Length()

	if index < 0 {
		index = 0
	} else if index >= lenList {
		index = lenList
	}

	// initialize node
	if lenList == 0 {
		l.node = temp
	} else {
		currNode := l.node

		if index <= lenList {
			// why index-1 ? to access previous node
			for i := 0; i < index-1; i++ {
				currNode = currNode.next
			}
			temp.next = currNode.next
			currNode.next = temp
		} else {
			for currNode.next != nil {
				currNode = currNode.next
			}
			currNode.next = temp
		}

	}
}

func (l *LinkedList) GetNode(pos int) *Node {
	currNode := l.node
	for i := 0; i < pos; i++ {
		currNode = currNode.next
	}
	return currNode

}

func (l *LinkedList) CircularAt(pos int) {
	getNode := l.GetNode(pos)

	currNode := l.node
	for currNode.next != nil {
		currNode = currNode.next
	}
	currNode.next = getNode
}
