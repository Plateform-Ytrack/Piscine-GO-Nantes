package student

func Index(s string, toFind string) int {

	cat := []rune(s)
	dog := []rune(toFind)

	a := 0
	b := 0

	for range cat {
		a++
	}

	for range dog {
		b++
	}
	for i := 0; i <= a-b; i++ {
		if toFind == s[i:(i+b)] {
			return i
		}
	}
	return -1
}