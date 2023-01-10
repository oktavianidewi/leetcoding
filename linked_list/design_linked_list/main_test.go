package structure

import (
	"fmt"
	"testing"
)

func Test_MyLinkedList(t *testing.T) {
	obj := NewLinkedList()
	var param_1 int

	fmt.Printf("length %v \n", obj.Length())

	obj.DeleteAtIndex(2)

	obj.AddAtHead(100)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	obj.AddAtTail(203)

	fmt.Printf("length %v \n", obj.Length())
	obj.Traverse()

	param_1 = obj.Get(0)
	fmt.Printf("param_1 %v \n", param_1)

	param_1 = obj.Get(2)
	fmt.Printf("param_1 %v \n", param_1)

	obj.DeleteAtIndex(2)
	obj.Traverse()

	// obj.DeleteAtIndex(3)
	// obj.Traverse()

	// obj.Traverse()

	// before deleted
	// var param_1 int
	// param_1 = obj.Get(0)
	// fmt.Printf("param_1 %v \n", param_1)
	// for x := 0; x < 10; x++ {
	// 	obj.AddAtHead(x)
	// }
	// fmt.Printf("length %v \n", obj.Length())
	// // fmt.Printf("length %v \n", obj.Length())

	// obj.AddAtIndex(10, 999)
	// obj.AddAtTail(1000)
	// param_1 = obj.Get(11)
	// fmt.Printf("param_1 %v \n", param_1)
	// obj.AddAtIndex(1, 2)

	// after deleted
	// param_1 = obj.Get(1)
	// fmt.Printf("param_1 %v \n", param_1)
}
