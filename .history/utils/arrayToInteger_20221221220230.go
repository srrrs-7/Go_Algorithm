package utils

import (
	"fmt"
	"strconv"
)

func ArrayToInteger() {
	arr := []int{1, 2, 3, 4, 5}

	res := arrayToInteger(arr)
	fmt.Println("array to integer", res)
}

func arrayToInteger(arr []int) int {
	st := ""
	for _, x := range arr {
		s := strconv.Itoa(x)
		st += s
	}

	res, err := strconv.Atoi(st)
	if err != nil {
		return -1
	}

	return res
}
