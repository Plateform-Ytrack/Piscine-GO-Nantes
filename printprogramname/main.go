package main


import "github.com/01-edu/z01"
import "os"

func PrintStr(s string) {

	s_rune := []rune(s)
	for i, u := range s_rune {
		if u != '.' && u != '/'{
			z01.PrintRune(s_rune[i])
		}
	}
}

func main() {
	ac := os.Args
	var av string = ac[0]
	PrintStr(av)
	z01.PrintRune('\n')
}