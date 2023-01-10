package main

import (
	"fmt"
	"testing"
)

func Test_MyHashMap(t *testing.T) {
	hash := Constructor()
	hash.Put(1, 1)
	hash.Put(2, 2)
	a := hash.Get(1)
	fmt.Printf("hash value %v \n", a)

	a = hash.Get(3)
	fmt.Printf("hash value %v \n", a)

	hash.Put(2, 1)
	a = hash.Get(2)
	fmt.Printf("hash value %v \n", a)

	hash.Remove(2)
	a = hash.Get(2)
	fmt.Printf("hash value %v \n", a)
}
