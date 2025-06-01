package sudoku

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func createTempPuzzleFile(t *testing.T, data [][]int) string {
	t.Helper()

	filePath := filepath.Join(".", "puzzle.json")

	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create puzzle.json: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		t.Fatalf("Failed to write JSON to puzzle.json: %v", err)
	}

	return filePath
}

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
	filePath := createTempPuzzleFile(t, validPuzzle)
	defer os.Remove(filePath)
	solved, err := SolvePuzzle()

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

	filePath := createTempPuzzleFile(t, invalidPuzzle)
	defer os.Remove(filePath)
	solved, err := SolvePuzzle()
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

	filePath := createTempPuzzleFile(t, invalidPuzzle)
	defer os.Remove(filePath)
	solved, err := SolvePuzzle()
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

	filePath := createTempPuzzleFile(t, invalidPuzzle)
	defer os.Remove(filePath)
	solved, err := SolvePuzzle()
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

	filePath := createTempPuzzleFile(t, invalidPuzzle)
	defer os.Remove(filePath)
	solved, err := SolvePuzzle()
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

	filePath := createTempPuzzleFile(t, invalidPuzzle)
	defer os.Remove(filePath)
	solved, err := SolvePuzzle()
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

	filePath := createTempPuzzleFile(t, invalidPuzzle)
	defer os.Remove(filePath)
	solved, err := SolvePuzzle()
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

	filePath := createTempPuzzleFile(t, unsolvablePuzzle)
	defer os.Remove(filePath)
	solved, err := SolvePuzzle()

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
