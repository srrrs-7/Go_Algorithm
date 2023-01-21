package codility

import "fmt"

func MaxSliceSum() {
	arr := []int{3, 2, -6, 4, 0}
	n := maxSliceSum(arr)
	fmt.Println(n)
}

func maxSliceSum(A []int) int {
	maxPair := A[0]
	currMax := A[0]

	for i := 1; i < len(A); i++ {
		maxPair = maxInt(A[i], maxPair+A[i])
		currMax = maxInt(currMax, maxPair)
	}
	return currMax
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
