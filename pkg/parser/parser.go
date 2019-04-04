package parser

import (
	"errors"
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/scanner"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
)

var (
	consumeTokenError = errors.New("unable to consume token")
	endReachedError   = errors.New("reached end of recursion")
)

type Parser struct {
	lexer        scanner.Scanner
	currentToken token.Token
}

func (p *Parser) Consume(tokenType int) error {
	if tokenType == p.currentToken.Type {
		p.currentToken = p.lexer.Next()
		return nil
	}
	return consumeTokenError
}

func (p *Parser) AssignmentStmt() (ast.Expr, error) {
	left, err := p.Variable()
	if err != nil {
		return nil, err
	}
	operator := p.currentToken
	err = p.Consume(token.Assign)
	if err != nil {
		return nil, err
	}
	right, err := p.Expr()
	if err != nil {
		return nil, err
	}

	node := &ast.Assign{
		Left:     left,
		Operator: operator,
		Right:    right,
	}

	return node, nil
}

func (p *Parser) Variable() (ast.Expr, error) {
	node := &ast.Variable{
		Name: p.currentToken.Lexeme,
		Token: token.Token{
			Type:   token.Identifier,
			Lexeme: p.currentToken.Lexeme,
		},
	}

	err := p.Consume(token.Identifier)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (p *Parser) Expr() (ast.Expr, error) {
	node, err := p.Term()
	if err != nil {
		return nil, err
	}
	// Check if current token is of type addition or subtraction
	for p.currentToken.Type == token.Add || p.currentToken.Type == token.Sub {
		operator := p.currentToken
		err = p.Consume(p.currentToken.Type)
		if err != nil {
			return nil, err
		}
		left := node

		// Get the value on the right
		node, err = p.Term()
		if err != nil {
			return nil, err
		}

		right := node

		node = &ast.BinOp{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return node, nil
}

func (p *Parser) Term() (ast.Expr, error) {
	node, err := p.Factor()
	if err != nil {
		return nil, err
	}
	// Check if current token is of type multiplication or division
	for p.currentToken.Type == token.Mul || p.currentToken.Type == token.Div {
		operator := p.currentToken
		err = p.Consume(p.currentToken.Type)
		if err != nil {
			return nil, err
		}

		left := node

		// Get the value on the right
		right, err := p.Factor()
		if err != nil {
			return nil, err
		}
		node = &ast.BinOp{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return node, nil
}

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

		factor, err := p.Factor()
		if err != nil {
			return nil, err
		}
		node.Expression = factor
		return node, nil
	case token.Sub:
		node := &ast.UnaryOp{
			Operator: p.currentToken,
		}
		err := p.Consume(token.Sub)
		if err != nil {
			return nil, err
		}
		factor, err := p.Factor()
		if err != nil {
			return nil, err
		}
		node.Expression = factor
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
	case token.Lparen:
		err := p.Consume(token.Lparen)
		if err != nil {
			return nil, err
		}
		expr, err := p.Expr()
		if err != nil {
			return nil, err
		}

		err = p.Consume(token.Rparen)
		if err != nil {
			return nil, err
		}
		return expr, nil
	case token.Identifier:
		return p.Variable()
	}
	return nil, endReachedError
}

// New creates the struct Parser.
func New(lexer scanner.Scanner) *Parser {
	parser := &Parser{}
	parser.lexer = lexer
	parser.currentToken = parser.lexer.Next()
	return parser
}
