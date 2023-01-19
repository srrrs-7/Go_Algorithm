package codility

import "fmt"

func MinAveSlice() {
	arr := []int{4, 2, 2, 5, 1, 5, 8}
	n := minAveSlice(arr)
	fmt.Println(n)
}

func minAveSlice(A []int) int {
	length := len(A)
	minAve := float64((A[0] + A[1]) / 2.0)
	pos := 0
	var temp1, temp2 float64

	if length == 2 {
		return 0
	}
	for i := 2; i < length; i++ {
		temp1 = float64(A[i-1]+A[i]) / 2.0
		temp2 = float64(A[i-2]+A[i-1]+A[i]) / 3.0
		if temp1 < minAve {
			pos = i - 1
			minAve = temp1
		}
		if temp2 < minAve {
			pos = i - 2
			minAve = temp2
		}
	}
	return pos
}
