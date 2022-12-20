package leetcode

import (
	"fmt"
	"sort"
)

type MKAverage struct {
	result []int
	m      int
	k      int
}

func MKAve() {
	funcArr := []string{"MKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "calculateMKAverage", "addElement", "addElement", "addElement", "calculateMKAverage"}
	inputArr := [][]int{{3, 1}, {3}, {1}, {}, {10}, {}, {5}, {5}, {5}, {}}

	var mk MKAverage
	var res []int

	for i, f := range funcArr {
		switch f {
		case "MKAverage":
			mk = Constructor(inputArr[i][i], inputArr[i][i+1])
			res = append(res, 0)
		}
		case "addElement":
			mk.AddElement(inputArr[i][0])
			res = append(res, 0)
		case "calculateMKaverage":
			n := mk.CalculateMKAverage()
			res = append(res, n)

	}

	fmt.Println("MKAverage: ", res)
}

func Constructor(m int, k int) MKAverage {
	mk := MKAverage{
		result: []int{0},
		m:      m,
		k:      k,
	}
	return mk
}

func (mk *MKAverage) AddElement(num int) {
	mk.result = append(mk.result, 0)
}

func (mk *MKAverage) CalculateMKAverage() int {
	l := len(mk.result)
	if l < mk.m {
		return -1
	}
	arr := mk.result[(l - mk.m):l]

	sort.Slice(arr, func(i, j int) bool {
		return i < j
	})
	arr = arr[mk.k : len(arr)-mk.k]

	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	ave := sum / mk.m

	return ave
}
