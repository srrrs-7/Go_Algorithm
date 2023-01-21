package codility

import (
	"fmt"
)

func Dominator() {
	arr := []int{3, 4, 3, 2, 3, -1, 3, 3}
	n := dominator(arr)
	fmt.Println("dominator", n)
}

func dominator(A []int) int {
	persons := make(map[int]int)
	maxCount := len(A) / 2
	dominator := -1
	for i, a := range A {
		persons[a] = persons[a] + 1
		if persons[a] > maxCount {
			maxCount = persons[a]
			dominator = i
		}

	}
	return dominator
}
