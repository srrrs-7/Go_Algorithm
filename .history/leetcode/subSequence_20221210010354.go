package leetcode

import (
	"fmt"
	"sort"
)

func SubSequence() {
	nums := []int{5, -7, 3, 5}
	goal := 6

	res := subSequence(nums, goal)
	fmt.Println("sub sequence sum: ", res)
}

func subSequence(nums []int, goal int) int {
	arr := getAllSum(nums)
	sort.Ints(arr)

	return 111111
}

func getAllSum(nums []int) []int {
	n := len(nums)

	var res []int

	var iter func(idx, sumSoFar int)
	iter = func(idx, sumSoFar int) {
		if idx == n {
			res = append(res, sumSoFar)
			return
		}

		iter(idx+1, sumSoFar)
		iter(idx+1, sumSoFar+nums[idx])
	}

	iter(0, 0)

	return res
}

func binarySearch(arr []int, findNum int) (result bool) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] <= arr[j]
	})

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if findNum > arr[mid] {
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
