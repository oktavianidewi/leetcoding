// func sortedSquares(nums []int) []int {
// 	sqs := []int{}
// 	left, right := 0, len(nums)-1

// 	for left <= right {
// 		// square
// 		sq := nums[left] * nums[left]
// 		// fmt.Printf("sq %v \n", sq)

// 		if len(sqs) == 0 {
// 			sqs = append(sqs, sq)
// 		} else {
// 			// compare and insert
// 			l, r := 0, len(sqs)-1
// 			for l <= r {
// 				if sq <= sqs[l] {
// 					sqs = insertNumInArr(sqs, l, sq)
// 					break
// 				} else if l == r && sq > sqs[r] {
// 					sqs = append(sqs, sq)
// 					break
// 				}
// 				l += 1
// 				// time.Sleep(1 * time.Second)
// 			}
// 		}
// 		left += 1
// 		// fmt.Printf("sqs %v \n", sqs)
// 		// time.Sleep(1 * time.Second)
// 	}
// 	return sqs
// }

// func insertNumInArr(nums []int, index, num int) []int {
// 	// extend array space
// 	nums = append(nums, 0)
// 	// fmt.Printf("nums index spec %v , %v, %v, %v \n", nums[index+1:], nums[index:], len(nums), index)
// 	// copy(dest, src)

// 	copy(nums[index+1:], nums[index:])
// 	nums[index] = num
// 	// fmt.Printf("nums after copy %v \n", nums)
// 	return nums
// }

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
		if nums[i] < 0 {
			nums[i] = -1 * nums[i]
		}
	}
	sort.Ints(nums)
	return nums
}