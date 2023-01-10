package maincoba

import (
	"fmt"
	"testing"
)

func Test_MyLinkedList(t *testing.T) {
	obj := Constructor()

	// before deleted
	param_1 := obj.Get(1)
	fmt.Printf("param_1 %v \n", param_1)

	obj.AddAtHead(1)
	// obj.AddAtTail(3)
	// obj.AddAtIndex(1, 2)
	// obj.DeleteAtIndex(1)

	// after deleted
	// param_1 = obj.Get(1)
	// fmt.Printf("param_1 %v \n", param_1)
}
