package leetcode

import "fmt"

func minimumCostAve() {
	nums := []int{1, 3, 5, 2}
	costs := []int{2, 3, 1, 4}

	res := MinimumCostAve(nums, costs)
	fmt.Println("minimum cost: ", res)
}

func MinimumCostAve(nums []int, costs []int) int {
	sum := 0
	for i := range nums {
		sum += i
	}
	ave := sum / len(nums)
	min := 0
}
