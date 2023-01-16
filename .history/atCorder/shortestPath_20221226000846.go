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
	go top(s, g, grid[s[0]-1][s[1]])
	go right(s, g, grid[s[0]][s[1]+1])
	go bottom(s, g, grid[s[0]+1][s[1]])
	go left(s, g, grid[s[0]][s[1]-1])
}

func top(sx, sy, gx, gy int, grid [][]string) int {}

func right(sx, sy, gx, gy int, grid [][]string) int {}

func bottom(sx, sy, gx, gy int, grid [][]string) int {}

func left(sx, sy, gx, gy int, grid [][]string) int {}
