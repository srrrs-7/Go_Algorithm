package codility

import "fmt"

func StoneWall() {
	highs := []int{8, 8, 5, 7, 9, 8, 7, 4, 8}
	n := stoneWall(highs)
	fmt.Println(n)
}

func stoneWall(H []int) int {
	stack := []int{}
	block := 0

	for i := range H {
		for len(stack) != 0 && stack[len(stack)-1] > H[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			block++
			stack = append(stack, H[i])

		} else if stack[len(stack)-1] == H[i] {
		} else if stack[len(stack)-1] < H[i] {
			block++
			stack = append(stack, H[i])
		}
	}
	return block
}
