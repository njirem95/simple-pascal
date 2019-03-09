/*
Package scanner is responsible for generating tokens of the input stream.
This process is also known as lexical analyzing, tokenization, and scanning.

For instance, the input "1 + 2 - 1" is equivalent to the following set of tokens:
	[]Token{
		{
			Type: Token.Int,
			Lexeme: "1",
		},
		{
			Type: Token.Add,
			Lexeme: "+",
		},
		{
			Type: Token.Int
			Lexeme: "2",
		},
		{
			Type: Token.Sub,
			Lexeme: "-",
		},
		{
			Type: Token.Int,
			Lexeme: "1",
		},
		{
			Type: Token.EOF,
			Lexeme: "",
		},
	}

As you can see, each token consists of a TokenType and a lexeme (an element of the input stream).
Every set ends with the EOF (End-of-file) token.
*/
package scanner

import (
	"errors"
)

// Scanner is the type that contains functions regarding tokenization of the input stream.
type Scanner struct {
	Stream string
}

// New creates the struct Scanner. Returns an error if the stream size is zero.
func New(stream string) (*Scanner, error) {
	if len(stream) == 0 {
		return nil, errors.New("input stream is empty")
	}
	scanner := &Scanner{}
	scanner.Stream = stream
	return scanner, nil
}
