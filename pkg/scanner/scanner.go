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
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
	"strings"
)

type Scanner interface {
	Next() token.Token
	Peek() string
	Advance()
}

// Scanner is the type that contains functions regarding tokenization of the input stream.
type scanner struct {
	Stream   string
	Position int
	Current  string
}

// Next returns the next token from the input stream.
func (s *scanner) Next() token.Token {
	for s.Current != "" {
		if s.Current == " " {
			s.Advance()
			continue
		}
		if s.Current >= "a" && s.Current <= "z" || s.Current >= "A" && s.Current <= "Z" {
			var sb strings.Builder
			sb.WriteString(s.Current)

			for s.Peek() >= "a" && s.Peek() <= "z" || s.Peek() >= "A" && s.Peek() <= "Z" {
				if strings.ToLower(sb.String()) == "begin" || strings.ToLower(sb.String()) == "end" {
					s.Advance()
					break
				}
				sb.WriteString(s.Peek())
				s.Advance()
			}
			result := strings.ToLower(sb.String())
			var newToken token.Token
			switch result {
			case "begin":
				newToken = token.Token{
					Type: token.Begin,
				}
				break
			case "end":
				newToken = token.Token{
					Type: token.End,
				}
				break
			default:
				newToken = token.Token{
					Type: token.Identifier,
				}
				break
			}

			s.Advance()

			newToken.Lexeme = result
			return newToken
		}

		if s.Current == ":" && s.Peek() == "=" {
			s.Advance()
			s.Advance()
			return token.Token{
				Type:   token.Assign,
				Lexeme: ":=",
			}
		}

		if s.Current == "." {
			s.Advance()
			return token.Token{
				Type:   token.Dot,
				Lexeme: ".",
			}
		}

		if s.Current == ";" {
			s.Advance()
			return token.Token{
				Type:   token.Semi,
				Lexeme: ";",
			}
		}

		if s.Current == "+" {
			s.Advance()
			return token.Token{
				Type:   token.Add,
				Lexeme: "+",
			}
		}

		if s.Current == "-" {
			s.Advance()
			return token.Token{
				Type:   token.Sub,
				Lexeme: "-",
			}
		}

		if s.Current == "*" {
			s.Advance()
			return token.Token{
				Type:   token.Mul,
				Lexeme: "*",
			}
		}

		if s.Current == "/" {
			s.Advance()
			return token.Token{
				Type:   token.Div,
				Lexeme: "/",
			}
		}

		if s.Current == "(" {
			s.Advance()
			return token.Token{
				Type:   token.Lparen,
				Lexeme: "(",
			}
		}

		if s.Current == ")" {
			s.Advance()
			return token.Token{
				Type:   token.Rparen,
				Lexeme: ")",
			}
		}

		if s.Current >= "0" && s.Current <= "9" {
			var sb strings.Builder
			sb.WriteString(s.Current)
			for s.Peek() >= "0" && s.Peek() <= "9" {
				sb.WriteString(s.Peek())
				s.Advance()
			}
			s.Advance()
			return token.Token{
				Type:   token.Int,
				Lexeme: sb.String(),
			}
		}

	}
	return token.Token{
		Type:   token.EOF,
		Lexeme: "",
	}
}

// Peek retrieves the next lexeme from the input stream without advancing the current position
// of the input stream.
func (s *scanner) Peek() string {
	if s.Position+1 >= len(s.Stream) {
		return ""
	}
	return string(s.Stream[s.Position+1])
}

// Advance changes the current position and assigns the new position to s.Current.
func (s *scanner) Advance() {
	if s.Position+1 >= len(s.Stream) {
		s.Current = ""
	} else {
		s.Position++
		s.Current = string(s.Stream[s.Position])
	}
}

// New creates the struct Scanner.
func New(stream string) (*scanner, error) {
	scanner := &scanner{}
	scanner.Stream = stream
	scanner.Current = string(stream[0])
	return scanner, nil
}
