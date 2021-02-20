package student

func IsUpper(s string) bool {

	a := false
	for _, i:= range s {
	if i >= 65 && i <= 90 {
		a = true
		} else {
		a = false
		break
		}
	}
	return a
}
