package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wreulicke/tincaml/interpreter"
	"github.com/wreulicke/tincaml/parse"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("0 argument is expected tincaml file")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()

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
