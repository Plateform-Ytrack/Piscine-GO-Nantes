package student

func IterativePower(Number int, Number2 int) int {

	result := 1

	if Number == 0 && Number2 == 0 {

		return 1

	} else if Number2 < 0 || Number2 > 12 {

		return 0

	}

	for i := 0; i < Number2; i++ {

		result = result * Number
	}

	return result

}
