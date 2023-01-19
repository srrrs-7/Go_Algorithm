package codility

import (
	"fmt"
	"sort"
)

func MissingInteger() {
	arr := []int{-2, -4}
	n := missingInteger(arr)
	fmt.Println(n)
}

func missingInteger(A []int) int {
	sort.Ints(A)
	arr := removeMinus(A)
	if len(arr) == 0 {
		return 1
	}
	arr = distinct(arr)

	for i := range arr {
		if arr[i] != i+1 {
			return i + 1
		}
	}
	return arr[len(arr)-1] + 1
}

func distinct(arr []int) []int {
	list := make(map[int]bool, len(arr)-1)
	var uniq []int

	for _, v := range arr {
		if !list[v] {
			list[v] = true
			uniq = append(uniq, v)
		}
	}
	return uniq
}

func removeMinus(arr []int) []int {
	var res []int

	for i := range arr {
		if arr[i] > 0 {
			res = append(res, arr[i])
		}
	}
	return res
}
