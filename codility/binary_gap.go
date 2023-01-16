package codility

import (
	"fmt"
	"strings"
)

func BinaryGap() {
	arr := []int{1, 2, 147, 487, 647}
	for _, v := range arr {
		n := binaryGap(561892)
		fmt.Println(n, v)
	}

}

func binaryGap(N int) int {
	s := fmt.Sprintf("%b", N)

	arr := strings.Split(s, "")

	cnt, state := 0, 0
	for i := range arr {
		if arr[i] == "0" {
			state++
		} else if arr[i] == "1" {
			if cnt < state {
				cnt = state
			}
			state = 0
		}
	}
	return cnt
}
