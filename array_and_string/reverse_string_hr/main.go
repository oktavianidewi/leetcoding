package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ModifyString(str string) string {
	// remove spaces
	// strArr := strings.Fields(str) // return array of string
	strArr := strings.TrimSpace(str) // return string

	// pilihan:
	// convert from string to array
	// pakai temp variable untuk mutasi string

	tempStr := make([]string, len(strArr))

	i, j := 0, len(strArr)-1
	// remove int
	for i <= j {
		// detect a character is integer or alphabet, alphabet returns err != nil
		_, err_i := strconv.Atoi(string(strArr[i]))
		if err_i != nil {
			tempStr = append(tempStr, string(strArr[i]))
		}
		i++
	}

	// restart i
	i = 0
	j = len(tempStr) - 1
	for i < j {
		fmt.Printf("compare %v and %v \n", tempStr[i], tempStr[j])
		tempStr[i], tempStr[j] = tempStr[j], tempStr[i]
		i++
		j--
	}
	return strings.Join(tempStr, "")
}

// https://yourbasic.org/golang/string-functions-reference-cheat-sheet/
