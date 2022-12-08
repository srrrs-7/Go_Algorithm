package leetcode

import "fmt"

func MinimumCostAve() {
	nums := []int{1, 3, 5, 2}
	costs := []int{2, 3, 1, 4}

	res := minimumCostAve(nums, costs)
	fmt.Println("minimum cost: ", res)
}

func minimumCostAve(nums []int, costs []int) int {
	sum := 0
	for i := range nums {
		sum += i
	}
	ave := sum / len(nums)
	minSost := 0

	for i, n := range nums {

	}
}
