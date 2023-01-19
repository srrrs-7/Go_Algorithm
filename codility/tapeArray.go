package codility

import (
	"fmt"
	"math"
)

func TapeArray() {
	arr := []int{3, 1, 2, 4, 3}
	n := tapeArray(arr)
	fmt.Println(n)
}

func tapeArray(A []int) int {
	var diff int
	var left int
	var right int
	var sum int

	for i := 0; i < len(A); i++ {
		sum += A[i]
	}

	left = A[0]
	right = sum - left
	diff = int(math.Abs(float64(left - right)))
	for i := 2; i < len(A); i++ {
		left += A[i-1]
		right -= A[i-1]
		n := int(math.Abs(float64(left - right)))
		if diff > int(n) {
			diff = int(n)
		}
	}
	return diff
}
