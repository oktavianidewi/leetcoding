package main

// you can also use imports, for example:
import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func readFile() []int {
	inputFile, _ := ioutil.ReadFile("test-input.txt")
	var inputs []int

	for _, input := range inputFile {
		intInput, _ := strconv.Atoi(string(input))
		inputs = append(inputs, intInput)
	}
	return inputs
}

func Solution(A []int) int {
	// write your code in Go 1.4
	fmt.Println(A)
	return 1
}

func main() {
	getInput := readFile()
	Solution(getInput)
}
