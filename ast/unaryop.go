package ast

import "interpreter/scanner/token"

type UnaryOp struct {
	Operator   token.Token
	Expression *Num
}
