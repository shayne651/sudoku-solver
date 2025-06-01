package main

import (
	"reflect"
	"testing"
)

var solvedPuzzle = [][]int{
	{5, 6, 2, 7, 8, 9, 1, 3, 4},
	{3, 7, 8, 4, 2, 1, 5, 6, 9},
	{9, 1, 4, 3, 5, 6, 7, 8, 2},
	{4, 3, 1, 5, 7, 8, 2, 9, 6},
	{6, 2, 7, 9, 1, 4, 3, 5, 8},
	{8, 5, 9, 2, 6, 3, 4, 1, 7},
	{7, 9, 5, 6, 3, 2, 8, 4, 1},
	{1, 4, 3, 8, 9, 7, 6, 2, 5},
	{2, 8, 6, 1, 4, 5, 9, 7, 3},
}

func Test_SolverValidInput(t *testing.T) {
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
	solved, err := solvePuzzle(validPuzzle, 0, 0)

	if err != nil {
		t.Error("Error retreived when the puzzle was valid and solvable", err)
	}

	if !reflect.DeepEqual(solved, solvedPuzzle) {
		t.Error("Output should be\n", solvedPuzzle, "\nbut was\n", solved)
	}
}

func Test_SolverInvalid_NumberGT9(t *testing.T) {
	var invalidPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, 0, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 8, 0, 0, 0, 5, 0, 0, 10},
	}

	solved, err := solvePuzzle(invalidPuzzle, 0, 0)
	if err != nil && err.Error() != "invalid puzzle: element is greater than 9 at row: 9, column: 9, value: 10" {
		t.Error("Error should be:\n invalid puzzle: element is greater than 9 at row: 9, column: 9, value: 10\nBut was: \n", err)
	}

	if solved != nil {
		t.Error("solved an invalid puzzle")
	}
}

func Test_SolverInvalid_NumberLT0(t *testing.T) {
	var invalidPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, -1, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 8, 0, 0, 0, 5, 0, 0, 9},
	}

	solved, err := solvePuzzle(invalidPuzzle, 0, 0)
	if err != nil && err.Error() != "invalid puzzle: element is less than 0 at row: 5, column: 5, value: -1" {
		t.Error("Error should be:\n invalid puzzle: element is less than 0 at row: 5, column: 5, value: -1\nBut was: \n", err)
	}

	if solved != nil {
		t.Error("solved an invalid puzzle")
	}
}

func Test_SolverInvalid_TooManyRows(t *testing.T) {
	var invalidPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, 0, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 8, 0, 0, 0, 5, 0, 0, 9},
		{2, 8, 0, 0, 0, 5, 0, 0, 9},
	}

	solved, err := solvePuzzle(invalidPuzzle, 0, 0)
	if err != nil && err.Error() != "puzzle should have a row length of 9 but had row length of 10" {
		t.Error("puzzle should have a row length of 9 but had row length of 10\nBut was: \n", err)
	}

	if solved != nil {
		t.Error("solved an invalid puzzle")
	}
}

func Test_SolverInvalid_TooFewRows(t *testing.T) {
	var invalidPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, 0, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	solved, err := solvePuzzle(invalidPuzzle, 0, 0)
	if err != nil && err.Error() != "puzzle should have a row length of 9 but had row length of 8" {
		t.Error("puzzle should have a row length of 9 but had row length of 8\nBut was: \n", err)
	}

	if solved != nil {
		t.Error("solved an invalid puzzle")
	}
}

func Test_SolverInvalid_TooManyCols(t *testing.T) {
	var invalidPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0, 9},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, 0, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 8, 0, 0, 0, 5, 0, 0, 9},
	}

	solved, err := solvePuzzle(invalidPuzzle, 0, 0)
	if err != nil && err.Error() != "puzzle should have a column length of 9 but had column length of 10" {
		t.Error("puzzle should have a column length of 9 but had column length of 10\nBut was: \n", err)
	}

	if solved != nil {
		t.Error("solved an invalid puzzle")
	}
}

func Test_SolverInvalid_TooFewCols(t *testing.T) {
	var invalidPuzzle = [][]int{
		{0, 0, 0, 7, 8, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 6, 9},
		{0, 0, 4, 0, 5, 6, 0, 0, 2},
		{0, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 2, 0, 9, 0, 4, 0, 0, 0},
		{0, 0, 9, 2, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 8, 4, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	solved, err := solvePuzzle(invalidPuzzle, 0, 0)
	if err != nil && err.Error() != "puzzle should have a row length of 9 but had row length of 8" {
		t.Error("puzzle should have a row length of 9 but had row length of 8\nBut was: \n", err)
	}

	if solved != nil {
		t.Error("solved an invalid puzzle")
	}
}

func Test_SolverUnsolvablePuzzle(t *testing.T) {
	var unsolvablePuzzle = [][]int{
		{5, 1, 6, 8, 4, 9, 7, 3, 5},
		{3, 0, 7, 6, 0, 5, 0, 0, 0},
		{8, 0, 9, 7, 0, 0, 0, 6, 0},
		{1, 3, 5, 0, 6, 0, 9, 0, 0},
		{4, 7, 2, 5, 9, 1, 0, 0, 6},
		{9, 6, 8, 3, 7, 0, 0, 5, 0},
		{2, 5, 3, 1, 8, 6, 0, 0, 7},
		{6, 8, 4, 2, 0, 7, 5, 0, 0},
		{7, 9, 1, 0, 5, 0, 6, 0, 0},
	}
	solved, err := solvePuzzle(unsolvablePuzzle, 0, 0)

	if err != nil {
		t.Error("Error retreived when the puzzle was valid", err)
	}

	if solved != nil {
		t.Error("Unsolvable puzzle was solved")
	}
}

func Test_FindMissingValues(t *testing.T) {
	var validPuzzle1 = [][]int{
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
	correctValues := []int{1, 7}
	validValues := findMissingValues(validPuzzle1, 1, 1)
	if !reflect.DeepEqual(validValues, correctValues) {
		t.Errorf("\nError getting valid missing values\nExpected:\n%v\nActual:\n%v", correctValues, validValues)
	}
}

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
