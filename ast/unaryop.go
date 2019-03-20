package ast

import "github.com/njirem95/simple-pascal/scanner/token"

type UnaryOp struct {
	Operator   token.Token
	Expression *Num
}
