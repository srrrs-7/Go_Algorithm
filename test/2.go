package test

import "sort"

func Solution2(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1; i >= 2; i-- {
		if A[i-2]+A[i-1] > A[i] {
			return A[i-2] + A[i-1] + A[i]
		}
	}
	return -1
}
