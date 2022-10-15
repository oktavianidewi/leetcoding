package plusone

import "fmt"

func cleanTrailingZero(input []int) []int {
	length := len(input) - 1
	var clean []int
	for i := length; i > 0; i-- {
		if input[i] != 0 {
			clean = append(clean, input[i])
		}
	}
	return clean
}
func plusone(input []int) []int {
	// add checking trailing 0 in front
	input = cleanTrailingZero(input)
	length := len(input)
	var output []int

	for i := length; i >= 0; i-- {
		if i == length {
			input[i] = input[i] + 1
		}

		output = append(output, input[i])
	}
	fmt.Println(">> output ", output)
	return output
	// if length == 0 {
	// 	input[length] = input[length] + 1
	// }
	// input[lastIndex] = input[lastIndex] + 1
	// if input[lastIndex] >= 10 {
	// 	input[lastIndex] = 0
	// }

	return input
}
