package student

func RuneToUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

func RuneToLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func Capitalize(s string) string {
	r := []rune(s)
	isFirstLetter := true
	lostWord := true
	for index := range s {
		if (r[index] >= 'a' && r[index] <= 'z') || (r[index] >= 'A' && r[index] <= 'Z') || (r[index] >= '0' && r[index] <= '9') {
			if isFirstLetter {
				r[index] = RuneToUpper(r[index])
			} else {
				r[index] = RuneToLower(r[index])
			}
			if isFirstLetter == true && lostWord == true {
				isFirstLetter = false
				lostWord = false
			}
		} else {
			lostWord = true
			isFirstLetter = true
		}
	}
	return string(r)
}