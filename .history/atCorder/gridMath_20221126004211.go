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

func gridMath(grid [][]int) [][]int {
	sumGrid := make([][]int, 5)

	for i,r := range grid {
		for ii := range r {
			var sum = 0
			for idx, y := range r {
				sum += y
				sum += grid[0][idx]
			}
			sumGrid[i][ii] = sum
		}
	}
	return sumGrid
}