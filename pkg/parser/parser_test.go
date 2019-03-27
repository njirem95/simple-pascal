package parser_test

import (
	"github.com/golang/mock/gomock"
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/parser"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
	"github.com/njirem95/simple-pascal/test/mock/scanner"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Consume(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_scanner.NewMockScanner(ctrl)

	expected := token.Token{
		Type: token.Int,
	}

	m.
		EXPECT().
		Next().
		Return(expected)

	parser := parser.New(m)

	expected = token.Token{
		Type: token.Add,
	}

	m.
		EXPECT().
		Next().
		Return(expected)

	assert.Nil(t, parser.Consume(token.Int))

	expected = token.Token{
		Type: token.Int,
	}
	m.
		EXPECT().
		Next().
		Return(expected)

	assert.Nil(t, parser.Consume(token.Add))

	expected = token.Token{
		Type: token.Sub,
	}
	m.
		EXPECT().
		Next().
		Return(expected)

	assert.Nil(t, parser.Consume(token.Int))

	expected = token.Token{
		Type: token.Int,
	}

	m.
		EXPECT().
		Next().
		Return(expected)

	assert.Nil(t, parser.Consume(token.Sub))

	expected = token.Token{
		Type: token.Mul,
	}

	m.
		EXPECT().
		Next().
		Return(expected)

	assert.Nil(t, parser.Consume(token.Int))

	expected = token.Token{
		Type: token.Int,
	}

	m.
		EXPECT().
		Next().
		Return(expected)

	assert.Nil(t, parser.Consume(token.Mul))

	expected = token.Token{
		Type: token.EOF,
	}

	m.
		EXPECT().
		Next().
		AnyTimes().
		Return(expected)

	assert.Nil(t, parser.Consume(token.Int))
}

func TestParser_Factor(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_scanner.NewMockScanner(ctrl)

	expected := token.Token{
		Type:   token.Int,
		Lexeme: "20",
	}

	m.
		EXPECT().
		Next().
		Return(expected)

	parser := parser.New(m)

	// I'm expecting Next() to be returning EOF from now on
	expected = token.Token{
		Type:   token.EOF,
		Lexeme: "",
	}
	m.EXPECT().
		Next().
		AnyTimes().
		Return(expected)

	expression, err := parser.Factor()
	assert.Nil(t, err)

	// sanity check
	num, ok := expression.(*ast.Num)
	assert.True(t, ok)

	assert.Equal(t, num.Token.Type, token.Int)
	assert.Equal(t, num.Token.Lexeme, "20")
	assert.Equal(t, num.Lexeme, "20")
}

func TestParser_Factor_TestUnaryAdd(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_scanner.NewMockScanner(ctrl)
	expected := token.Token{
		Type:   token.Add,
		Lexeme: "+",
	}
	m.
		EXPECT().
		Next().
		Return(expected)

	parser := parser.New(m)
	expected = token.Token{
		Lexeme: "20",
		Type:   token.Int,
	}
	m.
		EXPECT().
		Next().
		AnyTimes().
		Return(expected)
	expression, err := parser.Factor()
	if err != nil {
		t.Error(err)
	}

	// sanity check
	node, ok := expression.(*ast.UnaryOp)
	assert.True(t, ok)

	assert.Equal(t, "20", node.Expression.Lexeme)
	assert.Equal(t, "+", node.Operator.Lexeme)
	assert.Equal(t, token.Add, node.Operator.Type)
}

func TestParser_Factor_TestUnarySub(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_scanner.NewMockScanner(ctrl)
	expected := token.Token{
		Type:   token.Sub,
		Lexeme: "-",
	}
	m.
		EXPECT().
		Next().
		Return(expected)

	parser := parser.New(m)
	expected = token.Token{
		Lexeme: "20",
		Type:   token.Int,
	}
	m.
		EXPECT().
		Next().
		AnyTimes().
		Return(expected)
	expression, err := parser.Factor()
	if err != nil {
		t.Error(err)
	}

	// sanity check
	node, ok := expression.(*ast.UnaryOp)
	assert.True(t, ok)

	assert.Equal(t, "20", node.Expression.Lexeme)
	assert.Equal(t, "-", node.Operator.Lexeme)
	assert.Equal(t, token.Sub, node.Operator.Type)
}

func TestParser_Factor_TestLParenExprRparen(t *testing.T) {
	expected := &ast.BinOp{
		Left: &ast.Num{
			Token: token.Token{
				Type:   token.Int,
				Lexeme: "1",
			},
			Lexeme: "1",
		},
		Operator: token.Token{
			Type:   token.Mul,
			Lexeme: "*",
		},
		Right: &ast.Num{
			Token: token.Token{
				Type:   token.Int,
				Lexeme: "20",
			},
			Lexeme: "20",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tok := token.Token{
		Type:   token.Lparen,
		Lexeme: "(",
	}

	m := mock_scanner.NewMockScanner(ctrl)
	m.
		EXPECT().
		Next().
		Return(tok)

	parser := parser.New(m)

	tok = token.Token{
		Type:   token.Int,
		Lexeme: "1",
	}
	before := m.
		EXPECT().
		Next().
		Return(tok)

	tok = token.Token{
		Type:   token.Mul,
		Lexeme: "*",
	}

	before = m.
		EXPECT().
		Next().
		After(before).
		Return(tok)

	tok = token.Token{
		Type:   token.Int,
		Lexeme: "20",
	}
	before = m.
		EXPECT().
		Next().
		Return(tok).
		After(before)

	tok = token.Token{
		Type:   token.Rparen,
		Lexeme: ")",
	}
	before = m.
		EXPECT().
		Next().
		Return(tok).
		After(before)

	tok = token.Token{
		Type:   token.EOF,
		Lexeme: "",
	}

	m.
		EXPECT().
		Next().
		AnyTimes().
		Return(tok).
		After(before)

	expression, err := parser.Factor()
	assert.Nil(t, err)

	// sanity check
	node, ok := expression.(*ast.BinOp)
	assert.True(t, ok)

	assert.Equal(t, node.Left, expected.Left)
	assert.Equal(t, node.Right, expected.Right)
	assert.Equal(t, node.Operator, expected.Operator)
}
