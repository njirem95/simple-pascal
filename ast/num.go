package ast

import "github.com/njirem95/simple-pascal/scanner/token"

type Num struct {
	Token  token.Token
	Lexeme string
}
