package codility

import "fmt"

func PassingCar() {
	arr := []int{0, 1, 0, 1, 1}
	n := passingCar(arr)
	fmt.Println(n)
}

func passingCar(A []int) int {
	var east int
	var cnt int

	for i := range A {
		if A[i] == 0 {
			east += 1
		} else {
			cnt += east
		}
	}

	if cnt > 1000000000 {
		return -1
	}

	return cnt
}
