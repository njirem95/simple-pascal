package visitor_test

import (
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
	"github.com/njirem95/simple-pascal/pkg/visitor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Visit_BinOp(t *testing.T) {
	expected := 35
	input := &ast.BinOp{
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
	}

	visitor := visitor.New()
	assert.Equal(t, expected, visitor.Visit(input))
}
