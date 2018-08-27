package timecomplexity

func Solution(A []int) int {
	sumRight := A[0]
	sumLeft := sumArray(A[1:])
	minNumber := Abs(sumRight - sumLeft)
	n := len(A)
	if n == 1 {
		return sumRight
	}
	if n == 2 {
		diff := Abs(sumRight - sumLeft)
		if minNumber > diff {
			minNumber = diff
		}
		return minNumber
	}
	for i := 1; i < n-1; i++ {
		sumRight += A[i]
		sumLeft -= A[i]
		diff := Abs(sumRight - sumLeft)
		if minNumber > diff {
			minNumber = diff
		}
	}

	return minNumber
}

func sumArray(arr []int) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
