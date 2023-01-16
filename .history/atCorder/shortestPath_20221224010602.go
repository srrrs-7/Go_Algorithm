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
	grid := [5]string{[5]string{}}

	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			n := rand.Intn(1)

		}
	}

	return 0
}

func searchPath() {}
