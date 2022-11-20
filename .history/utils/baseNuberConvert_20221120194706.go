package utils

import (
	"fmt"
	"strconv"
)

func BaseNumberConvert() {
	num := 10
	base := 2

	res := baseNumberConvert(num, base)
	fmt.Println("convert base number", res)
}

func baseNumberConvert(num, base int) int {
	remainder := 0
	quotient := 0
	result := []int{}

	for {
		remainder = num % base
		quotient = num / base
		num = quotient
		result = append(result, remainder)
		if quotient == 1 { break }

	}

	st := fmt.Sprintf("%d", result[0:])
	fmt.Println(st)
	res, err := strconv.Atoi(st)
	if err != nil {
		return -1
	}

	return res
}