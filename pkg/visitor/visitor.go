// Package visitor is responsible for interpreting the abstract syntax tree
package visitor

import (
	"github.com/njirem95/simple-pascal/pkg/ast"
	"log"
)

type Visitor struct {
}

func (v *Visitor) Visit(expression ast.Expr) ast.Expr {
	switch expr := expression.(type) {
	case *ast.BinOp:
		node := BinOpVisitor{}
		visit, err := node.Visit(expr)
		if err != nil {
			log.Fatal(err)
		}
		return visit
	case *ast.Num:
		node := NumVisitor{}
		visit, err := node.Visit(expr)
		if err != nil {
			log.Fatal(err)
		}
		return visit
	case *ast.UnaryOp:
		node := UnaryVisitor{}
		visit, err := node.Visit(expr)
		if err != nil {
			log.Fatal(err)
		}
		return visit
	}

	return "nope"
}
