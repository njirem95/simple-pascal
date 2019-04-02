package integration

import (
	parser2 "github.com/njirem95/simple-pascal/pkg/parser"
	"github.com/njirem95/simple-pascal/pkg/scanner"
	"github.com/njirem95/simple-pascal/pkg/visitor"
	"github.com/stretchr/testify/assert"
	"testing"
)

// We'll be testing the integration of the visitor pattern. The visitor pattern
// 'visits' the abstract syntax tree; therefore interpreting the expression.
func TestVisitor_Expression(t *testing.T) {
	inputs := make(map[string]int) // [ast.Expr] is the given input and int is the result
	inputs["2"] = 2
	inputs["2 + 2"] = 4
	inputs["9 * 2 - 2 + 4"] = 20
	inputs["2 + 2 * 4"] = 10
	inputs["(2 + 2) * 4"] = 16
	inputs["(512 * 2) - (28 - (16 / 4))"] = 1000
	inputs["6 - - - + - 4"] = 10
	inputs["6 - - - + - (3 + 4) - +1"] = 12

	for input, result := range inputs {
		lexer, err := scanner.New(input)
		assert.Nil(t, err)

		parser := parser2.New(lexer)

		expression, err := parser.Expr()
		assert.Nil(t, err)

		visitor := visitor.Visitor{}
		visit, err := visitor.Visit(expression)
		assert.Nil(t, err)

		assert.Equal(t, result, visit)
	}
}
