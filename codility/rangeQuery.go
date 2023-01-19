package codility

import (
	"fmt"
	"strings"
)

func RangeQuery() {
	S := "CAGCCTA"
	P := []int{2, 5, 0}
	Q := []int{4, 5, 6}

	n := rangeQuery(S, P, Q)
	fmt.Println(n)
}

func rangeQuery(S string, P []int, Q []int) []int {
	arr := make([]int, len(P))

	for i := range P {
		str := S[P[i] : Q[i]+1]

		switch {
		case strings.ContainsRune(str, 'A'):
			arr[i] = 1
		case strings.ContainsRune(str, 'C'):
			arr[i] = 2
		case strings.ContainsRune(str, 'G'):
			arr[i] = 3
		default:
			arr[i] = 4
		}
	}
	return arr
}
