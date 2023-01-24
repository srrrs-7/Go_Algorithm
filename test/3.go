package test

import "fmt"

func Solution3() {
	arr := []int{1, 1, 1, 1, 1, 0}
	n := solution3(arr)
	fmt.Println(n)
}

func solution3(A []int) int {
	n := len(A)
	if n <= 1 {
		return 0
	}
	res := 0
	// pair count
	for i := 0; i < n-1; i++ {
		if A[i] == A[i+1] {
			res++
		}
	}
	// [1,1], [1,1,1]
	if res == n-1 {
		return res - 1
	}
	// reverse count
	r := 0
	for i := 0; i < n; i++ {
		cnt := 0
		// max increment -> 2
		if i > 0 {
			// reverse -> equal
			if A[i-1] != A[i] {
				cnt++
			} else {
				cnt--
			}
		}
		if i < n-1 {
			// reverse -> equal
			if A[i+1] != A[i] {
				cnt++
			} else {
				cnt--
			}
		}
		r = max(r, cnt)
	}
	return res + r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
