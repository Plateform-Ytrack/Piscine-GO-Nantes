package student

func intfor(r rune) int {
	temp2 := -1
	for temp := '0'; temp <= r; temp++ {
		temp2++
	}
	return temp2
}

func BasicAtoi(s string) int {
	arrayStr := []rune(s)
	n := 0
	for range arrayStr {
		n++
	}
	ans := 0
	for i := 0; i < n; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {
			return ans
		} else {
			ans *= 10
			ans += intfor(arrayStr[i])
		}
	}
	return ans
}