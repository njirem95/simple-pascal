package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/njirem95/simple-pascal/pkg/parser"
	"github.com/njirem95/simple-pascal/pkg/scanner"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("unable to read file", err)
	}

	lexer, err := scanner.New(string(file))
	if err != nil {
		log.Fatal("lexer error:", err)
	}

	parser := parser.New(lexer)
	statements, err := parser.Program()
	if err != nil {
		log.Fatal("parse error:", err)
	}

	// TODO interpret statements
	spew.Dump(statements)
}
