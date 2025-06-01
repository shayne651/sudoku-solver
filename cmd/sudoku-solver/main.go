package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func validateRow(puzzle [][]int, r int, errorOnZeros bool) bool {
	usedMap := make(map[int]bool)
	for i := range puzzle[r] {
		if puzzle[r][i] > 0 && puzzle[r][i] <= 9 {
			if usedMap[puzzle[r][i]] {
				return false
			}
			usedMap[puzzle[r][i]] = true
		} else if errorOnZeros && puzzle[r][i] == 0 {
			return false
		}
	}
	return true
}

func validateCol(puzzle [][]int, c int, errorOnZeros bool) bool {
	usedMap := make(map[int]bool)
	for i := range 9 {
		if puzzle[i][c] > 0 && puzzle[i][c] <= 9 {
			if usedMap[puzzle[i][c]] {
				return false
			}
			usedMap[puzzle[i][c]] = true
		} else if errorOnZeros && puzzle[i][c] == 0 {
			return false
		}
	}
	return true
}

func validateSubGrid(puzzle [][]int, col, row int, errorOnZeros bool) bool {
	colStart := col - (col % 3)
	rowStart := row - (row % 3)
	usedMap := make(map[int]bool)

	for i := range 3 {
		for j := range 3 {
			if puzzle[i+rowStart][j+colStart] > 0 && puzzle[i+rowStart][j+colStart] <= 9 {
				if usedMap[puzzle[i+rowStart][j+colStart]] {
					return false
				}
				usedMap[puzzle[i+rowStart][j+colStart]] = true
			} else if errorOnZeros && puzzle[i+rowStart][j+colStart] == 0 {
				return false
			}
		}
	}
	return true
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

func getCol(colNum int, puzzle [][]int) []int {
	var col []int
	for i := range 9 {
		col = append(col, puzzle[i][colNum])
	}
	return col
}

func isPuzzleValid(puzzle [][]int) error {
	if len(puzzle) != 9 {
		return fmt.Errorf("puzzle should have a row length of 9 but had row length of %d", len(puzzle))
	}

	for _, col := range puzzle {
		if len(col) != 9 {
			return fmt.Errorf("puzzle should have a column length of 9 but had column length of %d", len(col))
		}
	}
	for i := range 9 {
		for j := range 9 {
			if puzzle[i][j] > 9 {
				return fmt.Errorf("invalid puzzle: element is greater than 9 at row: %d, column: %d, value: %d", i+1, j+1, puzzle[i][j])
			} else if puzzle[i][j] < 0 {
				return fmt.Errorf("invalid puzzle: element is less than 0 at row: %d, column: %d, value: %d", i+1, j+1, puzzle[i][j])
			}
		}
	}
	return nil
}

func solvePuzzle(puzzle [][]int, row, col int) ([][]int, error) {
	puzzleValidityError := isPuzzleValid(puzzle)
	if puzzleValidityError != nil {
		return nil, puzzleValidityError
	}

	solved := recurseBacktrack(puzzle, 0, 0)
	return solved, nil
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

func main() {
	puzzle, err := getPuzzleFromFile()

	if err != nil {
		return
	}

	solved, err := solvePuzzle(puzzle, 0, 0)

	if err != nil {
		fmt.Println(err)
	}
	if solved != nil {
		fmt.Println("Solved Puzzle:")
		for _, row := range solved {
			fmt.Println(row)
		}
	} else if err == nil && solved == nil {
		fmt.Println("No solution found.")
	}
}
