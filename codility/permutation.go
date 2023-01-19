package codility

import (
	"fmt"
	"sort"
)

func Permutation() {
	arr := []int{1, 4, 5, 6, 2}
	n := permutation(arr)
	fmt.Println(n)
}

func permutation(A []int) int {
	if len(A) == 0 {
		return 1
	}
	sort.Ints(A)
	for i := 0; i < len(A); i++ {
		if i+1 != A[i] {
			return i + 1
		}
	}
	return A[len(A)-1] + 1
}
