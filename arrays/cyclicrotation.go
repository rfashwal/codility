<<<<<<< HEAD
package arrays

/*
An array A consisting of N integers is given. Rotation of the array means that
each element is shifted right by one index, and the last element of the array
is moved to the first place. For example, the rotation of array A = [3, 8, 9, 7, 6]
is [6, 3, 8, 9, 7] (elements are shifted right by one index and 6 is moved to the first place).

The goal is to rotate array A K times; that is, each element of A will be shifted
to the right K times.

Write a function:
func Solution(A []int, K int) []int

that, given an array A consisting of N integers and an integer K,
returns the array A rotated K times.

*/

func CyclicRotation(A []int, K int) []int {
	lenA := len(A)
	if lenA == K || K == 0 || lenA == 0 {
		return A
	}
	if K > len(A) {
		K = K % lenA
	}
	rotattedArray := make([]int, lenA)
	for i := 0; i < lenA; i++ {
		newIndex := i + K
		if newIndex >= lenA {
			newIndex = newIndex - lenA
		}
		rotattedArray[newIndex] = A[i]
	}
	return rotattedArray
}
=======
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
>>>>>>> 11cc22cec302a686b6a64e41583deb7ddce7e95c
