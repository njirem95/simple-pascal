package integration

import (
	"github.com/njirem95/simple-pascal/pkg/ast"
	parser2 "github.com/njirem95/simple-pascal/pkg/parser"
	"github.com/njirem95/simple-pascal/pkg/scanner"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestParser_Expr tests the integration between the lexer and the parser
// and parses the expression 10 + (5 - 3)
func TestParser_Expr(t *testing.T) {
	expr := "10 + (5 - 3)"
	expected := &ast.BinOp{
		Left: &ast.Num{
			Token: token.Token{
				Type:   token.Int,
				Lexeme: "10",
			},
			Lexeme: "10",
		},
		Operator: token.Token{
			Type:   token.Add,
			Lexeme: "+",
		},
		Right: &ast.BinOp{
			Left: &ast.Num{
				Token: token.Token{
					Type:   token.Int,
					Lexeme: "5",
				},
				Lexeme: "5",
			},
			Operator: token.Token{
				Type:   token.Sub,
				Lexeme: "-",
			},
			Right: &ast.Num{
				Token: token.Token{
					Type:   token.Int,
					Lexeme: "3",
				},
				Lexeme: "3",
			},
		},
	}

	lexer, err := scanner.New(expr)
	assert.Nil(t, err)

	parser := parser2.New(lexer)
	expre, err := parser.Expr()
	assert.Nil(t, err)

	expression, ok := expre.(*ast.BinOp)
	assert.True(t, ok)

	assert.Equal(t, expression, expected)
}
