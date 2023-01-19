package codility

import (
	"fmt"
)

func FlogRiver() {
	arr := []int{1, 3, 1, 4, 2, 3, 5, 4}
	x := 5

	n := flogRiver(x, arr)
	fmt.Println(n)
}

func flogRiver(X int, A []int) int {
	list := make(map[int]bool, X)

	for i, v := range A {
		if !list[v] {
			list[v] = true
		}
		if len(list) == X {
			return i
		}
	}
	return -1
}
