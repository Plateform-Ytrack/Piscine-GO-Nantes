package student

func intfor2(r rune) int {
	temp2 := -1
	for temp := '0'; temp <= r; temp++ {
		temp2++
	}
	return temp2
}

func BasicAtoi2(s string) int {
	arrayStr := []rune(s)
	n := 0
	for range arrayStr {
		n++
	}
	ans := 0
	for i := 0; i < n; i++ {
		if arrayStr[i] < '0' || arrayStr[i] > '9' {
			return 0
		} else {
			ans *= 10
			ans += intfor2(arrayStr[i])
		}
	}
	return ans
}