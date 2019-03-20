package parser_test

import (
	"github.com/njirem95/simple-pascal/ast"
	"github.com/njirem95/simple-pascal/parser"
	"github.com/njirem95/simple-pascal/scanner"
	"github.com/njirem95/simple-pascal/scanner/token"
	"reflect"
	"testing"
)

func TestParser_Consume(t *testing.T) {
	lexer, err := scanner.New("1 + 1 - 1 * 3")
	if err != nil {
		t.Error(err)
	}
	parser := parser.New(lexer)
	if err != nil {
		t.Error(err)
	}
	err = parser.Consume(token.Int)
	if err != nil {
		t.Error("failed to consume integer")
	}
	err = parser.Consume(token.Add)
	if err != nil {
		t.Error("failed to consume addition operator")
	}
	err = parser.Consume(token.Int)
	if err != nil {
		t.Error("failed to consume integer")
	}
	err = parser.Consume(token.Sub)
	if err != nil {
		t.Error("failed to consume subtraction operator")
	}
	err = parser.Consume(token.Int)
	if err != nil {
		t.Error("failed to consume integer")
	}
	err = parser.Consume(token.Mul)
	if err != nil {
		t.Error("failed to consume multiplication operator")
	}
	err = parser.Consume(token.Int)
	if err != nil {
		t.Error("failed to consume integer")
	}
	err = parser.Consume(token.EOF)
	if err != nil {
		t.Error("failed to consume EOF")
	}
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
	lexer, err := scanner.New("+20")
	if err != nil {
		t.Error(err)
	}
	parser := parser.New(lexer)
	expression, err := parser.Factor()
	if err != nil {
		t.Error(err)
	}

	// sanity check
	node, ok := expression.(*ast.UnaryOp)
	if !ok {
		t.Fatal("expected *ast.UnaryOp")
	}

	if node.Expression.Lexeme != "20" {
		t.Error("expected lexeme to be 20")
	}

	expected := token.Token{
		Type:   token.Add,
		Lexeme: "+",
	}
	if !reflect.DeepEqual(expected, node.Operator) {
		t.Error("expected token to be token.Add")
	}
}

func TestParser_Factor_TestUnarySub(t *testing.T) {
	lexer, err := scanner.New("-20")
	if err != nil {
		t.Error(err)
	}
	parser := parser.New(lexer)
	expression, err := parser.Factor()
	if err != nil {
		t.Error(err)
	}

	// sanity check
	node, ok := expression.(*ast.UnaryOp)
	if !ok {
		t.Fatal("expected *ast.UnaryOp")
	}

	if node.Expression.Lexeme != "20" {
		t.Error("expected lexeme to be 20")
	}

	expected := token.Token{
		Type:   token.Sub,
		Lexeme: "-",
	}
	if !reflect.DeepEqual(expected, node.Operator) {
		t.Error("expected token to be token.Sub")
	}
}
