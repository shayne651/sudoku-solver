# Sudoku Solver
## Prerequisites

One of either go or docker will be required depending if you want to build and run it or just run it: 
* [Go](https://go.dev/dl/) - If you want to build the code and run it
* [Docker](https://docs.docker.com/get-started/get-docker/) - If you just want to run it

## Approach
- To start designing this Sudoku solver, I went through how I would step by step go through solving a Sudoku puzzle. 
	1. Look at the first empty cell
	2. Check what numbers could be used here (based on row, column and subgrid restrictions)
	3. Choose one of the valid numbers and try it
	4. Move on to the next
	5. If there is a conflict where there are no valid options for a cell, I would then go back and look at other cells in the row/column/subgrid to see if I would be able to change any of the values and make a valid option for the current cell
	6. Repeat steps
- This approach is straight forward and possible for computers to do but step 5 poses a problem. How can I look back at all of the other options and see which numbers I could change to make it valid. Not only would I have to analyze each cell in the row, column, and subgrid, but for each change I would then have to verify all rows, columns and subgrids again to ensure they are all still valid. 
- This thought process led me to the backtracking algorithm. An approach where you can use recursion to check each value and try to solve the puzzle based on that input. You can then go back to a previous state and try again with a new value if it does not work.

## Backtracking
- The naive backtracking algorithm is essentially a brute force technique. It will solve the problem but it is very costly to do so. Is the worst case scenario assuming the correct solution is the last one we try, that could leave us with 9^n iterations where n is the number of empty cells. 
- By adding some simple logic we could cut this down significantly by filtering out what values are invalid in the cell you want to try. Eg, if puzzle\[0]\[1]=5 then 5 is not valid for puzzle\[0]\[x] where x can be any value 0-8.

## Implementation
- One of the important parts of making a backtracking algorithm is ensuring the recursive base case is well defined to ensure it does not create an infinite loop. For a Sudoku solver we would be able to check if the base case is completed by checking what row we are on. If we are on index 9 for the row puzzle that means 0-8 have been solved already and there is no more to do
- Row, column and subgrid verification were done in separate methods to increase readability and maintainability. Although this can create more loops being run, the Sudoku puzzle is relatively small at the size of 9x9, any performance gain we would have by combining a loop would be negligible. 
- Goroutines are used to parallelize solving starting from the first empty cell. I locate the first cell containing a 0, determine all valid values for it, and then launch a goroutine for each valid option. I limited goroutine usage to this first step to avoid spawning thousands of concurrent routines. The initial value has the largest impact due to its higher branching factor. While parallelizing deeper levels could bring benefits, it must be done cautiously to avoid overwhelming the system with concurrent calls.

## Build and run

1. Go has been installed

2. Clone the repo

```sh

git clone git@github.com:shayne651/Sudoku-solver.git

```

3. Open puzzle.json in the main directory and enter the Sudoku puzzle you want to solve

```json

[
	[0, 0, 0, 7, 8, 0, 0, 0, 0],
	[0, 0, 0, 0, 0, 0, 5, 6, 9],
	[0, 0, 4, 0, 5, 6, 0, 0, 2],
	[0, 3, 0, 0, 7, 0, 0, 0, 0],
	[6, 2, 0, 9, 0, 4, 0, 0, 0],
	[0, 0, 9, 2, 0, 0, 0, 1, 0],
	[0, 0, 0, 0, 0, 0, 8, 4, 1],
	[0, 0, 0, 0, 0, 0, 0, 0, 0],
	[2, 8, 0, 0, 0, 5, 0, 0, 0]
]

```

4. Run the program

```sh

go run ./cmd/Sudoku-solver/main.go 

```

## Run the pre-built code

1. Docker has been installed

2. Clone the repo

```sh

git clone git@github.com:shayne651/Sudoku-solver.git

```

3. Open puzzle.json in the main directory and enter the Sudoku puzzle you want to solve

```json

[
	[0, 0, 0, 7, 8, 0, 0, 0, 0],
	[0, 0, 0, 0, 0, 0, 5, 6, 9],
	[0, 0, 4, 0, 5, 6, 0, 0, 2],
	[0, 3, 0, 0, 7, 0, 0, 0, 0],
	[6, 2, 0, 9, 0, 4, 0, 0, 0],
	[0, 0, 9, 2, 0, 0, 0, 1, 0],
	[0, 0, 0, 0, 0, 0, 8, 4, 1],
	[0, 0, 0, 0, 0, 0, 0, 0, 0],
	[2, 8, 0, 0, 0, 5, 0, 0, 0]
]

```

4. Run the container

```sh

docker compose -f ./build/compose.yml up

```