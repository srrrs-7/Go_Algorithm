package leetcode

type MKAverage struct {
	result []*int
}


func Constructor(m int, k int) MKAverage {
	funcArr := []string{"MKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "calculateMKAverage", "addElement", "addElement", "addElement", "calculateMKAverage"}
	inputArr := [][]int{{3, 1}, {3}, {1}, {}, {10}, {}, {5}, {5}, {5}, {}}
	for i, f := range funcArr {
		Constructor(funcArr[i], inputArr[i])
	}
	mk := MKAverage{
		result: []*int{},
	}

	switch
}

func (mk *MKAverage) AddElement(num int) {

}

func (mk *MKAverage) CalculateMKAverage() int {

}