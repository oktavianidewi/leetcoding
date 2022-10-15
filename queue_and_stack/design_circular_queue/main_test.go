package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyCircularQueue(t *testing.T) {
	assert := assert.New(t)
	// construct the queue object, or intanciate the queue struct
	circularQ := Constructor(5)

	assert.Equal(circularQ.IsEmpty(), true)
	assert.Equal(circularQ.IsFull(), false)

	// enqueue
	for i := 0; i <= 5; i++ {
		circularQ.EnQueue(i)
	}
	assert.Equal(circularQ.IsFull(), true)

	circularQ.DeQueue()
	fmt.Printf("show queue %v \n", circularQ.ShowQueue())

	circularQ.DeQueue()
	fmt.Printf("show queue %v \n", circularQ.ShowQueue())
	fmt.Printf("front queue %v \n", circularQ.Front())
	fmt.Printf("rear queue %v \n", circularQ.Rear())

}
