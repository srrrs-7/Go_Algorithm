package algorithm

import "fmt"

func MergeSort() {
	arr := []int{1,54,6,7,34,5,7,12}
	arr = append(arr, arr[3:]...)

	res := mergeSort(arr)
	fmt.Println("merge sort: ", res)
}

func mergeSort(arr []int) []int {
	var res []int
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	l := mergeSort(arr[:mid])
	r := mergeSort(arr[mid:])
	i, j := 0, 0

	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}

	res = append(res, l[i:]...)
	res = append(res, r[j:]...)

	return res
}

