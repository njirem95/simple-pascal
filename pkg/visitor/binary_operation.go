package visitor

import (
	"errors"
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
)

type BinOpVisitor struct {
}

func (b *BinOpVisitor) Visit(expression *ast.BinOp) (int, error) {
	visitor := New()

	node := visitor.Visit(expression.Left)
	left, ok := node.(int)
	if !ok {
		return 0, errors.New("expected left to be *ast.Num")
	}

	node = visitor.Visit(expression.Right)

	right, ok := node.(int)
	if !ok {
		return 0, errors.New("expected right to be *ast.Num")
	}

	switch expression.Operator.Type {
	case token.Add:
		return left + right, nil
	case token.Sub:
		return left - right, nil
	case token.Mul:
		return left * right, nil
	case token.Div:
		return left / right, nil
	}

	return 0, errors.New("unable to perform binary operation for unknown reasons")
}
