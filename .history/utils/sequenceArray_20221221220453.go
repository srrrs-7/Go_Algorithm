package utils

import "fmt"

func SequenceArray() {
	min := 10
	max := 20

	res := sequenceArray(min, max)
	fmt.Printf("%d \n", res)

	res = sameValueArray(max)
	fmt.Printf("%d \n", res)
}

func sequenceArray(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func sameValueArray[T int | string](num, v T) []T {
	a := make([]int, num)
	for i := range a {
		a[i] = v
	}
	return a
}
