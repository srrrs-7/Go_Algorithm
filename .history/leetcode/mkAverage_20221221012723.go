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

	for i, f := range funcArr {
		switch f {
		case "addElement":
			mk.AddElement(inputArr[i][0])
		case "calculateMKaverage":
			mk.CalculateMKAverage()
		case "MKAverage":
			mk = Constructor(inputArr[i][i], inputArr[i][i+1])
		}
	}

	fmt.Println("MKAverage: ", mk.result)
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

func (mk *MKAverage) CalculateMKAverage() {
	l := len(mk.result)
	if l < mk.m {
		mk.result = append(mk.result, -1)
		return
	}
	fmt.Println(l, mk.m)
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

	mk.result = append(mk.result, ave)
}
