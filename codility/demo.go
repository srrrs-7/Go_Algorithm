package codility

import "sort"

func Demo() {
	arr := []int{3, 4, 6, 4, 1}
	sequenceArray(arr)
}

func sequenceArray(A []int) int {
	sort.Ints(A)
	if A[len(A)-1] < 0 {
		return 1
	}

	uniq := distinct(A)
	for i, v := range uniq {
		if v != i+1 {
			return i + 1
		}
	}
	return uniq[len(uniq)-1] + 1
}

func distinct(arr []int) (unique []int) {
	list := make(map[int]bool, len(arr)-1)

	for _, v := range arr {
		if !list[v] {
			list[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}
