package sudoku

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func getCol(colNum int, puzzle [][]int) []int {
	var col []int
	for i := range 9 {
		col = append(col, puzzle[i][colNum])
	}
	return col
}

func findMissingValues(puzzle [][]int, rowIndex, colIndex int) []int {
	usedColsMap := make(map[int]bool)
	usedRowMap := make(map[int]bool)
	usedGridMap := make(map[int]bool)

	row := puzzle[rowIndex]
	col := getCol(colIndex, puzzle)

	colStart := colIndex - (colIndex % 3)
	rowStart := rowIndex - (rowIndex % 3)

	for i := range 3 {
		for j := range 3 {
			usedGridMap[puzzle[i+rowStart][j+colStart]] = true
		}
	}
	for i := range 9 {

		usedColsMap[col[i]] = true
		usedRowMap[row[i]] = true
	}
	var validNumbers []int
	for i := range 10 {
		if i > 0 && !usedColsMap[i] && !usedRowMap[i] && !usedGridMap[i] {
			validNumbers = append(validNumbers, i)
		}
	}

	return validNumbers
}

func recurseBacktrack(puzzle [][]int, row, col int) [][]int {
	if row == 9 {
		return puzzle
	}

	nextRow, nextCol := row, col+1
	if col == 8 {
		nextRow++
		nextCol = 0
	}

	if puzzle[row][col] != 0 {
		return recurseBacktrack(puzzle, nextRow, nextCol)
	}

	validValues := findMissingValues(puzzle, row, col)
	for _, val := range validValues {
		puzzle[row][col] = val

		if validateRow(puzzle, row, false) &&
			validateCol(puzzle, col, false) &&
			validateSubGrid(puzzle, col, row, false) {

			if solved := recurseBacktrack(puzzle, nextRow, nextCol); solved != nil {
				return solved
			}
		}

		puzzle[row][col] = 0
	}

	return nil
}

func getPuzzleFromFile() ([][]int, error) {
	file, err := os.Open("puzzle.json")
	if err != nil {
		fmt.Println("Error opening puzzle file", err)
		return nil, err
	}

	defer file.Close()
	bytes, _ := io.ReadAll(file)

	var puzzle [][]int

	err = json.Unmarshal(bytes, &puzzle)
	if err != nil {
		fmt.Println("Error unmarshaling puzzle", err)
		return nil, err
	}

	return puzzle, nil
}

func SolvePuzzle() ([][]int, error) {

	puzzle, err := getPuzzleFromFile()

	if err != nil {
		return nil, err
	}

	puzzleValidityError := isPuzzleValid(puzzle)
	if puzzleValidityError != nil {
		return nil, puzzleValidityError
	}

	// var firstZeroRow, firstZeroCol int
	// zeroFound := false

	// for i := range 9 {
	// 	for j := range 9 {
	// 		if puzzle[i][j] == 0 {
	// 			firstZeroRow, firstZeroCol = i, j
	// 			zeroFound = true
	// 		}
	// 	}
	// }

	// if !zeroFound {
	// 	return puzzle, nil
	// }

	// validValues := findMissingValues(puzzle, firstZeroCol, firstZeroRow)

	// ctx,

	solved := recurseBacktrack(puzzle, 0, 0)
	return solved, nil
}
