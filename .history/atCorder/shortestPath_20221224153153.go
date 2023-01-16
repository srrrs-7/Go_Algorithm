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
	for x := range grid {
		for y := range grid
	}
	return 7
}

func top(grid [][]int) int {
	if grid
}

func left(grid [][]int) int {}

func bottom(grid [][]int) int {}

func right(grid [][]int) int {}
