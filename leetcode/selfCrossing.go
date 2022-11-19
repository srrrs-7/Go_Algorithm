package leetcode

import "fmt"

func SelfCrossing() {
	arr := [4]int{1,1,1,1}
	res := selfCrossing(arr)

	fmt.Println(res)
}

func selfCrossing(arr [4]int) bool {
	north := arr[0]
	west := arr[1]
	south := arr[2]
	east := arr[3]

	if north - south < 0 {
		return false
	} else if west - east > 0 {
		return false
	}

	return true
}