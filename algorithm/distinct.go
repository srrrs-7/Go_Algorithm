package algorithm

import (
	"fmt"
	"strings"
)

func Distinct() {
	v := "mydream"
	arr := strings.Split(v, "")
	distinct(arr)

	st := distinctString(v)
	fmt.Println("distinct string", st)
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

func distinctString(s string) string {
	var sb strings.Builder
	dict := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := dict[s[i]]; !ok {
			dict[s[i]] = i
			sb.WriteString(string(s[i]))
		}
	}
	return sb.String()
}
