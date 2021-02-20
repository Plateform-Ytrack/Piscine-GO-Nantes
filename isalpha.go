package student

func IsAlpha(str string) bool {

	a := false 
	for _, i := range str {
		if i > 47 && i < 58 || i > 64 && i < 91 || i > 96 && i < 123 {
			a = true
		} else {
		a = false
		break
		}
	}
	return a
}