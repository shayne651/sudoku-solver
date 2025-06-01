package sudoku

import "testing"

func Test_ValidateRowsAllowZeros_Pass(t *testing.T) {
	var validPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, 0, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 8, 0, 0, 0, 5, 0, 0, 0},
	}
	isRowValid := validateRow(validPuzzle, 3, false)

	if !isRowValid {
		t.Error("\nValidated row should be valid when zeros do not fail")
	}
}

func Test_ValidateRowsAllowZeros_Fail(t *testing.T) {
	var unsolvablePuzzle = [][]int{
		{5, 1, 6, 8, 4, 9, 7, 1, 5},
		{3, 0, 7, 6, 0, 5, 0, 0, 0},
		{8, 0, 9, 7, 0, 0, 0, 6, 0},
		{1, 3, 5, 0, 6, 0, 9, 0, 0},
		{4, 7, 2, 5, 9, 1, 0, 0, 6},
		{9, 6, 8, 3, 7, 0, 0, 5, 0},
		{2, 5, 3, 1, 8, 6, 0, 0, 7},
		{6, 8, 4, 2, 0, 7, 5, 0, 0},
		{7, 9, 1, 0, 5, 0, 6, 0, 0},
	}
	isRowValid := validateRow(unsolvablePuzzle, 0, false)

	if isRowValid {
		t.Error("\nValidated row should not be valid (1 is used twice in the row)")
	}
}

func Test_ValidateRowsNotAllowZeros_Pass(t *testing.T) {
	isRowValid := validateRow(solvedPuzzle, 0, true)

	if !isRowValid {
		t.Error("\nValidated row should be valid (No zeros present and all values between 1-9)")
	}
}

func Test_ValidateRowsNotAllowZeros_Fail(t *testing.T) {
	var validPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, 0, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 8, 0, 0, 0, 5, 0, 0, 0},
	}
	isRowValid := validateRow(validPuzzle, 0, true)

	if isRowValid {
		t.Error("\nValidated row should not be valid (zeros erroring is enabled and zeros are present in the row)")
	}
}

func Test_ValidateColAllowZeros_Pass(t *testing.T) {
	var validPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, 0, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 8, 0, 0, 0, 5, 0, 0, 0},
	}
	isRowValid := validateCol(validPuzzle, 3, false)

	if !isRowValid {
		t.Error("\nValidated column should be valid when zeros do not fail")
	}
}

func Test_ValidateColAllowZeros_Fail(t *testing.T) {
	var unsolvablePuzzle = [][]int{
		{5, 1, 6, 8, 4, 9, 7, 1, 5},
		{3, 0, 7, 6, 0, 5, 0, 0, 0},
		{8, 0, 9, 7, 0, 0, 0, 6, 0},
		{1, 3, 5, 5, 6, 0, 9, 0, 0},
		{4, 7, 2, 5, 9, 1, 0, 0, 6},
		{9, 6, 8, 3, 7, 0, 0, 5, 0},
		{2, 5, 3, 1, 8, 6, 0, 0, 7},
		{6, 8, 4, 2, 0, 7, 5, 0, 0},
		{7, 9, 1, 0, 5, 0, 6, 0, 0},
	}
	isRowValid := validateCol(unsolvablePuzzle, 3, false)

	if isRowValid {
		t.Error("\nValidated column should not be valid (5 is used twice in row 3)")
	}
}

func Test_ValidateColNotAllowZeros_Pass(t *testing.T) {
	isRowValid := validateCol(solvedPuzzle, 0, true)

	if !isRowValid {
		t.Error("\nValidated column should be valid (No zeros present and all values between 1-9)")
	}
}

func Test_ValidateColNotAllowZeros_Fail(t *testing.T) {
	var validPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, 0, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 8, 0, 0, 0, 5, 0, 0, 0},
	}
	isRowValid := validateCol(validPuzzle, 0, true)

	if isRowValid {
		t.Error("\nValidated column should not be valid (zeros erroring is enabled and zeros are present in the row)")
	}
}

func Test_ValidateSubGridAllowZeros_Pass(t *testing.T) {
	var validPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, 0, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 8, 0, 0, 0, 5, 0, 0, 0},
	}
	isRowValid := validateSubGrid(validPuzzle, 3, 3, false)

	if !isRowValid {
		t.Error("\nValidated subgrid should be valid when zeros do not fail")
	}
}

func Test_ValidateSubGridAllowZeros_Fail(t *testing.T) {
	var unsolvablePuzzle = [][]int{
		{5, 1, 6, 8, 4, 9, 7, 1, 5},
		{3, 0, 7, 6, 0, 5, 0, 0, 0},
		{8, 0, 9, 7, 0, 0, 0, 6, 0},
		{1, 3, 5, 5, 6, 0, 9, 0, 0},
		{4, 4, 2, 5, 9, 1, 0, 0, 6},
		{9, 6, 8, 3, 7, 0, 0, 5, 0},
		{2, 5, 3, 1, 8, 6, 0, 0, 7},
		{6, 8, 4, 2, 0, 7, 5, 0, 0},
		{7, 9, 1, 0, 5, 0, 6, 0, 0},
	}
	isRowValid := validateSubGrid(unsolvablePuzzle, 3, 3, false)

	if isRowValid {
		t.Error("\nValidated subgrid should not be valid (5 is used twice in row 3)")
	}
}

func Test_ValidateSubGridNotAllowZeros_Pass(t *testing.T) {
	isRowValid := validateSubGrid(solvedPuzzle, 0, 3, true)

	if !isRowValid {
		t.Error("\nValidated column should be valid (No zeros present and all values between 1-9)")
	}
}

func Test_ValidateSubGridNotAllowZeros_Fail(t *testing.T) {
	var validPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, 0, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 8, 0, 0, 0, 5, 0, 0, 0},
	}
	isRowValid := validateSubGrid(validPuzzle, 0, 3, true)

	if isRowValid {
		t.Error("\nValidated column should not be valid (zeros erroring is enabled and zeros are present in the row)")
	}
}
