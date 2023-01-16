package atCoder

import (
	"fmt"
	"math/rand"
	"time"
)

func ShortestPath() {
	row := 5
	col := 5

	grid := createGrid(row, col)
	s, g := findStartGoal(grid)

	fmt.Println("shortest path: ", cnt)
}

// init function
func createGrid(row, col int) [][]string {
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

	return grid
}

func findStartGoal(grid [][]string) (s, g []int) {
	s = make([]int, 1)
	g = make([]int, 1)

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
	return s, g
}

func searchPath() int {
	s, g := findStartGoal()
	// search path
	count := search(s, g, grid)

	return 7
}

func search(s, g []int, grid [][]string) (pos int) {
	go top(s, g, grid[x-1][y])
	go right(s, g, grid[x][y+1])
	go bottom(s, g, grid[x+1][y])
	go left(s, g, grid[x][y-1])
}

func top(s, g []int, grid [][]string) [][]string {}

func right(s, g []int, grid [][]string) [][]string {}

func bottom(s, g []int, grid [][]string) [][]string {}

func left(s, g []int, grid [][]string) [][]string {}
