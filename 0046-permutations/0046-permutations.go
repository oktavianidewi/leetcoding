func permute(nums []int) [][]int {
	r := len(nums)
	if r == 1 {
		//Convert every item in L to List and
		//Append it to List of List
		temp := make([][]int, 0)
		for _, rr := range nums {
			t := make([]int, 0)
			t = append(t, rr)
			temp = append(temp, [][]int{t}...)
		}
		return temp
	} else {
		res := make([][]int, 0)
		for i := 0; i < len(nums); i++ {
			//Create List Without L[i] element
			perms := make([]int, 0)

			// fmt.Printf("i %v \n", i)

			perms = append(perms, nums[:i]...)
			// fmt.Printf("first perms %v \n", perms)

			perms = append(perms, nums[i+1:]...)
			// fmt.Printf("second perms %v \n", perms)

			//Call recursively to Permutations
			for _, x := range permute(perms) {
				t := append(x, nums[i])
				res = append(res, [][]int{t}...)
			}
		}
		return res
	}
}