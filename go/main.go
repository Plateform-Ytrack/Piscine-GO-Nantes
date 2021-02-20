package student

import "github.com/01-edu/z01"

func piscine() {

	for i := 97; i <= 122; i++ {

		z01.PrintRune(rune(i))
	}

	z01.PrintRune('\n')
}
