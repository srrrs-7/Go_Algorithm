package utils

import (
	"fmt"
	"strconv"
)

func BaseNumberConvert() {
	num := 100
	base := 2

	res := baseNumberConvert(num, base)
	fmt.Println("convert base number", res)
}

func baseNumberConvert(num, base int) []int {
	remainder := 0
	quotient := 0
	results := []int{}

	for {
		remainder = num % base
		quotient = num / base
		num = quotient
		results = append(results, remainder)
		if quotient == 1 {
			results = append(results, quotient)
			break
		}
	}

	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}

	st := ""
	for _, i := range results {
		s := strconv.Itoa(i)
		st += s
	}

	res, err := strconv.Atoi(st)
	if err != nil {}

	return
}