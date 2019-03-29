// Package visitor is responsible for interpreting the abstract syntax tree
package visitor

import (
	"github.com/njirem95/simple-pascal/pkg/ast"
)

type Visitor interface {
	Visit(expression ast.Expr)
}

type visitor struct {
}

func (v *visitor) Visit(expression ast.Expr) ast.Expr {
	switch expr := expression.(type) {
	case *ast.BinOp:
		node := BinOpVisitor{}
		visit, _ := node.Visit(expr)
		return visit
	case *ast.Num:
		node := NumVisitor{}
		visit, _ := node.Visit(expr)
		return visit
	case *ast.UnaryOp:
		// Use the UnaryOp visitor
		break
	}

	return "nope"
}

func New() *visitor {
	return &visitor{}
}
