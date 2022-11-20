package utils

import (
	"fmt"
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
	result := 0
	cnt := 1

	for quotient != 1 {
		remainder = num % base
		quotient = num / base
		num = quotient
		result = remainder * (10^cnt)
		cnt += 1

	}

	return result
}