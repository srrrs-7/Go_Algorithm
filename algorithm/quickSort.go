package algorithm

import "fmt"

func QuickSort() {
	arr := []int{2, 5, 4, 2, 6, 1}
	ret := quickSort(arr)
	fmt.Println(ret)
}

func quickSort(arr []int) (ret []int) {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]

	left := []int{}
	right := []int{}
	for _, v := range arr[1:] {
		if v > pivot {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	ret = append(left, pivot)
	ret = append(ret, right...)

	return
}
