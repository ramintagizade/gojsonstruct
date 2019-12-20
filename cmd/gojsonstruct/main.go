package main

import (
	"fmt"
	"github.com/ramintagizade/gojsonstruct"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing filename, provide file name!")
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
	}

	gojsonstruct.Compile(string(data))
}
