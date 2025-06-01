# Sudoku Solver
### Prerequisites

One of either go or docker will be required depending if you want to build and run it or just build
* [Go](https://go.dev/dl/) (If you want to build the code and run it)
* [Docker](https://docs.docker.com/get-started/get-docker/) (If you just want to run it)


## Build and run

1. Go has been installed

2. Clone the repo

```sh

git clone git@github.com:shayne651/sudoku-solver.git

```

3. Open puzzle.json in the main directory and enter the sudoku puzzle you want to solve

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

go run ./cmd/sudoku-solver/main.go 

```



## Run the pre-built code

1. Docker has been installed

2. Clone the repo

```sh

git clone git@github.com:shayne651/sudoku-solver.git

```

3. Open puzzle.json in the main directory and enter the sudoku puzzle you want to solve

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

docker compose -f ./build/compose.yml up

```
