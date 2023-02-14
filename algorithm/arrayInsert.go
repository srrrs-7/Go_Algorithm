package algorithm

import "fmt"

func Insert() {
	arr := []int{1, 3, 5, 7}
	num := 2
	res := insert(arr, num)
	fmt.Println("insert function", res)

	res2 := insert2(arr, num)
	fmt.Println("insert function2", res2)
}

func insert(arr []int, num int) []int {
	for i, v := range arr {
		if v > num {
			arr = append(arr, 0)
			copy(arr[i+1:], arr[i:])
			arr[i] = num
			break
		}
	}
	return arr
}

func insert2(arr []int, num int) []int {
	for i, v := range arr {
		if v > num {
			arr = append(arr[:i], append([]int{num}, arr[i:]...)...)
			break
		}
	}
	return arr
}
