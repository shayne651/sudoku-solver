package sudoku

import "fmt"

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
