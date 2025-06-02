package sudoku

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

func getCol(colNum int, puzzle [][]int) []int {
	var col []int
	for i := range 9 {
		col = append(col, puzzle[i][colNum])
	}
	return col
}

func findMissingValues(puzzle [][]int, rowIndex, colIndex int) []int {
	// tracks if numbers have been used in the row, column, subgrid
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
	// If they have not been used they are valid
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

		if solved := recurseBacktrack(puzzle, nextRow, nextCol); solved != nil {
			return solved
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

	// Locate starting position
	var firstZeroRow, firstZeroCol int
	zeroFound := false

	for i := range 9 {
		for j := range 9 {
			if puzzle[i][j] == 0 {
				firstZeroRow, firstZeroCol = i, j
				zeroFound = true
				break
			}
		}
		if zeroFound {
			break
		}
	}

	if !zeroFound {
		return puzzle, nil
	}

	validValues := findMissingValues(puzzle, firstZeroRow, firstZeroCol)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	solutionChan := make(chan [][]int, 1)
	var once sync.Once

	nextRow, nextCol := firstZeroRow, firstZeroCol+1
	if firstZeroCol == 8 {
		nextRow++
		nextCol = 0
	}

	// Start a goroutine for each valid starting value
	// Each routine will preform backtracking
	for _, validValue := range validValues {
		wg.Add(1)
		go func(validValue int) {
			defer wg.Done()
			newPuzzle := make([][]int, 9)
			for i := 0; i < 9; i++ {
				newPuzzle[i] = make([]int, 9)
				copy(newPuzzle[i], puzzle[i])
			}
			newPuzzle[firstZeroRow][firstZeroCol] = validValue
			solution := recurseBacktrack(newPuzzle, nextRow, nextCol)
			if solution != nil {
				select {
				case <-ctx.Done():
					return
				default:
					// Allows solutionsChan to be written to once
					// Cancels other goroutines
					once.Do(func() {
						cancel()
						solutionChan <- solution
					})
				}
			}
		}(validValue)

	}

	// Waits for routine to finish, closes channel
	wg.Wait()
	close(solutionChan)
	var solution [][]int
	select {
	case solution = <-solutionChan:
		return solution, nil
	case <-ctx.Done():
		return nil, nil
	}
}
