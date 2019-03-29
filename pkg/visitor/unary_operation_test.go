package visitor_test

import (
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
	visitor2 "github.com/njirem95/simple-pascal/pkg/visitor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnaryVisitor_Visit(t *testing.T) {
	// TODO better testing
	input := -6
	result := &ast.UnaryOp{
		Operator: token.Token{
			Type:   token.Sub,
			Lexeme: "-",
		},
		Expression: &ast.Num{
			Lexeme: "6",
			Token: token.Token{
				Type:   token.Int,
				Lexeme: "6",
			},
		},
	}

	visitor := visitor2.UnaryVisitor{}
	res, err := visitor.Visit(result)
	assert.Nil(t, err)
	assert.Equal(t, input, res)

	input = +6
	result.Operator = token.Token{
		Type:   token.Add,
		Lexeme: "+",
	}

	res, err = visitor.Visit(result)
	assert.Nil(t, err)
	assert.Equal(t, input, res)
}
