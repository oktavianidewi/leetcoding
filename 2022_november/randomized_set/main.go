package main

import (
	"fmt"
	"math/rand"
)

// type RandomizedSet struct {
// 	storage map[int]int
// 	data    []int
// }

// func Constructor() RandomizedSet {
// 	return RandomizedSet{
// 		storage: make(map[int]int),
// 		data:    []int{},
// 	}
// }

// func (this *RandomizedSet) Insert(val int) bool {
// 	if _, ok := this.storage[val]; !ok {
// 		this.storage[val] = len(this.data) + 1 // harus dicoba, kenapa ga simpan val as value
// 		this.data = append(this.data, val)
// 		fmt.Printf("insert storage %v \n", this.storage)
// 		fmt.Printf("insert data %v \n", this.data)
// 		return true
// 	}
// 	return false
// }

// func (this *RandomizedSet) Remove(val int) bool {
// // harus remove data, gimana caranya supaya remove dari array of data cost nya o(1)
// 	if _, ok := this.storage[val]; ok {
// 		this.storage[this.data[len(this.data)-1]] = this.storage[val]
// 		this.data[this.storage[val]-1] = this.data[len(this.data)-1]
// 		this.data = this.data[:len(this.data)-1]

// 		fmt.Printf("delete storage %v \n", this.storage)
// 		fmt.Printf("delete data %v \n", this.data)

// 		delete(this.storage, val)
// 		return true
// 	}
// 	return false
// }

// func (this *RandomizedSet) GetRandom() int {
// 	x := rand.Intn(len(this.storage))
// 	return this.storage[x]
// }

/*
requirement:
RandomizedSet() Initializes the RandomizedSet object.
bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.

explanation:
- the key problem here is: "You must implement the functions of the class such that each function works in average O(1) time complexity."
- what kind of data types are available and time complexity const? butuh array karena untuk getRandom operation, array cost nya hanya o(1). https://leetcode.com/problems/insert-delete-getrandom-o1/discuss/684232/Python-w-Explanation
- https://leetcode.com/problems/insert-delete-getrandom-o1/discuss/2585661/Python-HashMap-Solution-with-Explanation
*/

type RandomizedSet struct {
	mapper map[int]int
	data   []int
}

func Constructor() RandomizedSet {

	r := RandomizedSet{}

	// hashmap -> supaya time complexity for read and write access o(1)
	r.mapper = make(map[int]int)

	// buat apa? untuk getRandom operation
	r.data = []int{}
	return r

}

func (this *RandomizedSet) Insert(val int) bool {
	if this.mapper[val] == 0 {
		// data apa yang disimpan mapper? lebar dari array data + 1 => kenapa?
		this.mapper[val] = len(this.data) + 1
		this.data = append(this.data, val)
		fmt.Printf("insert storage %v \n", this.mapper)
		fmt.Printf("insert data %v \n", this.data)
		return true
	}

	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, e := this.mapper[val]; e {
		this.mapper[this.data[len(this.data)-1]] = this.mapper[val]
		// this.mapper[]
		this.data[this.mapper[val]-1] = this.data[len(this.data)-1]
		this.data = this.data[:len(this.data)-1]
		delete(this.mapper, val)
		fmt.Printf("delete storage %v \n", this.mapper)
		fmt.Printf("delete data %v \n", this.data)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	n := rand.Intn(len(this.data))
	return this.data[n]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	random := Constructor()
	var boolResult bool
	var intResult int

	boolResult = random.Insert(1)
	fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	boolResult = random.Insert(2)
	fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	boolResult = random.Insert(3)
	fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	boolResult = random.Insert(4)
	fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	boolResult = random.Insert(5)
	fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	boolResult = random.Insert(6)
	fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	boolResult = random.Insert(7)
	fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	boolResult = random.Insert(8)
	fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	boolResult = random.Remove(3)
	fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	boolResult = random.Remove(1)
	fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	fmt.Printf("intResult %v \n", intResult)

	// boolResult = random.Remove(2)
	// fmt.Printf("boolResult %v , expected %v \n", boolResult, false)

	// boolResult = random.Insert(2)
	// fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	// intResult = random.GetRandom()
	// fmt.Printf("intResult %v , expected %v \n", intResult, 2)

	// boolResult = random.Remove(1)
	// fmt.Printf("boolResult %v , expected %v \n", boolResult, true)

	// boolResult = random.Insert(2)
	// fmt.Printf("boolResult %v , expected %v \n", boolResult, false)

	// intResult = random.GetRandom()
	// fmt.Printf("intResult %v , expected %v \n", intResult, 2)

	// [[], [1], [2], [2], [], [1], [2], []]
	// Output : [null, true, false, true, 2, true, false, 2]
}
