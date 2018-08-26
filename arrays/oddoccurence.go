package arrays

func OddOccurence(A []int) int {
	odd := 0

	for i := 0; i < len(A); i++ {
		odd ^= A[i] //xor
	}
	return odd
}
