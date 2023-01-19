package codility

import (
	"fmt"
	"math"
	"sort"
)

func MaxCounter() {
	arr := []int{3, 4, 4, 6, 1, 4, 4}
	n := 5
	a := maxCounter(n, arr)
	fmt.Println(a)
}

func maxCounter2(N int, A []int) []int {
	arr := make([]int, N)

	for _, v := range A {
		if v > N {
			m := max(arr)
			for i := range arr {
				arr[i] = m
			}
		} else {
			arr[v-1] += 1
		}
	}
	return arr
}

func maxCounter(N int, A []int) []int {
	max := 0
	lastMax := 0
	counters := make([]int, N)

	for k := 0; k < len(A); k++ {
		if A[k] >= 1 && A[k] <= N {
			counters[A[k]-1] = int(math.Max(float64(counters[A[k]-1]), float64(lastMax)))
			counters[A[k]-1] += 1
			max = int(math.Max(float64(max), float64(counters[A[k]-1])))
		} else if A[k] == N+1 {
			lastMax = max
		}
	}

	fmt.Println(counters)

	for i := 0; i < len(counters); i++ {
		counters[i] = int(math.Max(float64(counters[i]), float64(lastMax)))
	}

	fmt.Println(counters)

	return counters
}

func max(arr []int) int {
	a := arr
	sort.Ints(a)
	return a[len(a)-1]
}

func min(arr []int) int {
	a := arr
	sort.Ints(a)
	return a[0]
}
