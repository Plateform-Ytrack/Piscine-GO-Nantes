package student

func IsNumeric(str string) bool {
	runes := []rune(str)
	for _, v := range runes {
		if !(v >= 48 && v <= 57) {
			return false
		}
	}
	return true
}