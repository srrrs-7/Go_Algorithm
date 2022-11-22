package utils

func SequenceArray() {
	min := 10
	max := 20

	res := sequenceArray(min, max)
}

func sequenceArray(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}