package codility

import "fmt"

func Distinct() {
	arr := []int{1, 2, 3, 1, 2, 3}
	n := distinctCount(arr)
	fmt.Println(n)
}

func distinctCount(A []int) int {
	list := make(map[int]bool, len(A))
	uniq := make([]int, 0)

	for _, v := range A {
		if !list[v] {
			list[v] = true
			uniq = append(uniq, v)
		}
	}
	return len(uniq)
}
