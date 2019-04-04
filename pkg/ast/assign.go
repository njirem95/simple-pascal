package ast

import "github.com/njirem95/simple-pascal/pkg/scanner/token"

type Assign struct {
	Left     interface{}
	Operator token.Token
	Right    Expr
}
