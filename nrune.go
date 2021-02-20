package student

func NRune(s string, n int) rune {

	A := []rune(s)

	for index, word := range A {

		if index == n-1 {
			return word
		}
	}

	return 0
}