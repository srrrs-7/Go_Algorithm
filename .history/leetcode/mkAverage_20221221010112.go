package leetcode

import "fmt"

type MKAverage struct {
	result []*int
	m      int
	k      int
}

func MKAve() {
	funcArr := []string{"MKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "calculateMKAverage", "addElement", "addElement", "addElement", "calculateMKAverage"}
	inputArr := [][]int{{3, 1}, {3}, {1}, {}, {10}, {}, {5}, {5}, {5}, {}}

	var mk MKAverage

	for i, f := range funcArr {
		switch f {
		case "MKAverage":
			mk = Constructor(inputArr[i][i], inputArr[i][i+1])
			mk.result = append(mk.result, nil)
		case "addElement":
			mk.AddElement(inputArr[i][0])
		case "calculateMKaverage":
			mk.CalculateMKAverage()
		}
	}

	fmt.Println(mk.result)
}

func Constructor(m int, k int) MKAverage {
	mk := MKAverage{
		result: []*int{},
		m:      m,
		k:      k,
	}
	return mk
}

func (mk *MKAverage) AddElement(num int) {
	mk.result = append(mk.result, &num)
}

func (mk *MKAverage) CalculateMKAverage() {
	l := len(mk.result)

}
