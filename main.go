package main

import (
	"fmt"
	"flag"
	"os"
	"strconv"
)

type Op interface {
	Int() int
}

type Int int
func (i Int) Int() int {
	return int(i)
}

type Add struct{
	cm *CalcMachine
	ptr int
}

func (a Add) Int() int {
	return a.cm.opcodes[a.ptr+1].Int() + a.cm.opcodes[a.ptr+2].Int()
}

type CalcMachine struct {
	opcodes []Op
	opcount int
}

// Push a new value into the machine
func (cm *CalcMachine) PushVal(i int) {
	cm.opcodes = append(cm.opcodes, Int(i))
	cm.opcount++
}

func (cm *CalcMachine) PushAdd() {
	a := Add{}
	a.ptr = cm.opcount
	a.cm = cm
	cm.opcodes = append(cm.opcodes, a)
	cm.opcount++
}

func (cm *CalcMachine) Int() int {
	return cm.opcodes[0].Int()
}

var (
	machine = &CalcMachine{opcodes: make([]Op, 0), opcount: 0}
)

func to_int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return i
}

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
	fmt.Printf("Result: %v", machine.Int())
}
