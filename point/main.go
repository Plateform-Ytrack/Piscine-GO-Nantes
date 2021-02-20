package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	printString("x = ", points.x, ", y = ", points.y)
}

func printString(s1 string, x int, s2 string, y int) {
	for _, c := range s1 {
		z01.PrintRune(c)
	}

	if x > 9 {
		z01.PrintRune(getRune(x / 10))
		z01.PrintRune(getRune(x % 10))
	}

	for _, c := range s2 {
		z01.PrintRune(c)
	}

	if y > 9 {
		z01.PrintRune(getRune(y / 10))
		z01.PrintRune(getRune(y % 10))
	}

	z01.PrintRune(10)

}

func getRune(i int) rune {
	num := 1
	r := '1'
	for num < i {
		num++
		r++
	}
	return r
}