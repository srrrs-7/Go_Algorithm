package utils

import "strconv"

func ArrayToInteger() {
	arr := []int{1,2,3,4,5}
}

func arrayToInteger(arr []int) int {
	st := ""
	for _, x := range arr {
		s := strconv.Itoa(i)
		st += s
	}

	res, err := strconv.Atoi(st)
	if err != nil {
		return -1
	}
}