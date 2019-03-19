package parser_test

import (
	"interpreter/pkg/ast"
	"interpreter/pkg/parser"
	"interpreter/pkg/scanner"
	"interpreter/pkg/scanner/token"
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
		t.Fatal("expected *ast.Num")
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
		Type: token.Add,
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
		Type: token.Sub,
		Lexeme: "-",
	}
	if !reflect.DeepEqual(expected, node.Operator) {
		t.Error("expected token to be token.Sub")
	}
}