package main

import "fmt"

// XOR https://stackoverflow.com/questions/14526584/what-does-the-xor-operator-do; XOR ini logical operator, jika dicari padanannya dengan math operator, XOR = ^ atau addition modulo 2 in bit
/*
func subsetXORSum(nums []int) int {
	dict := make([][]int, 1)
	fmt.Printf("initial dict %v \n", dict)
	for _, n := range nums {
		fmt.Println(n) // 5; 1; 6
		for _, arr := range dict {
			fmt.Printf("arr %v \n", arr)     // []; [] [5]
			temp := append([]int{n}, arr...) // [5] di-append dengan [] = [5]; [1] [1 5]
			fmt.Printf("temp %v \n", temp)
			dict = append(dict, temp) // [[], [5], [1], [1,5]]
			fmt.Printf("dict %v \n", dict)
		}
	}
	var ans int
	for _, arr := range dict {
		var temp int
		for _, n := range arr {
			temp = temp ^ n
		}
		ans += temp
	}
	return ans
}
*/

func subsetXORSum(nums []int) int {
	// di-jelentrehkan. gimana proses jelentrehkannya?
	// construct a set of combination of number (order doesn't matter) => [1,6] = [6,1]
	maps := make([][]int, 1)
	// fmt.Println(maps)
	for _, i := range nums {
		var temp []int
		for _, arr := range maps {
			temp = append([]int{i}, arr...)
			maps = append(maps, temp)
		}
		// fmt.Println(maps)
	}
	// sum xor process
	var res int
	for _, arr := range maps {
		tmp := 0
		for _, num := range arr {
			tmp = tmp ^ num
		}
		res += tmp
	}
	return res
}

/*
// 0ms solution example
func subsetXORSum(nums []int) int {
	sum := 0
	subs := []int{} // subset xor

	for i := 0; i < len(nums); i++ {
		a := nums[i]
		prevLen := len(subs)
		for j := 0; j < prevLen; j++ {
			b := subs[j] ^ a
			sum += b
			subs = append(subs, b)
		}
		subs = append(subs, a)
		sum += a
	}
	return sum
}
*/

func main() {
	// nums, exp := []int{1, 3}, 6
	// nums, exp := []int{5, 1, 6}, 28
	nums, exp := []int{2, 4, 4}, 24
	res := subsetXORSum(nums)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
