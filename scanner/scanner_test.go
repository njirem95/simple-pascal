package scanner_test

import (
	"interpreter/scanner"
	"interpreter/scanner/token"
	"reflect"
	"testing"
)

var (
	instantiationError   = "expected (%s) does not equal scanner.Stream (%s)"
	errorExpectedError   = "expected to receive an error, but err is nil"
	unexpectedTokenError = "didn't expect this token while getting next lexeme"
	currentLexemeError   = "expected lexeme to be %s, got %s"
	peekingError         = "peeking failed, got %s instead of %s"
)

func TestScanner_Advance(t *testing.T) {
	input := "1+2"

	lexer, err := scanner.New(input)
	if err != nil {
		t.Fatal(err)
	}
	expected := [4]string{"1", "+", "2", ""}
	for _, current := range expected {
		if current != lexer.Current {
			t.Errorf(currentLexemeError, current, lexer.Current)
		}
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
		if !reflect.DeepEqual(next, lexer.Next()) {
			t.Error(unexpectedTokenError)
		}
	}
}

func TestScanner_Next_MulDiv(t *testing.T) {
	input := "5 * 2 / 3 + 8 - 5"
	lexer, err := scanner.New(input)
	if err != nil {
		t.Error(err)
	}
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

	for _, next := range expected {
		if !reflect.DeepEqual(next, lexer.Next()) {
			t.Error(unexpectedTokenError)
		}
	}

}

func TestScanner_Next_WithParentheses(t *testing.T) {
	input := "(10 + 5) * (9 / 2 * (5 - 3))"
	lexer, err := scanner.New(input)
	if err != nil {
		t.Error(err)
	}
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
	for _, next := range expected {
		if !reflect.DeepEqual(next, lexer.Next()) {
			t.Error(unexpectedTokenError)
		}
	}
}

func TestNew_Peek(t *testing.T) {
	input := "15 + 2"
	lexer, err := scanner.New(input)
	if err != nil {
		t.Error(err)
	}

	nextToken := lexer.Next()
	expected := []string{"+", "2", ""}
	for _, next := range expected {
		nextToken = lexer.Next()
		if nextToken.Lexeme != next {
			t.Errorf(peekingError, nextToken.Lexeme, next)
		}
	}
}

// TestNew checks if we can instantiate the scanner when we provide a valid input stream
func TestNew(t *testing.T) {
	expected := "1"
	scan, _ := scanner.New(expected)
	if scan.Stream != expected {
		t.Errorf(instantiationError, expected, scan.Stream)
	}
}

// TestNew_EmptyInput fails when the instantiation of the scanner succeeds, because
// the scanner should return an error if we provide an empty input stream
func TestNew_EmptyInput(t *testing.T) {
	input := ""
	_, err := scanner.New(input)
	if err == nil {
		t.Error(errorExpectedError)
	}
}
