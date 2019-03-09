// Package scanner is the lexical analyzer. It takes an input stream and converts it to a slice of tokens.
package scanner

import (
	"github.com/pkg/errors"
)

// Scanner is responsible for 'tokenizing' the input stream.
// For example:
// 1 will become Token{TokenType.Integer, "1"}
// 1 + 2 will become Token{TokenType.Integer, "1"}, Token{TokenType.Addition, "+"}, Token{TokenType.Integer, "2"}
type Scanner struct {
	Stream string
}

// New creates the struct Scanner. Returns nil if the stream size is empty.
func New(stream string) (*Scanner, error) {
	if len(stream) == 0 {
		return nil, errors.New("input stream is empty")
	}
	scanner := &Scanner{}
	scanner.Stream = stream
	return scanner, nil
}
