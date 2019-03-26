package parser_test

import (
	"github.com/golang/mock/gomock"
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/parser"
	"github.com/njirem95/simple-pascal/pkg/scanner"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
	"github.com/njirem95/simple-pascal/test/mock/scanner"
	"github.com/stretchr/testify/assert"
	"reflect"
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
		if err != nil {
			t.Error(err)
		}
		parser := parser.New(lexer)
		expression, err := parser.Expr()
		if err != nil {
			t.Error(err)
		}

		// sanity check
		operation, ok := expression.(*ast.BinOp)
		if !ok {
			t.Fatalf("expected *ast.BinOp, got %s", reflect.TypeOf(expression))
		}

		// Take the memory location of expected[index] and check if it matches with the expression.
		if !reflect.DeepEqual(&expected[index], operation) {
			t.Error("binary operation does not match expected")
		}
	}
}

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
		if err != nil {
			t.Error(err)
		}
		parser := parser.New(lexer)
		expression, err := parser.Term()
		if err != nil {
			t.Error(err)
		}

		// sanity check
		operation, ok := expression.(*ast.BinOp)
		if !ok {
			t.Fatalf("expected *ast.BinOp, got %s", reflect.TypeOf(expression))
		}

		// Take the memory location of expected[index] and check if it matches with the expression.
		if !reflect.DeepEqual(&expected[index], operation) {
			t.Error("binary operation does not match expected")
		}
	}
}

func TestParser_Factor(t *testing.T) {
	lexer, err := scanner.New("20")
	if err != nil {
		t.Error(err)
	}
	parser := parser.New(lexer)
	expression, err := parser.Factor()
	if err != nil {
		t.Error(err)
	}

	// sanity check
	num, ok := expression.(*ast.Num)
	if !ok {
		t.Fatalf("expected *ast.Num, got %s", reflect.TypeOf(expression))
	}

	if num.Lexeme != "20" {
		t.Errorf("expected lexeme to be 20, got %s instead", num.Lexeme)
	}

	expected := token.Token{
		Type:   token.Int,
		Lexeme: "20",
	}
	if !reflect.DeepEqual(expected, num.Token) {
		t.Error("expected does not equal num.Token")
	}
}

func TestParser_Factor_TestUnaryAdd(t *testing.T) {
	ctrl := gomock.NewController(t)

	m := mock_scanner.NewMockScanner(ctrl)
	expected := token.Token{
		Type: token.Add,
		Lexeme: "+",
	}
	m.
		EXPECT().
		Next().
		Return(expected)

	parser := parser.New(m)
	expected = token.Token{
		Lexeme: "20",
		Type: token.Int,
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

	m := mock_scanner.NewMockScanner(ctrl)
	expected := token.Token{
		Type: token.Sub,
		Lexeme: "-",
	}
	m.
		EXPECT().
		Next().
		Return(expected)

	parser := parser.New(m)
	expected = token.Token{
		Lexeme: "20",
		Type: token.Int,
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
	lexer, err := scanner.New("(1 * 20)")
	if err != nil {
		t.Error(err)
	}
	parser := parser.New(lexer)
	expression, err := parser.Factor()
	if err != nil {
		t.Error(err)
	}

	// sanity check
	node, ok := expression.(*ast.BinOp)
	if !ok {
		t.Fatal("expected *ast.UnaryOp")
	}

	if !reflect.DeepEqual(node, expected) {
		t.Error("node does not equal expected")
	}
}
