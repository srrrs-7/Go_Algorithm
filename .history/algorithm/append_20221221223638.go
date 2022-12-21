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
	arr = append(arr, 6, 7, 8, 9)
	fmt.Println("add array: ", arr)

	arr = []int{1, 2, 3, 4, 5}
	// cut elements
	arr = append(arr[:3], arr[6:]...)
	fmt.Println("add array: ", arr)

	arr = []int{1, 2, 3, 4, 5}
	// delete one element
	arr = append(arr[:2], arr[3:]...)
	fmt.Println("add array: ", arr)

	arr = []int{1, 2, 3, 4, 5}
	//
	arr[0] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	fmt.Println("add array: ", arr)
}
