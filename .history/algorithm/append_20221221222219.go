package algorithm

import "fmt"

func AddArray() {
	addArray()
}

func addArray() {
	arr := make([]int, 0)

	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Println("add array: ", arr)

	arr = append(arr, arr...)
	fmt.Println("add array: ", arr)

	arr = append(arr[:2], arr[4:]...)
	fmt.Println("add array: ", arr)
}
