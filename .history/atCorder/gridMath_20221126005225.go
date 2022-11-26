package atCorder

import "fmt"

func GridMath() {
	grid := [][]int{
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
}

	res := gridMath(grid)
	fmt.Println("grid sum: ", res)
}

func gridMath(grid [][]int) [5][5]int {
	sumGrid := grid

	for i, r := range grid {
		for ii := range r {
			var sum = 0
			for idx, y := range r {
				sum += y
				if idx != ii { sum += grid[idx][ii] }
			}
			sumGrid[i][ii] = sum
		}
	}

	return sumGrid
}