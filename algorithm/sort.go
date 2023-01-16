package algorithm

import (
	"fmt"
	"sort"
)

func SortAlgo() {
	var arrInt = []int{2, 4, 5, 6, 1, 3, 0}
	var arrStr = []string{"bsa", "cds", "adas"}

	SortInt(arrInt)
	SortStr(arrStr)
	SortReverse(arrInt)
}

func SortInt(arr []int) {
	sort.Ints(arr)
	fmt.Println(arr)
}

func SortStr(arr []string) {
	sort.Strings(arr)
	fmt.Println(arr)
}

func SortReverse(arr []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
}
