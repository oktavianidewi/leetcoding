package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func DisplayList(head *Node) {
	fmt.Printf("The list is: ")
	for ; head != nil; head = head.next {
		fmt.Printf("data : %v\n", head.val)
	}
}

// calculate the length
func Length(head *Node) int {
	len := 0
	for ; head != nil; head = head.next {
		len++
	}
	fmt.Printf("len %v \n", len)
	return len
}

// insertion at the end
func InsertAtEnd(head *Node, x int) *Node {
	temp := &Node{
		val:  x,
		next: nil,
	}
	curr := head
	// apakah linkedList kosong?
	if curr == nil {
		// ya, replace current node dengan temp node
		curr = temp
	} else {
		for ; curr.next != nil; curr = curr.next {
		}
		// chain the current node dengan temp node
		curr.next = temp
	}
	return curr
}

// This modifies the head pointer in the calling function
func InsertAtEndInPlace(head **Node, x int) {
	fmt.Println("\nIn InsertAtEndInPlace:")
	tmp := Node{next: nil, val: x}
	fmt.Println("Inserting", tmp, "at the end")
	if *head == nil {
		*head = &tmp
	} else {
		var p *Node
		p = *head
		for ; p.next != nil; p = p.next {
		}
		p.next = &tmp
	}
	fmt.Printf("Address of head node %p\n", &**head)
	fmt.Println("Address of head pointer ", &*head)
	fmt.Println("Returning")
}

func main() {
	var p *Node
	p = nil

	// Linked List Insertion
	InsertAtEndInPlace(&p, 2020)
	fmt.Printf("---- of node p %p\n", &*p)
	fmt.Println("---- of pointer p ", &p)
	fmt.Println("---- of value p ", *p)
	fmt.Println("---- of p ", p)
	DisplayList(p)

	// insertion
	// InsertAtEnd(p, 202)
	// fmt.Printf("Address of node p %v \n", &*p)
	// fmt.Printf("Address of pointer p %v \n", &p)
	// DisplayList(p)
}
