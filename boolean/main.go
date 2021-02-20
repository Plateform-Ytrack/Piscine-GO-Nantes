package main

import (
	"os"

	"github.com/01-edu/z01"
)

var args []string = os.Args

func getArgsLen() int {
	counter := 0
	for range args {
		counter++
	}
	return counter
}

var lengthOfArg int = getArgsLen()

type boolean int

var yes boolean = 1
var no boolean = 0

var EvenMsg string = "I have an even number of arguments"

var OddMsg string = "I have an odd number of arguments"

func even(nbr int) int {
	return nbr % 2
}

func printStr(str string) {
	arrayStr := []rune(str)
	counter := 0
	for range arrayStr {
		counter++
	}
	for i := 0; i < counter; i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	}
	return no
}

func main() {
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}