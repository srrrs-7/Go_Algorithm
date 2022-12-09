package leetcode

import (
	"fmt"
	"sort"
)

func SubSequence() {
	nums := []int{7, -9, 15, -2}
	goal := -5

	res := subSequence(nums, goal)
	fmt.Println("sub sequence sum: ", res)
}

func subSequence(nums []int, goal int) int {
	arr := getAllSum(nums)
	sort.Ints(arr)

	res := getMinVal(arr, goal)

	return res
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

func getMinVal(arr []int, goal int) int {
	for i, n := range arr {
		if arr[i] <= goal {

		}
	}
}
