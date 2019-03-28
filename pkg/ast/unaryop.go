package ast

import "github.com/njirem95/simple-pascal/pkg/scanner/token"

type UnaryOp struct {
	Operator   token.Token
	Expression interface{}
}
