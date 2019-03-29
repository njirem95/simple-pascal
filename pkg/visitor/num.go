package visitor

import (
	"github.com/njirem95/simple-pascal/pkg/ast"
	"strconv"
)

type NumVisitor struct {
}

func (n *NumVisitor) Visit(expression *ast.Num) (int, error) {
	number, err := strconv.Atoi(expression.Lexeme)
	if err != nil {
		return 0, err
	}

	return number, nil
}
