package leetcode

import "fmt"

func AggregateComponent() {
	num := []int16{6,2,2,2,6}
	edges := [][]int16{{0,1},{1,2},{1,3},{3,4}}

	res := aggregateComponent(num, edges)
	fmt.Println(res)
}

func aggregateComponent(num []int16, edges [][]int16) int16 {
	res := int16(0)

	for _, e := range edges {
		if num[e[0]] == num[e[1]] {
			res += 1
		}
	}
	return res
}