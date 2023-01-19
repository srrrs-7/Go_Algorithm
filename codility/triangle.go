package codility

import (
	"fmt"
	"sort"
)

func Triangle() {
	arr := []int{10, 2, 5, 1, 8, 20}
	b := triangle(arr)
	fmt.Println(b)
}

func triangle(A []int) int {
	sort.Ints(A)

	for i := 2; i < len(A); i++ {
		edge1 := A[i-2] + A[i-1]
		edge2 := A[i-1] + A[i]
		edge3 := A[i] + A[i-2]

		if edge1 > A[i] && edge2 > A[i-2] && edge3 > A[i-1] {
			return 1
		}
	}
	return 0
}
