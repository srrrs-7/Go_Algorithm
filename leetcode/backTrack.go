package leetcode

import "fmt"

func BackTrack() {
	board := [][]byte{[]byte{'A', 'B', 'C', 'E'}, []byte{'S', 'F', 'C', 'S'}, []byte{'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	res := exist(board, word)

	fmt.Println("word search", res)
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backtrack(board, visited, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, visited [][]bool, i, j int, word string, k int) bool {
	if board[i][j] != word[k] {
		return false
	} else if k == len(word)-1 {
		return true
	}
	visited[i][j] = true
	defer func() { visited[i][j] = false }()
	if i > 0 && !visited[i-1][j] && backtrack(board, visited, i-1, j, word, k+1) {
		return true
	}
	if i < len(board)-1 && !visited[i+1][j] && backtrack(board, visited, i+1, j, word, k+1) {
		return true
	}
	if j > 0 && !visited[i][j-1] && backtrack(board, visited, i, j-1, word, k+1) {
		return true
	}
	if j < len(board[0])-1 && !visited[i][j+1] && backtrack(board, visited, i, j+1, word, k+1) {
		return true
	}
	return false
}
