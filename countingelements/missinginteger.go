package countingelements

func MissingInteger(A []int) int {
	n := len(A)
	seen := make([]bool, n)

	for i := 0; i < n; i++ {
		if A[i] > 0 && A[i] <= n {
			seen[A[i]-1] = true
		}
	}
	for i := 0; i < len(seen); i++ {
		if !seen[i] {
			return i + 1
		}
	}
	return n + 1
}
