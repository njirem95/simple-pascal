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

	result := visitor.Visit(expression.Expression).(int)
	if expression.Operator.Type == token.Add {
		return +result, nil
	} else if expression.Operator.Type == token.Sub {
		return -result, nil
	}

	return 0, errors.New("unable to visit node")
}
