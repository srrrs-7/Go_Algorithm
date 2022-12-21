package algorithm

import "fmt"

func AddArray() {
	addArray()
}

func addArray() {
	arr := make([]int, 0)

	// add elements
	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Println("add array: ", arr)

	// add array
	arr = append(arr, arr...)
	fmt.Println("add array: ", arr)

	// cut
	arr = append(arr[:3], arr[6:]...)
	fmt.Println("add array: ", arr)

	// delete
	arr = append(arr[:2], arr[2+1:]...)
	fmt.Println("add array: ", arr)
}
