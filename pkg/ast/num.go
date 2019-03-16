package ast

import "interpreter/pkg/scanner/token"

type Num struct {
	Token  token.Token
	Lexeme string
}
