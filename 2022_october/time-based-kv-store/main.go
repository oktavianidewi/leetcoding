package main

import (
	"fmt"
)

type Node struct {
	timestamp int
	value     string
}

// a hash table that stores array of node (why node? because we need to store 2 values: timestamp and int value)
type TimeMap struct {
	storage map[string][]Node
}

func Constructor() TimeMap {
	return TimeMap{
		storage: make(map[string][]Node),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.storage[key] = append(this.storage[key], Node{timestamp: timestamp, value: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	nodes, ok := this.storage[key]
	if !ok {
		return ""
	}

	// find by timestamp, kenapa pakai binary search? karena value timestamp selalu sorted ascendingly
	left := 0
	right := len(nodes)
	mid := (left + right) / 2
	for i := left; i < right; i++ {
		// fmt.Printf("find %v timestamp %v \n", key, timestamp)
		// fmt.Printf("storage %v \n", this.storage)
		// fmt.Printf("left %v , mid %v , right %v \n", left, mid, right)

		if timestamp == nodes[mid].timestamp {
			return nodes[mid].value
		}

		// kondisi apa? yang memenuhi nodes[mid+1].timestamp > timestamp ?
		if (mid == right-1 || nodes[mid+1].timestamp > timestamp) && nodes[mid].timestamp <= timestamp {
			return nodes[mid].value
		}

		if timestamp < nodes[mid].timestamp {
			right = mid - 1
		} else if timestamp > nodes[mid].timestamp {
			left = mid + 1
		}

		// time.Sleep(1 * time.Second)
	}
	return ""
}

func main() {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)

	res := timeMap.Get("foo", 1)
	fmt.Printf("result foo-1 %v \n", res)

	res = timeMap.Get("abc", 1)
	fmt.Printf("result abc %v \n", res)

	res = timeMap.Get("foo", 3)
	fmt.Printf("result foo-3 %v \n", res)

	timeMap.Set("foo", "bar2", 4)

	res = timeMap.Get("foo", 4)
	fmt.Printf("result foo-4 %v \n", res)

	res = timeMap.Get("foo", 5)
	fmt.Printf("result foo-5 %v \n", res)

	// masih error disini, kenapa ya? yang asli nggak error!
	/*
		["TimeMap","set","set","get","get","get","get","get"]
		[[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
		output: [null,null,null,"","","","low","low"]
		expeced: [null,null,null,"","high","high","low","low"]
	*/
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
