package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	length := 0
	for i := range args {
		length = i + 1
	}

	if length == 0 {
		fmt.Println("File name missing")
		return
	}
	if length != 1 {
		fmt.Println("Too many arguments")
		return
	}

	file_name := args[0]
	file, err := os.Open(file_name)

	entered := false
	if err != nil {
		entered = true
		fmt.Println(err.Error())
	}

	b, er := ioutil.ReadAll(file)
	fmt.Print(string(b))
	if er != nil && !entered {
		fmt.Println(err.Error())
	}

}