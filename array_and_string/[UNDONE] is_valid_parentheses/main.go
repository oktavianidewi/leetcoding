package main

import "fmt"

// solusiku mirip ini: https://leetcode.com/problems/valid-parentheses/discuss/2158788/Golang

var validBracketMap = map[rune]rune{
	rune('('): ')',
	rune('{'): '}',
	rune('['): ']',
}

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}
	var bracketStack []rune
	isValidFlag := true
	for i := 0; i < len(s); i++ {
		if _, isOpenBracket := validBracketMap[rune(s[i])]; isOpenBracket {
			bracketStack = append(bracketStack, rune(s[i]))
		} else {
			if len(bracketStack) == 0 {
				return false
			}
			closeBracket := validBracketMap[rune(bracketStack[len(bracketStack)-1])]
			fmt.Printf("bracketStack popped %v \n", bracketStack)

			if rune(s[i]) == closeBracket {
				isValidFlag = isValidFlag && true

				if len(bracketStack)-1 > 0 {
					bracketStack = bracketStack[:len(bracketStack)-1]
				} else {
					bracketStack = nil
				}
			} else {
				return false
			}
		}
		// fmt.Printf("key %v, val %v \n", i, s[i])
		if len(bracketStack) > 0 {
			return false
		}
	}
	return isValidFlag
}

func main() {
	// error
	// s, expected := "([)]", false
	s, expected := "{[]}", true
	// s, expected := "[", false
	// s, expected := "((", false

	// test-case
	// s, expected := "()", true
	// s, expected := "()[]{}", true
	// s, expected := "(())", true
	// s, expected := "(()}", false
	// s, expected := "(]", false

	result := isValid(s)
	fmt.Printf("expected %v, result %v \n", expected, result)
}
