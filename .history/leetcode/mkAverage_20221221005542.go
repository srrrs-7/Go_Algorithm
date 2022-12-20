package leetcode

type MKAverage struct {
	result []*int
}

func MKAverage() {
	funcArr := []string{"MKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "calculateMKAverage", "addElement", "addElement", "addElement", "calculateMKAverage"}
	inputArr := [][]int{{3, 1}, {3}, {1}, {}, {10}, {}, {5}, {5}, {5}, {}}

	var mk MKAverage

	for i, f := range funcArr {
		switch f {
		case "MKAverage":
			mk = Constructor(inputArr[i][i], inputArr[i][i+1])
			mk.result = append(mk.result, null)
		case "addElement":
			mk.AddElement(inputArr[i])
		case "calculateMKaverage":
			mk.CalculateMKAverage()
		}
	}

	return mk
}

func Constructor(m int, k int) MKAverage {
	mk := MKAverage{
		result: []*int{},
	}
	return mk
}

func (mk *MKAverage) AddElement(num int) {

}

func (mk *MKAverage) CalculateMKAverage() int {

}
