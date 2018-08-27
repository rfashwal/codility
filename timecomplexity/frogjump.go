package timecomplexity

func FrogJump(X int, Y int, D int) int {

	K := Y - X
	
	result := K / D

	if K%D != 0 {
		result++
	}
	return result
}
