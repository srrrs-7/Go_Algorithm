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
	s := make([]int, 1)
	g := make([]int, 1)

	// find start and goal
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == "S" {
				s = []int{x, y}
			}
			if grid[x][y] == "G" {
				s = []int{x, y}
			}
		}
	}
	// search path
	count := search(s, g, grid)

	return 7
}

func search(s, g []int, grid [][]string) (count int) {
	for x := range grid {
		for y := range grid[x] {
			top(s, g, grid[x][y])
			right(s, g, grid[x][y])
			bottom(s, g, grid[x][y])
			left(s, g, grid[x][y])
		}
	}


}

func top(s, g []int, grid [][]string) [][]string {
	if grid
}

func right(s, g []int, grid [][]string) [][]string {}

func bottom(s, g []int, grid [][]string) [][]string {}

func left(s, g []int, grid [][]string) [][]string {}
