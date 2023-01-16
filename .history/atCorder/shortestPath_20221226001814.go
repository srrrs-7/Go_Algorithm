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
	sx, sy, gx, gy := findStartGoal(grid)

	cnt := 0

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

func findStartGoal(grid [][]string) (sx, sy, gx, gy int) {
	// find start and goal
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == "S" {
				sx, sy = x, y
			}
			if grid[x][y] == "G" {
				gx, gy = x, y
			}
		}
	}
	return sx, sy, gx, gy
}

func search(sx, sy, gx, gy int, grid [][]string) {
	go top(sx, sy, gx, gy, grid)
	go right(sx, sy, gx, gy, grid)
	go bottom(sx, sy, gx, gy, grid)
	go left(sx, sy, gx, gy, grid)
}

func top(sx, sy, gx, gy int, grid [][]string) bool {
	if sx-1 < 0 {
		return false
	}
	if grid[sx-1][sy] == "1" {
		return false
	}
	if grid[sx-1][sy] == "G" {
		return true
	}
	top(sx+1, sy, gx, gy, grid)
	return true
}

func right(sx, sy, gx, gy int, grid [][]string) bool {}

func bottom(sx, sy, gx, gy int, grid [][]string) bool {}

func left(sx, sy, gx, gy int, grid [][]string) bool {}
