package codility

import (
	"fmt"
	"sort"
)

func PermCheck() {
	arr := []int{2, 3, 5, 4, 1}
	b := permCheck(arr)
	fmt.Println(b)
}

func permCheck(A []int) int {
	sort.Ints(A)
	for i := 0; i < len(A); i++ {
		if i+1 != A[i] {
			return 0
		}
	}
	return 1
}
