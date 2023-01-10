package main

// naive approach
func checkIfExist(arr []int) bool {
	var half float32
	left, right := 0, len(arr)-1
	pointer := 0
	for left <= right && pointer < right {
		half = float32(arr[pointer]) / 2
		// fmt.Printf("half %v, %v \n", half, arr[left])
		// time.Sleep(1 * time.Second)

		if half == float32(arr[left]) && half != 0 {
			// fmt.Printf("sama %v, %v, %v, %v \n", half, arr[left], pointer, left)
			return true
		} else if half == float32(arr[left]) && half == 0 {
			// fmt.Printf("half = 0, %v, %v \n", pointer, left)
			if left != pointer {
				return true
			}
		}

		if left == right {
			pointer++
			left = -1
		}

		left++
	}
	return false
}
