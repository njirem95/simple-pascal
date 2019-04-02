package visitor

import (
	"errors"
	"fmt"
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
)

type BinOpVisitor struct {
}

func (b *BinOpVisitor) Visit(expression *ast.BinOp) (int, error) {
	visitor := Visitor{}

	node, err := visitor.Visit(expression.Left)
	if err != nil {
		return 0, err
	}
	left, ok := node.(int)
	if !ok {
		return 0, errors.New("expected left to be an integer")
	}

	node, err = visitor.Visit(expression.Right)
	if err != nil {
		return 0, err
	}
	right, ok := node.(int)
	if !ok {
		return 0, errors.New("expected right to be an integer")
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

	return 0, fmt.Errorf("unknown operator type %s", expression.Operator.Lexeme)
}
