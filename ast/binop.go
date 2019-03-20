package ast

import "github.com/njirem95/simple-pascal/scanner/token"

type BinOp struct {
	Left     *Num
	Operator token.Token
	Right    *Num
}
