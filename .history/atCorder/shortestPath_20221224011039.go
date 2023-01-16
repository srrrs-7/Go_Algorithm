package atCoder

import (
	"fmt"
	"math/rand"
)

func ShortestPath() {
	row := 5
	col := 5

	cnt := createGrid(row, col)
	fmt.Println("shortest path: ", cnt)
}

func createGrid(row, col int) int {
	grid := make([][]string, row*col)

	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			n := rand.Intn(2)
			grid[x][y] = fmt.Sprint(n)
		}
		grid[0][2] = "S"
		grid[4][3] = "G"
	}

	return 0
}

func searchPath() {}
