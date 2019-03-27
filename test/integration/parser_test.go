package integration

import (
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/parser"
	"github.com/njirem95/simple-pascal/pkg/scanner"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestParser_Expr tests the arithmetic operators '+' (addition) and '-' (subtraction)
func TestParser_Expr(t *testing.T) {
	expected := []ast.BinOp{
		{
			Left: &ast.Num{
				Token: token.Token{
					Type:   token.Int,
					Lexeme: "20",
				},
				Lexeme: "20",
			},
			Operator: token.Token{
				Type:   token.Add,
				Lexeme: "+",
			},
			Right: &ast.Num{
				Token: token.Token{
					Type:   token.Int,
					Lexeme: "15",
				},
				Lexeme: "15",
			},
		},
		{
			Left: &ast.Num{
				Token: token.Token{
					Type:   token.Int,
					Lexeme: "20",
				},
				Lexeme: "20",
			},
			Operator: token.Token{
				Type:   token.Sub,
				Lexeme: "-",
			},
			Right: &ast.Num{
				Token: token.Token{
					Type:   token.Int,
					Lexeme: "15",
				},
				Lexeme: "15",
			},
		},
	}
	inputs := [2]string{"20 + 15", "20 - 15"}
	for index, input := range inputs {
		lexer, err := scanner.New(input)
		assert.Nil(t, err)

		parser := parser.New(lexer)
		expression, err := parser.Expr()
		assert.Nil(t, err)

		// sanity check
		operation, ok := expression.(*ast.BinOp)
		assert.True(t, ok)

		assert.Equal(t, expected[index].Operator, operation.Operator)
		assert.Equal(t, expected[index].Right, operation.Right)
		assert.Equal(t, expected[index].Left, operation.Left)
	}
}

// TestParser_Term tests the arithmetic operators '*' (multiplication) and '/' (division)
func TestParser_Term(t *testing.T) {
	expected := []ast.BinOp{
		{
			Left: &ast.Num{
				Token: token.Token{
					Type:   token.Int,
					Lexeme: "20",
				},
				Lexeme: "20",
			},
			Operator: token.Token{
				Type:   token.Mul,
				Lexeme: "*",
			},
			Right: &ast.Num{
				Token: token.Token{
					Type:   token.Int,
					Lexeme: "5",
				},
				Lexeme: "5",
			},
		},
		{
			Left: &ast.Num{
				Token: token.Token{
					Type:   token.Int,
					Lexeme: "20",
				},
				Lexeme: "20",
			},
			Operator: token.Token{
				Type:   token.Div,
				Lexeme: "/",
			},
			Right: &ast.Num{
				Token: token.Token{
					Type:   token.Int,
					Lexeme: "5",
				},
				Lexeme: "5",
			},
		},
	}
	inputs := [2]string{"20 * 5", "20 / 5"}
	for index, input := range inputs {
		lexer, err := scanner.New(input)
		assert.Nil(t, err)

		parser := parser.New(lexer)
		expression, err := parser.Term()
		assert.Nil(t, err)

		// sanity check
		operation, ok := expression.(*ast.BinOp)
		assert.True(t, ok)

		assert.Equal(t, expected[index].Operator, operation.Operator)
		assert.Equal(t, expected[index].Right, operation.Right)
		assert.Equal(t, expected[index].Left, operation.Left)
	}
}
