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
	"interpreter/pkg/scanner/token"
)

// Scanner is the type that contains functions regarding tokenization of the input stream.
type Scanner struct {
	Stream   string
	Position int
	Current  string
}

// Next returns the next token from the input stream.
func (s *Scanner) Next() token.Token {
	for s.Current != "" {
		if s.Current == " " {
			s.Advance()
			continue
		}

		if s.Current == "+" {
			s.Advance()
			return token.Token{
				Type:   token.Add,
				Lexeme: "+",
			}
		}
		if s.Current > "0" {
			next := token.Token{
				Type:   token.Int,
				Lexeme: s.Current,
			}
			s.Advance()
			return next
		}

	}
	return token.Token{
		Type:   token.EOF,
		Lexeme: "",
	}
}

// Advance changes the current position and assigns the new position to s.Current.
func (s *Scanner) Advance() {
	if s.Position+1 >= len(s.Stream) {
		s.Current = ""
	} else {
		s.Position++
		s.Current = string(s.Stream[s.Position])
	}
}

// New creates the struct Scanner. Returns an error if the stream size is zero.
func New(stream string) (*Scanner, error) {
	if len(stream) == 0 {
		return nil, errors.New("input stream is empty")
	}
	scanner := &Scanner{}
	scanner.Stream = stream
	scanner.Current = string(stream[0])
	return scanner, nil
}
