package atCoder

import (
	"crypto/rand"
	"fmt"
)

func ShortestPath() {
	row := 5
	col := 5

	cnt := createGrid(row, col)
	fmt.Println("shortest path: ", cnt)
}

func createGrid(row, col int) int {
	grid := [][]string{}

	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			n := rand.Intn()
		}
	}

	return 0
}

func searchPath() {}
