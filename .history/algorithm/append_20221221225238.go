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
	arr = append(arr[:2], arr[4:]...)
	fmt.Println("add array: ", arr)

	arr = []int{1, 2, 3, 4, 5}
	// delete one element
	arr = append(arr[:2], arr[3:]...)
	fmt.Println("add array: ", arr)

	arr = []int{1, 2, 3, 4, 5}
	// insert last element
	arr[2] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	fmt.Println("add array: ", arr)

	arr = []int{1, 2, 3, 4, 5}
	// expand
	arr = append(arr[:2], append(make([]int, 5), arr[2:]...)...)
	fmt.Println("add array: ", arr)

	arr = []int{1, 2, 3, 4, 5}
	// expand
	arr = append(arr, make([]int, 5)...)
	fmt.Println("add array: ", arr)

	arr = []int{1, 2, 3, 4, 5}
	// insert
	arr = append(arr[:2], append([]int{3}, arr[2:]...)...)
	fmt.Println("add array: ", arr)

	arr = []int{1, 2, 3, 4, 5}
	// queue
	x, a := arr[0], arr[1:]
	fmt.Println("add array: ", x, a)

	arr = []int{1, 2, 3, 4, 5}
	// stack
	y, w := arr[:len(arr)-2], arr[len(arr)-1]
	fmt.Println("add array: ", y, w)

	arr = []int{1, 2, 3, 4, 5}
	// push front
	b, c := arr[], arr[]
	fmt.Println("add array: ", b, c)
}
