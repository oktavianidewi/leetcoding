package main

// seharusnya gak boleh pakai built-in func di golang
// example pakai struct dan array -> https://shareablecode.com/snippets/golang-solution-for-leetcode-problem-design-hashmap-1AWJ-QCms

// array dan slice, bedanya apa? slice -> pre-defined array
// map gunanya apa? dict
// kapan pakai make(), kapan tidak? bebas, salah satu pilihan untuk initialize map

type MyHashMap struct {
	hash map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		// 2 ways to initialize map
		// hash: make(map[int]int),
		hash: map[int]int{},
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.hash[key] = value
}

func (this *MyHashMap) Get(key int) int {
	if _, ok := this.hash[key]; ok {
		return this.hash[key]
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	if _, ok := this.hash[key]; ok {
		delete(this.hash, key)
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
