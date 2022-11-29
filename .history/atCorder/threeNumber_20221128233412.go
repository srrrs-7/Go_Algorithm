package atCorder

import "fmt"

func ThreeNumber() {
	arr := []int{1, 4, 9}
	multi := 7

	res := threeNumber(arr, multi)
	fmt.Println("multi number count: ", res)
}

func threeNumber(arr []int, multi int) int {
	array := make([]int, len(arr))
}