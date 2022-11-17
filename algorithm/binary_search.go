package algorithm

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, findNum int) (result bool) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] <= arr[j]
	})

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if  findNum > arr[mid]{
			low = mid + 1
		} else {
			high = mid - 1
		}

		if findNum == arr[mid] {
			return true
		}
	}

	return false
}

func BinarySearch() {
	arr := []int{3,5,7,1,2,32,3,5,9}
	findNum := 5
	res := binarySearch(arr, findNum)
	fmt.Println("binary sort: ", res)
}