package algorithm

import "strings"

func Distinct() {
	v := "mydream"
	arr := strings.Split(v, "")

	distinct(arr)
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
