package ast

import "interpreter/pkg/scanner/token"

type UnaryOp struct {
	operator   token.Token
	Expression Num
}

func (u *UnaryOp) Token() token.Token {
	return u.operator
}

func (u *UnaryOp) SetToken(token token.Token) {
	u.operator = token
}