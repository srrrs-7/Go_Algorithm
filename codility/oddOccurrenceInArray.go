package codility

import (
	"fmt"
)

func OddOccurrenceInArray() {
	arr := []int{9, 3, 9, 3, 9, 7, 9}
	n := oddOccurrenceInArray(arr)
	fmt.Println(n)
}

func oddOccurrenceInArray(A []int) int {
	var elmt int
	for i := 0; i < len(A); i++ {
		elmt ^= A[i]
	}
	return elmt
}
