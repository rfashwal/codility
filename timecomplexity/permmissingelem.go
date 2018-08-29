package timecomplexity

func PermMissingElem(A []int) int {

	n := len(A)
	total := (n + 1) * (n + 2) / 2

	for i := 0; i < n; i++ {
		total -= A[i]
	}
	return total
}
