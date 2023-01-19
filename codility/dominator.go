package codility

import (
	"fmt"
	"sort"
)

func Dominator() {
	arr := []int{3, 4, 3, 2, 3, -1, 3, 3}
	n := dominator(arr)
	fmt.Println("dominator", n)
}

func dominator(A []int) int {
	if len(A) == 0 {
		return -1
	}
	list := make(map[int]int, 0)
	for _, v := range A {
		list[v] += 1
	}

	var key, cnt int
	for k, v := range list {
		if key < v {
			key = k
			cnt = v
		}
	}

	if cnt < len(A)/2 {
		return -1
	}

	for k, v := range list {
		if k != key && v == cnt {
			return -1
		}
	}

	return sort.SearchInts(A, key)
}
