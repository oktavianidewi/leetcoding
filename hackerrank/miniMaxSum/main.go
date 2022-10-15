package main

import "fmt"

func miniMaxSum(arr []int32) {
	// Write your code here
	// get min and max value from array, while summing
	var sumMid, min, max int64
	min = int64(arr[0])
	max = int64(arr[0])
	for i := 0; i <= len(arr)-1; i++ {
		if i < len(arr)-1 {
			// get the max
			if max < int64(arr[i+1]) {
				max = int64(arr[i+1])
			}
			// get the min
			if min > int64(arr[i+1]) {
				min = int64(arr[i+1])
			}
		}
		sumMid = sumMid + int64(arr[i])
	}
	minSumVal := sumMid - max
	maxSumVal := sumMid - min
	fmt.Printf("%v %v\n", minSumVal, maxSumVal)
}

func main() {
	arr := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	miniMaxSum(arr)
	// 2063136757 2744467344
}
