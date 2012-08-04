package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	flag.Parse()
	
	args := flag.Args()
	
	if len(args) == 0 {
		fmt.Printf("You must pass an expression.\n")
		os.Exit(1)
	}
	
	calc := &Calculator{}
	calc.Buffer = args[0]
	
	calc.Init()
	err := calc.Parse()
	
	if err != nil {
		fmt.Printf("Oops, Error! cause: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("No errors found. Syntax tree: \n")
	calc.PrintSyntaxTree()
}
