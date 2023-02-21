package leetcode

import "fmt"

func Subset() {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
}

func subsets(nums []int) [][]int {
	result := [][]int{}
	generateSubsets(nums, []int{}, &result)
	return result
}

func generateSubsets(nums []int, current []int, result *[][]int) {
	// response pointer
	*result = append(*result, current)
	// nums length loop -> for and recursion double loop
	for i := 0; i < len(nums); i++ {
		// generate array loop and recursion
		newSubset := make([]int, len(current))
		copy(newSubset, current)
		newSubset = append(newSubset, nums[i])
		// recursive function -> decrease one in array
		generateSubsets(nums[i+1:], newSubset, result)
	}
}
