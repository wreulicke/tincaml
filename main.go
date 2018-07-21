package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wreulicke/tincaml/interpreter"
	"github.com/wreulicke/tincaml/parse"
)

func main() {
	f, _ := os.Open("simple")
	n, err := parse.Parse(bufio.NewReader(f))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = interpreter.Evaluate(n)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
