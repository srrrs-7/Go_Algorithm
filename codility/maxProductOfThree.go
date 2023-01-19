package codility

import (
	"fmt"
	"sort"
)

func MaxProductOfThree() {
	arr := []int{-3, 1, 2, -2, 5, 6}
	n := maxProductOfThree(arr)
	fmt.Println(n)
}

func maxProductOfThree(A []int) int {
	sort.Ints(A)

	max := A[0] * A[1] * A[len(A)-1]
	max2 := A[len(A)-3] * A[len(A)-2] * A[len(A)-1]

	if max < max2 {
		max = max2
	}

	return max
}
