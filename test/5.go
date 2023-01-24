package test

import (
	"fmt"
	"sort"
)

func Solution5() {
	arr := []int{6, 10, 6, 9, 7, 8}
	amp := solution5(arr)
	fmt.Println("amplitude", amp)
}

func solution5(A []int) int {
	s := 0
	e := s + 1
	longest := 0

	sort.Ints(A)
	for e <= len(A) {
		sub := A[s:e]
		if maxAmp(sub) < 2 {
			e++
			if longest > len(sub) {
				longest = len(sub)
			}
		} else {
			s = e
			e = s + 1
		}
	}
	return longest
}

func maxAmp(sub []int) int {
	if len(sub) < 2 {
		return 0
	}
	return sub[len(sub)-1] - sub[0]
}
