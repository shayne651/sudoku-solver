package main

import (
	"fmt"

	sudoku "github.com/shayne651/sudoku-solver/internal/sudoku"
)

func main() {

	solved, err := sudoku.SolvePuzzle()

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
