package leetcode

type MKAverage struct {
	result []*int
}

func MKAverage() {
	funcArr := []string{"MKAverage","addElement","addElement","calculateMKAverage","addElement","calculateMKAverage","addElement","addElement","addElement","calculateMKAverage"}
	inputArr := [][]int{{3,1},{3},{1},{},{10},{},{5},{5},{5},{}}
	for i := range funcArr {
		Constructor(funcArr[i], inputArr[i])
	}


}

func Constructor(m int, k int) MKAverage {
	mk := MKAverage{
		result: ,
	}
}

func (mk *MKAverage) AddElement(num int) {

}

func (mk *MKAverage) CalculateMKAverage() int {

}
