package main

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	prevNums := make([]int, rowIndex)
	lenRow := 0
	pascal := _getRow(lenRow, rowIndex, prevNums)
	return pascal
}

func _getRow(lenRow, rowIndex int, prevNums []int) []int {
	lenRowPascal := lenRow + 1
	var currNums = make([]int, lenRowPascal)

	// exit loop
	if lenRow > rowIndex {
		return prevNums
	}

	// pascal main func
	for i := 0; i < lenRowPascal; i++ {

		// if empty len(prevNums)
		if len(prevNums) == 1 {
			currNums[i] = 1
		} else {
			// set 1 on the first and last index
			if i == 0 || i == lenRowPascal-1 {
				currNums[i] = 1
			} else {
				currNums[i] = prevNums[i] + prevNums[i-1]
			}
		}
	}

	// recurring loops
	return _getRow(lenRow+1, rowIndex, currNums)
}
