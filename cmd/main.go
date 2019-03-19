package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jmhobbs/calculator"
)

func main() {
	ops, err := calculator.Parse(strings.Join(os.Args[1:], ""))
	if err != nil {
		log.Fatal(err)
	}
	x := 0
	for _, op := range ops {
		switch op.Operator {
		case 0:
			x = op.Operand
			fmt.Print(x)
		case '+':
			x = x + op.Operand
			fmt.Print(" + ", op.Operand)
		case '-':
			x = x - op.Operand
			fmt.Print(" - ", op.Operand)
		case '*':
			x = x * op.Operand
			fmt.Print(" * ", op.Operand)
		case '/':
			x = x / op.Operand
			fmt.Print(" / ", op.Operand)
		}
	}
	fmt.Println(" =", x)
}
