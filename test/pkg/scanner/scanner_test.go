package scanner_test

import (
	"interpreter/pkg/scanner"
	"interpreter/pkg/scanner/token"
	"reflect"
	"testing"
)

func TestScanner_Advance(t *testing.T) {
	input := "1+2"

	scan, err := scanner.New(input)
	if err != nil {
		t.Fatal(err)
	}
	if scan.Current != "1" {
		t.Fatalf("expected lexeme to be 1, got %s", scan.Current)
	}
	scan.Advance()
	if scan.Current != "+" {
		t.Fatalf("expected lexeme to be +, got %s", scan.Current)
	}
	scan.Advance()
	if scan.Current != "2" {
		t.Fatalf("expected lexeme to be 2, got %s", scan.Current)
	}
	scan.Advance()
	if scan.Current != "" {
		t.Fatalf("expected lexeme to be empty, got %s", scan.Current)
	}
}

func TestScanner_Next_AddSub(t *testing.T) {
	scan, _ := scanner.New("1 +  2")
	expected := token.Token{
		Type:   token.Int,
		Lexeme: "1",
	}
	next := scan.Next()

	if !reflect.DeepEqual(expected, next) {
		t.Error("unexpected token")
	}
	expected = token.Token{
		Type:   token.Add,
		Lexeme: "+",
	}
	next = scan.Next()
	if !reflect.DeepEqual(expected, next) {
		t.Error("unexpected token")
	}
	next = scan.Next()
	expected = token.Token{
		Type:   token.Int,
		Lexeme: "2",
	}
	if !reflect.DeepEqual(expected, next) {
		t.Error("unexpected token")
	}
	next = scan.Next()
	expected = token.Token{
		Type:   token.EOF,
		Lexeme: "",
	}
	if !reflect.DeepEqual(expected, next) {
		t.Error("unexpected token")
	}
}

// TestNew checks if we can instantiate the scanner when we provide a valid input stream
func TestNew(t *testing.T) {
	expected := "1"
	scan, _ := scanner.New(expected)
	if scan.Stream != expected {
		t.Errorf("expected (%s) does not equal scanner.Stream (%s)", expected, scan.Stream)
	}
}

// TestNew_EmptyInput fails when the instantiation of the scanner succeeds, because
// the scanner should return an error if we provide an empty input stream
func TestNew_EmptyInput(t *testing.T) {
	input := ""
	_, err := scanner.New(input)
	if err == nil {
		t.Errorf("expected to receive an error, but err is nil")
	}
}
