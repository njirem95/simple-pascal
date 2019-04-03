package ast

import "github.com/njirem95/simple-pascal/pkg/scanner/token"

type Variable struct {
	Name  string
	Token token.Token
}
