package codility

func CyclicRotation() {
	A := []int{}
	K := 3
	cyclicRotation(A, K)
}

func cyclicRotation(A []int, K int) []int {
	arr := make([]int, len(A))
	if len(A) == 0 {
		return arr
	}

	cnt, r := 0, (K % len(arr))
	for i := range A {
		if i+r < len(arr) {
			arr[i+r] = A[i]
		} else {
			arr[cnt] = A[i]
			cnt++
		}
	}

	return arr
}
