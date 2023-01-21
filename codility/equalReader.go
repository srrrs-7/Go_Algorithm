package codility

import (
	"fmt"
)

func EqualReader() {
	arr := []int{4, 4, 2, 5, 3, 4, 4, 4}
	n := equalReader(arr)
	fmt.Println("leader", n)
}

func equalReader(A []int) int {
	a_map := make(map[int]int)
	leader := 0
	leaderCnt := 0
	for _, a := range A {
		a_map[a] += 1
		if a_map[a] > leaderCnt {
			leaderCnt = a_map[a]
			leader = a
		}
	}

	leftLeader := 0
	rightLeader := a_map[leader]
	result := 0
	for i, a := range A {
		if a == leader {
			leftLeader += 1
			rightLeader -= 1
		}
		if (i+1)/2 < leftLeader && (len(A)-1-i)/2 < rightLeader {
			result = result + 1
		}
	}
	return result
}
