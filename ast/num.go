package ast

import "interpreter/scanner/token"

type Num struct {
	Token  token.Token
	Lexeme string
}
