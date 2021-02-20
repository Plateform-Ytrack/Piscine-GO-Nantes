package student

func TrimAtoi(s string) int {
	runes := []rune(s)
	var digits []int
	minus := false
	plus := false
	counter := 0
	result := 0
	for _, v := range runes {
		if int(v) == 45 {
			if digits != nil {
				continue
			} else if minus == true || plus == true {
				return 0
			} else {
				minus = true
			}
		}
		if int(v) == 43 {
			if digits != nil {
				continue
			} else if plus == true || minus == true {
				return 0
			} else {
				plus = true
			}
		}
		if int(v) >= 48 && int(v) <= 57 {
			digits = append(digits, int(v))
			counter++
		}
	}
	j := 1
	for i := counter - 1; i >= 0; i-- {
		digits[i] = (digits[i] - 48) * j
		j *= 10
		result += digits[i]
	}
	if minus == true {
		result *= -1
	}
	return result
}