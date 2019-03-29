package visitor_test

import (
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
	"github.com/njirem95/simple-pascal/pkg/visitor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumVisitor_Visit(t *testing.T) {
	input := &ast.Num{
		Token: token.Token{
			Type:   token.Int,
			Lexeme: "12",
		},
		Lexeme: "12",
	}

	visitor := visitor.NumVisitor{}
	result, err := visitor.Visit(input)

	assert.Nil(t, err)
	assert.Equal(t, 12, result)

	input.Lexeme = "1.5"
	result, err = visitor.Visit(input)
	assert.NotNil(t, err)
}
