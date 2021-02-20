package student

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			z01.PrintRune('9')
			n = 223372036854775808
		} else {
			n *= -1
		}
	}
	if n >= 0 {
		if n > 9 {
			PrintNbr(n / 10)
		}
		r := '0'
		for i := 0; i < n%10; i++ {
			r++
		}
		z01.PrintRune(r)
	}
}