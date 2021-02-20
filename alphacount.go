package student

func AlphaCount(str string) int {

	A := []rune(str)
	count := 0

	for index := range A {

		if (A[index] > 96 && 123 > A[index]) || (A[index] > 64 && 91 > A[index]) {

			count++

		}
	}

	return count
}