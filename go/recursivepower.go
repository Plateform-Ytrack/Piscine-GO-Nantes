package student

func RecursivePower(Number int, Number2 int) int {

	if Number2 == 1 {

		return Number
	}

	if Number2 < 0 {

		return 0

	} else if Number2 < 0 || Number2 > 12 {

		return 0

	} else if Number2 == 0 {

		return 1

	} else {

		return Number * RecursivePower(Number, Number2-1)

	}

}
