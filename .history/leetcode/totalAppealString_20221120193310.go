package leetcode

import (
	"fmt"
	"strings"
)

func TotalAppealString() {
	str := "code"

	res := totalAppealString(str)
	fmt.Println("distinct arr", res)
}

func totalAppealString(str string) int {
	res := 0
	arrStr := strings.Split(str, "")

	l := len(str)
	for i := 0; i <= l; i++ {
		for ii := 0; ii <= l; ii++ {
			if i >= ii { continue }
			uniqueArr := distinct(arrStr[i:ii])
			res += len(uniqueArr)
		}
	}

	return res
}


func distinct[T string | int](arr []T) (unique []T) {
	list := make(map[T]bool, len(arr)-1)

	for _, v := range arr {
		if !list[v] {
			list[v] = true
			unique = append(unique, v)
		}
	}
	return unique

}