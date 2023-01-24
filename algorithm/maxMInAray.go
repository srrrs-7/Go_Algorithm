package algorithm

import (
	"fmt"
	"sort"
)

func MaxMinArray() {
	arr := []int{2, 4, 1, 7, 4}
	max := maxArray(arr)
	min := minArray(arr)
	fmt.Println("max", max, "min", min)
}

func maxArray(A []int) int {
	sort.Ints(A)
	return A[len(A)-1]
}

func minArray(A []int) int {
	sort.Ints(A)
	return A[0]
}
