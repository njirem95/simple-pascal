package visitor_test

import (
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
	"github.com/njirem95/simple-pascal/pkg/visitor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinOpVisitor_Visit(t *testing.T) {
	input := &ast.BinOp{
		Left: &ast.Num{
			Token: token.Token{
				Type:   token.Int,
				Lexeme: "12",
			},
			Lexeme: "12",
		},
		Operator: token.Token{
			Type:   token.Add,
			Lexeme: "+",
		},
		Right: &ast.Num{
			Token: token.Token{
				Type:   token.Int,
				Lexeme: "8",
			},
			Lexeme: "8",
		},
	}

	visitor := visitor.BinOpVisitor{}
	result, err := visitor.Visit(input)

	assert.Nil(t, err)
	assert.Equal(t, 20, result)
}
