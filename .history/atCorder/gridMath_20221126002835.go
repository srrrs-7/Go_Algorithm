package atCorder

func GridMath() {
	grid := [][]int{
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
		{1,2,3,4,5},
}

	res := grigMath(grid)
}

func gridMath(grid [][]int) [][]int {
	sumGrid := [][]int{}

	for i,r := range grid {
		for ii := range r {
			var sum = 0
			for _, y := range r { sum += y }
			sumGrid[i][ii] = sum
		}
	}

	return
}