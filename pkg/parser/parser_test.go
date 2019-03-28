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

// TestParser_Consume tests the consumption of tokens.
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

// TestParser_Factor_Integer tests the consumption of integers.
func TestParser_Factor_Integer(t *testing.T) {
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

// TestParser_Expr_Addition tests the binary operation addition.
func TestParser_Expr_Addition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := &ast.BinOp{
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

	m := mock_scanner.NewMockScanner(ctrl)
	currentToken := token.Token{
		Type:   token.Int,
		Lexeme: "20",
	}
	m.
		EXPECT().
		Next().
		Return(currentToken)

	parser := parser.New(m)

	currentToken = token.Token{
		Type:   token.Add,
		Lexeme: "+",
	}

	// Creating mock
	after := m.EXPECT().Next().Return(currentToken)

	currentToken = token.Token{
		Type:   token.Int,
		Lexeme: "15",
	}

	after = m.
		EXPECT().
		Next().
		Return(currentToken).
		After(after)

	currentToken = token.Token{
		Type:   token.EOF,
		Lexeme: "",
	}

	m.
		EXPECT().
		Next().
		AnyTimes().
		Return(currentToken)

	expression, err := parser.Expr()
	assert.Nil(t, err)

	// sanity check
	operation, ok := expression.(*ast.BinOp)
	assert.True(t, ok)

	assert.Equal(t, expected.Operator, operation.Operator)
	assert.Equal(t, expected.Right, operation.Right)
	assert.Equal(t, expected.Left, operation.Left)
}

// TestParser_Expr_Addition tests the binary operation subtraction.
func TestParser_Expr_Subtraction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := &ast.BinOp{
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
	}

	m := mock_scanner.NewMockScanner(ctrl)
	currentToken := token.Token{
		Type:   token.Int,
		Lexeme: "20",
	}
	m.
		EXPECT().
		Next().
		Return(currentToken)

	parser := parser.New(m)

	currentToken = token.Token{
		Type:   token.Sub,
		Lexeme: "-",
	}

	// Creating mock
	after := m.EXPECT().Next().Return(currentToken)

	currentToken = token.Token{
		Type:   token.Int,
		Lexeme: "15",
	}

	after = m.EXPECT().Next().Return(currentToken).After(after)

	currentToken = token.Token{
		Type:   token.EOF,
		Lexeme: "",
	}

	m.
		EXPECT().
		Next().
		AnyTimes().
		Return(currentToken)

	expression, err := parser.Expr()
	assert.Nil(t, err)

	// sanity check
	operation, ok := expression.(*ast.BinOp)
	assert.True(t, ok)

	assert.Equal(t, expected.Operator, operation.Operator)
	assert.Equal(t, expected.Right, operation.Right)
	assert.Equal(t, expected.Left, operation.Left)
}

// TestParser_Expr_Addition tests the binary operation multiplication.
func TestParser_Term_Multiplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := &ast.BinOp{
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
				Lexeme: "15",
			},
			Lexeme: "15",
		},
	}

	m := mock_scanner.NewMockScanner(ctrl)
	currentToken := token.Token{
		Type:   token.Int,
		Lexeme: "20",
	}
	m.
		EXPECT().
		Next().
		Return(currentToken)

	parser := parser.New(m)

	currentToken = token.Token{
		Type:   token.Mul,
		Lexeme: "*",
	}

	// Creating mock
	after := m.
		EXPECT().
		Next().
		Return(currentToken)

	currentToken = token.Token{
		Type:   token.Int,
		Lexeme: "15",
	}

	after = m.EXPECT().Next().Return(currentToken).After(after)

	currentToken = token.Token{
		Type:   token.EOF,
		Lexeme: "",
	}

	m.
		EXPECT().
		Next().
		AnyTimes().
		Return(currentToken)

	expression, err := parser.Term()
	assert.Nil(t, err)

	// sanity check
	operation, ok := expression.(*ast.BinOp)
	assert.True(t, ok)

	assert.Equal(t, expected.Operator, operation.Operator)
	assert.Equal(t, expected.Right, operation.Right)
	assert.Equal(t, expected.Left, operation.Left)
}

// TestParser_Expr_Addition tests the binary operation division.
func TestParser_Term_Division(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := &ast.BinOp{
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
				Lexeme: "15",
			},
			Lexeme: "15",
		},
	}

	m := mock_scanner.NewMockScanner(ctrl)
	currentToken := token.Token{
		Type:   token.Int,
		Lexeme: "20",
	}
	m.
		EXPECT().
		Next().
		Return(currentToken)

	parser := parser.New(m)

	currentToken = token.Token{
		Type:   token.Div,
		Lexeme: "/",
	}

	// Creating mock
	after := m.EXPECT().Next().Return(currentToken)

	currentToken = token.Token{
		Type:   token.Int,
		Lexeme: "15",
	}

	after = m.
		EXPECT().
		Next().
		Return(currentToken).After(after)

	currentToken = token.Token{
		Type:   token.EOF,
		Lexeme: "",
	}

	m.
		EXPECT().
		Next().
		AnyTimes().
		Return(currentToken)

	expression, err := parser.Expr()
	assert.Nil(t, err)

	// sanity check
	operation, ok := expression.(*ast.BinOp)
	assert.True(t, ok)

	assert.Equal(t, expected.Operator, operation.Operator)
	assert.Equal(t, expected.Right, operation.Right)
	assert.Equal(t, expected.Left, operation.Left)
}

// TestParser_Factor_TestUnaryAdd tests the unary add operator.
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

	// sanity checks
	node, ok := expression.(*ast.UnaryOp)
	assert.True(t, ok)

	num, ok := node.Expression.(*ast.Num)
	assert.True(t, ok)

	assert.Equal(t, "20", num.Lexeme)
	assert.Equal(t, "+", node.Operator.Lexeme)
	assert.Equal(t, token.Add, node.Operator.Type)
}

// TestParser_Factor_TestUnarySub tests the unary sub operator.
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

	// sanity checks
	node, ok := expression.(*ast.UnaryOp)
	assert.True(t, ok)

	num, ok := node.Expression.(*ast.Num)
	assert.True(t, ok)

	assert.Equal(t, "20", num.Lexeme)
	assert.Equal(t, "-", node.Operator.Lexeme)
	assert.Equal(t, token.Sub, node.Operator.Type)
}

// TestParser_Factor_TestLParenExprRparen tests the 'lparen expr rparen' expression.
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
