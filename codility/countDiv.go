package codility

import (
	"fmt"
	"math"
)

func CountDiv() {
	n := countDiv(11, 345, 17)
	fmt.Println(n)

	n = countDiv2(11, 345, 17)
	fmt.Println(n)
}

func countDiv2(A int, B int, K int) int {
	return int(math.Floor(float64(B)/float64(K))) - int(math.Ceil(float64(A)/float64(K))) + 1
}

func countDiv(A int, B int, K int) int {
	arr := sequence(A, B)

	cnt := 0
	for i := range arr {
		if arr[i]%K == 0 {
			cnt += 1
		}
	}
	return cnt
}

func sequence(min, max int) []int {
	arr := make([]int, max-min+1)
	for i := range arr {
		arr[i] = min + i
	}
	return arr
}
