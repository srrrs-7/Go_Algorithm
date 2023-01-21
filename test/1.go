package test

func Solution1(A []int) int {
	currIndex := 0
	cnt := 0
	for A[currIndex] != -1 {
		currIndex = A[currIndex]
		cnt++
	}
	return cnt + 1
}
