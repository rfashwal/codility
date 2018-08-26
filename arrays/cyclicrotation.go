package arrays

func CyclicRotation(A []int, K int) []int {
	n := len(A)
	rotatedArray := make([]int, n)
	if K > n {
		K = K % n
	}

	for i := 0; i < n; i++ {
		index := i + K
		if index >= n {
			index -= n
		}
		rotatedArray[index] = A[i]
	}
	return rotatedArray
}
