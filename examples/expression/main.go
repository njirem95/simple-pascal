package main

import (
	"fmt"
	"github.com/njirem95/simple-pascal/pkg/parser"
	"github.com/njirem95/simple-pascal/pkg/scanner"
	"github.com/njirem95/simple-pascal/pkg/visitor"
	"log"
)

func main() {
	inputs := [4]string{"2", "2 + 2", "2 + 2 * 4", "(2 + 2) * 4"}
	for _, input := range inputs {
		lexer, err := scanner.New(input)
		if err != nil {
			log.Fatal(err)
		}

		parser := parser.New(lexer)

		expression, err := parser.Expr()
		if err != nil {
			log.Fatal(err)
		}

		visitor := visitor.Visitor{}
		result := visitor.Visit(expression)

		fmt.Println(result)
	}

}
