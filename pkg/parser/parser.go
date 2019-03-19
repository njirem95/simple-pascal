package parser

import (
	"errors"
	"interpreter/pkg/ast"
	"interpreter/pkg/scanner"
	"interpreter/pkg/scanner/token"
)

type Parser struct {
	lexer        *scanner.Scanner
	currentToken token.Token
}

func (p *Parser) Consume(tokenType int) error {
	if tokenType == p.currentToken.Type {
		p.currentToken = p.lexer.Next()
		return nil
	}
	return errors.New("unable to consume token")
}

// factor : integer
func (p *Parser) Factor() (ast.Expr, error) {
	switch p.currentToken.Type {
	case token.Add:
		node := &ast.UnaryOp{
			Operator: p.currentToken,
		}
		err := p.Consume(token.Add)
		if err != nil {
			return nil, err
		}

		// assign expression
		factor, err := p.Factor()
		if err != nil {
			return nil, err
		}
		child, ok := factor.(*ast.Num)
		if !ok {
			return nil, errors.New("expected unary child to be a num")
		}
		node.Expression = child
		return node, nil
	case token.Int:
		node := &ast.Num{
			Token:  p.currentToken,
			Lexeme: p.currentToken.Lexeme,
		}
		err := p.Consume(token.Int)
		if err != nil {
			return nil, err
		}
		return node, nil
	}
	return nil, errors.New("end of factor reached")
}

// New creates the struct Parser.
func New(lexer *scanner.Scanner) *Parser {
	parser := &Parser{}
	parser.lexer = lexer
	parser.currentToken = parser.lexer.Next()
	return parser
}
