package scanner_test

import (
	"github.com/njirem95/simple-pascal/pkg/scanner"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

var unexpectedTokenError = "unexpected token"

func TestScanner_Advance(t *testing.T) {
	input := "1+2"
	expected := [4]string{"1", "+", "2", ""}

	lexer, err := scanner.New(input)
	assert.Nil(t, err)

	for _, current := range expected {
		assert.Equal(t, current, lexer.Current, "invalid lexeme")
		lexer.Advance()
	}
}

func TestScanner_Next_AddSub(t *testing.T) {
	lexer, _ := scanner.New("1 +  2 - 1")
	expected := []token.Token{
		{
			Type:   token.Int,
			Lexeme: "1",
		},
		{
			Type:   token.Add,
			Lexeme: "+",
		},
		{
			Type:   token.Int,
			Lexeme: "2",
		},
		{
			Type:   token.Sub,
			Lexeme: "-",
		},
		{
			Type:   token.Int,
			Lexeme: "1",
		},
		{
			Type:   token.EOF,
			Lexeme: "",
		},
	}

	for _, next := range expected {
		assert.Equal(t, next, lexer.Next(), unexpectedTokenError)
	}
}

func TestScanner_Next_MulDiv(t *testing.T) {
	input := "5 * 2 / 3 + 8 - 5"
	expected := []token.Token{
		{
			Type:   token.Int,
			Lexeme: "5",
		},
		{
			Type:   token.Mul,
			Lexeme: "*",
		},
		{
			Type:   token.Int,
			Lexeme: "2",
		},
		{
			Type:   token.Div,
			Lexeme: "/",
		},
		{
			Type:   token.Int,
			Lexeme: "3",
		},
		{
			Type:   token.Add,
			Lexeme: "+",
		},
		{
			Type:   token.Int,
			Lexeme: "8",
		},
		{
			Type:   token.Sub,
			Lexeme: "-",
		},
		{
			Type:   token.Int,
			Lexeme: "5",
		},
		{
			Type:   token.EOF,
			Lexeme: "",
		},
	}

	lexer, err := scanner.New(input)

	assert.Nil(t, err)
	for _, next := range expected {
		assert.Equal(t, next, lexer.Next(), unexpectedTokenError)
	}
}

func TestScanner_Next_WithParentheses(t *testing.T) {
	input := "(10 + 5) * (9 / 2 * (5 - 3))"
	expected := []token.Token{
		{
			Type:   token.Lparen,
			Lexeme: "(",
		},
		{
			Type:   token.Int,
			Lexeme: "10",
		},
		{
			Type:   token.Add,
			Lexeme: "+",
		},
		{
			Type:   token.Int,
			Lexeme: "5",
		},
		{
			Type:   token.Rparen,
			Lexeme: ")",
		},
		{
			Type:   token.Mul,
			Lexeme: "*",
		},
		{
			Type:   token.Lparen,
			Lexeme: "(",
		},
		{
			Type:   token.Int,
			Lexeme: "9",
		},
		{
			Type:   token.Div,
			Lexeme: "/",
		},
		{
			Type:   token.Int,
			Lexeme: "2",
		},
		{
			Type:   token.Mul,
			Lexeme: "*",
		},
		{
			Type:   token.Lparen,
			Lexeme: "(",
		},
		{
			Type:   token.Int,
			Lexeme: "5",
		},
		{
			Type:   token.Sub,
			Lexeme: "-",
		},
		{
			Type:   token.Int,
			Lexeme: "3",
		},
		{
			Type:   token.Rparen,
			Lexeme: ")",
		},
		{
			Type:   token.Rparen,
			Lexeme: ")",
		},
		{
			Type:   token.EOF,
			Lexeme: "",
		},
	}

	lexer, err := scanner.New(input)

	assert.Nil(t, err)
	for _, next := range expected {
		assert.Equal(t, next, lexer.Next(), unexpectedTokenError)
	}
}

func TestNew_Peek(t *testing.T) {
	input := "15 + 2"
	expected := []string{"+", "2", ""}

	lexer, err := scanner.New(input)
	nextToken := lexer.Next()

	assert.Nil(t, err)
	for _, next := range expected {
		nextToken = lexer.Next()
		assert.Equal(t, nextToken.Lexeme, next, unexpectedTokenError)
	}
}

func TestNew(t *testing.T) {
	expected := "1"
	scan, err := scanner.New(expected)

	assert.Nil(t, err)
	assert.Equal(t, expected, scan.Stream, "unable to instantiate")
}

func TestNew_EmptyInput(t *testing.T) {
	input := ""
	_, err := scanner.New(input)
	assert.NotNil(t, err)
}
