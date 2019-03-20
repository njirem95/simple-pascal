package ast

import "interpreter/scanner/token"

type BinOp struct {
	Left     *Num
	Operator token.Token
	Right    *Num
}
