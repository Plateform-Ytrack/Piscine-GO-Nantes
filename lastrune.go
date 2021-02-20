package student

func LastRune(s string) rune {

	A := []rune(s)

	Size := 0

	for range A {
		Size++
	}

	for index, word := range A {

		if index == Size-1 {
			return word
		}
	}

	return 0
}