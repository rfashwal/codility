package countingelements

func FrogRiverOne(X int, A []int) int {

	n := len(A)

	visitedPositions := make(map[int]bool, n)

	for i := 0; i < n; i++ {
		if _, ok := visitedPositions[A[i]-1]; !ok {
			visitedPositions[A[i]-1] = true
			X--
			if X == 0 {
				return i
			}
		}
	}
	return -1
}
