package test

import "fmt"

func Solution6() {
	matrix := [][]int{
		{3, 4, 4},
		{4, 3, 4},
		{3, 2, 4},
		{2, 2, 2},
		{3, 3, 4},
		{1, 4, 4},
		{4, 1, 1},
	}
	cnt := solution6(matrix)
	fmt.Println(cnt)
}

func solution6(matrix [][]int) int {
	cnt := 0
	row := len(matrix)
	col := len(matrix[0])

	if row == 0 || col == 0 {
		return 0
	}

	// matrix copy logic
	matrix2 := make([][]int, len(matrix))
	for i := range matrix {
		matrix2[i] = make([]int, len(matrix[i]))
		copy(matrix2[i], matrix[i])
	}

	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if matrix2[x][y] >= 0 {
				checkNeighbor(matrix, matrix2, x, y, row, col) // matrix[x][y] = -1
				cnt++
			}
		}
	}
	matrix2[0][0] = 1
	matrix[1][1] = 1
	fmt.Printf("%p\n", &matrix)
	fmt.Printf("%p\n", &matrix2)
	fmt.Println(matrix)
	fmt.Println(matrix2)
	return cnt
}

func checkNeighbor(A, B [][]int, x, y, row, col int) {
	if B[x][y] == -1 {
		return
	}
	B[x][y] = -1

	// check north
	if x-1 >= 0 && A[x-1][y] == A[x][y] {
		checkNeighbor(A, B, x-1, y, row, col)
	}

	// check east
	if y+1 < col && A[x][y+1] == A[x][y] {
		checkNeighbor(A, B, x, y+1, row, col)
	}
	// check south
	if x+1 < row && A[x+1][y] == A[x][y] {
		checkNeighbor(A, B, x+1, y, row, col)
	}
	// check west
	if y-1 >= 0 && A[x][y-1] == A[x][y] {
		checkNeighbor(A, B, x, y-1, row, col)
	}
}
