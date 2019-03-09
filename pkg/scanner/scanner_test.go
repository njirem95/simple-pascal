package scanner_test

import (
	"interpreter/pkg/scanner"
	"testing"
)

// TestNew checks if we can instantiate the scanner when we provide a valid input stream.
func TestNew(t *testing.T) {
	expected := "1"
	scan, _ := scanner.New(expected)
	if scan.Stream != expected {
		t.Errorf("expected (%s) does not equal scanner.Stream (%s)", expected, scan.Stream)
	}
}

// TestNew_EmptyStream fails when the instantiation of the scanner succeeds, because
// the scanner should return an error if we provide an empty input stream.
func TestNew_EmptyStream(t *testing.T) {
	input := ""
	_, err := scanner.New(input)
	if err == nil {
		t.Errorf("expected to receive an error, but err is nil")
	}
}
