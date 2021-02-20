package student

func Fibonacci(Number int) int {

	if Number < 0 {

		return -1

	}

	if Number < 0 || Number > 500 {

		return 0
	}

	if Number == 0 {
		return 0
	} else if Number == 1 {

		return 1
	} else {

		return Fibonacci(Number-1) + Fibonacci(Number-2)
	}

}