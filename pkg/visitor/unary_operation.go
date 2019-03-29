package visitor

import (
	"errors"
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
)

type UnaryVisitor struct {
}

func (u *UnaryVisitor) Visit(expression *ast.UnaryOp) (int, error) {
	visitor := Visitor{}
	if expression.Operator.Type == token.Add {
		return +visitor.Visit(expression.Expression).(int), nil
	} else if expression.Operator.Type == token.Sub {
		return -visitor.Visit(expression.Expression).(int), nil
	}

	return 0, errors.New("unable to visit node")
}
