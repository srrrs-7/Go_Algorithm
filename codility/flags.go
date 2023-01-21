package codility

import (
	"fmt"
)

func Flags() {
	arr := []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2}
	n := flags(arr)
	fmt.Println(n)
}

func flags(A []int) int {
	peak := make([]int, len(A))
	peakCnt := 0
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peak[i] = 1
			peakCnt++
		}
	}
	if peakCnt <= 2 {
		return peakCnt
	}
	next := make([]int, len(A))
	for i := len(A) - 2; i >= 0; i-- {
		if peak[i] == 1 {
			next[i] = i
		} else {
			next[i] = next[i+1]
		}
	}

	flag := 0
	i := 1
	for (i-1)*i <= len(A) {
		curFlag := 0
		curPos := 0
		for curFlag < i && curPos < len(A) {
			curPos = next[curPos]
			if curPos == 0 {
				break
			}
			curFlag++
			curPos = curPos + i
		}

		if flag < curFlag {
			flag = curFlag
		}
		i++
	}
	return flag
}
