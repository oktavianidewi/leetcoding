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
// 		this.storage[val] = len(this.data) + 1
// 		this.data = append(this.data, val)
// 		return true
// 	}
// 	return false
// }

// func (this *RandomizedSet) Remove(val int) bool {
// 	if _, ok := this.storage[val]; ok {
// 		this.storage[this.data[len(this.data)-1]] = this.storage[val]
// 		this.data[this.storage[val]-1] = this.data[len(this.data)-1]
// 		this.data = this.data[:len(this.data)-1]
// 		delete(this.storage, val)
// 		return true
// 	}
// 	return false
// }

// func (this *RandomizedSet) GetRandom() int {
// 	x := rand.Intn(len(this.storage))
// 	return this.storage[x]
// }

type RandomizedSet struct {
    mapper map[int]int
    data []int
}

func Constructor() RandomizedSet {
    
    r := RandomizedSet{}
    
    r.mapper = make(map[int]int)
    r.data = []int{}
    return r

}

func (this *RandomizedSet) Insert(val int) bool {
    if this.mapper[val] == 0 {
        this.mapper[val] = len(this.data) + 1
        this.data = append(this.data, val)
        return true
    }
    
    return false
}

func (this *RandomizedSet) Remove(val int) bool {
    if _, e := this.mapper[val]; e {
        this.mapper[this.data[len(this.data)-1]] = this.mapper[val]
        this.data[this.mapper[val]-1] = this.data[len(this.data)-1]
        this.data = this.data[:len(this.data)-1]
        delete(this.mapper, val)
        return true
    }
    return false
}

func (this *RandomizedSet) GetRandom() int {
    n := rand.Intn(len(this.data))
    return this.data[n]
}