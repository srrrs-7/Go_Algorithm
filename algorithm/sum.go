package algorithm

import "fmt"

func Sum() {
	arr := []int{1, 2, 3, 4, 5}
	n := sum(arr)
	fmt.Println(n)
}

func sum(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}
