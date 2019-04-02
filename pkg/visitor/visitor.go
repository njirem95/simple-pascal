// Package visitor is responsible for interpreting the abstract syntax tree
package visitor

import (
	"errors"
	"github.com/njirem95/simple-pascal/pkg/ast"
)

type Visitor struct {
}

func (v *Visitor) Visit(expression ast.Expr) (ast.Expr, error) {
	switch expr := expression.(type) {
	case *ast.BinOp:
		node := BinOpVisitor{}
		visit, err := node.Visit(expr)
		return visit, err
	case *ast.Num:
		node := NumVisitor{}
		visit, err := node.Visit(expr)
		return visit, err
	case *ast.UnaryOp:
		node := UnaryVisitor{}
		visit, err := node.Visit(expr)
		return visit, err
	}

	return nil, errors.New("visitor not found")
}
