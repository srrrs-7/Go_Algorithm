package atCoder

import (
	"fmt"
	"math/rand"
	"time"
)

func ShortestPath() {
	row := 5
	col := 5

	cnt := createGrid(row, col)
	fmt.Println("shortest path: ", cnt)
}

// init function
func createGrid(row, col int) int {
	grid := make([][]string, row)

	for x := 0; x < row; x++ {
		grid[x] = make([]string, col)
		for y := 0; y < col; y++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(2)
			grid[x][y] = fmt.Sprint(n)
		}
	}
	grid[0][2] = "S"
	grid[4][3] = "G"

	return 0
}

func searchPath(grid [][]string) int {
	s := make([][]int, 1)
	g := make([][]int, 1)

	// find start and goal
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == "S" {
				s = grid[x][y]
			}
			if grid[x][y] == "G" {
				s = grid[x][y]
			}
		}
	}
	// search path
	search()

	return 7
}

func search(grid, s, g [][]int) (count int) {
	top(grid, s, g)
	left(grid, s, g)
	bottom(grid, s, g)
	right(grid, s, g)
}

func top(grid, s, g [][]int) int {}

func left(grid, s, g [][]int) int {}

func bottom(grid, s, g [][]int) int {}

func right(grid, s, g [][]int) int {}
