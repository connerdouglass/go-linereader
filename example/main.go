package main

import (
	"fmt"
	"os"

	"github.com/connerdouglass/go-linereader"
)

func main() {

	file, err := os.Open("linereader.go")
	if err != nil {
		panic(err)
	}

	lr := linereader.New(file)

	for {
		line, err := lr.Line()
		if err != nil {
			panic(err)
		}
		if line == nil {
			break
		}
		fmt.Println("LINE: ", string(line))
	}

}
