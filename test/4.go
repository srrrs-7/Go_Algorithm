package test

import (
	"fmt"
	"strconv"
	"strings"
)

func Solution4() {
	num := 12144
	n := solution4(num)
	fmt.Println(n)
}

func solution4(N int) int {
	perm := 1
	arr := strings.Split(fmt.Sprint(N), "")
	for i := 2; i <= len(arr); i++ {
		perm *= i
	}

	fmt.Println(len(arr))

	list := make(map[int]int, 0)
	for _, v := range arr {
		n, _ := strconv.Atoi(v)
		list[n]++
	}

	for _, v := range list {
		if v != 1 {
			perm /= v
		}
	}

	return perm

}
