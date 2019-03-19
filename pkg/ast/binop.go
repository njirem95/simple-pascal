package ast

import "interpreter/pkg/scanner/token"

type BinOp struct {
	Left     Num
	operator token.Token
	Right    Num
}

func (b *BinOp) Token() token.Token {
	return b.operator
}

func (b *BinOp) SetToken(token token.Token) {
	b.operator = token
}