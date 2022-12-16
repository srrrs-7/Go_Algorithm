package utils

import "fmt"

func Contains() {
	arr := []int{1, 2, 3, 4, 5}
	v := 3

	res := contains[int32](arr, v)
	fmt.Println("contains: ", res)
}

func contains[T string | int](arr []T, v T) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}
	return false
}
