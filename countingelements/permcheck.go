package countingelements

func PremCheck(A []int) int {

	n := len(A)
	seen := make([]bool, n)
	for i := 0; i < n; i++ {
		if !(0 < A[i] && A[i] <= n) {
			return 0
		}
		if seen[A[i]-1] {
			return 0
		}
		seen[A[i]-1] = true
	}
	return 1

}
