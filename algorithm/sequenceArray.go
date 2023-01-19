package algorithm

import "fmt"

func SequenceArray() {
	min, max := 10, 20

	arr := sequenceArray(min, max)
	fmt.Println("sequence array", arr)
}

func sequenceArray(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
