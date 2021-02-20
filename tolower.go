package student

func ToLower(s string) string {
	r := []rune(s)
	for index := range r {
		if r[index] >= 'A' && r[index] <= 'Z' {
			r[index] += 32
		}
	}
	return string(r)
}