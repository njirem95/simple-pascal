package ast

import "interpreter/pkg/scanner/token"

type UnaryOp struct {
	Operator   token.Token
	Expression *Num
}
