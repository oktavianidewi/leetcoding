package main

type MyHashSet struct {
	list map[int]int
}

func Constructor() MyHashSet {
	return MyHashSet{map[int]int{}}
}

func (this *MyHashSet) Add(key int) {
	this.list[key] = key
}

func (this *MyHashSet) Remove(key int) {
	delete(this.list, key)
}

func (this *MyHashSet) Contains(key int) bool {
	_, exist := this.list[key]
	return exist
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
