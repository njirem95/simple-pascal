package ast

import "interpreter/pkg/scanner/token"

type BinOp struct {
	Left     Num
	Operator token.Token
	Right    Num
}
