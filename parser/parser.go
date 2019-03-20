package parser

import (
	"errors"
	"interpreter/ast"
	"interpreter/scanner"
	"interpreter/scanner/token"
)

var (
	consumeTokenError = errors.New("unable to consume token")
	endReachedError   = errors.New("reached end of recursion")
	unaryNumError     = errors.New("expected unary child to be a num")
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
	return consumeTokenError
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

		// We expect left to be of type num.
		left, ok := node.(*ast.Num)
		if !ok {
			return nil, errors.New("expected type num")
		}

		// Get the value on the right
		node, err = p.Term()
		if err != nil {
			return nil, err
		}

		right, ok := node.(*ast.Num)
		if !ok {
			return nil, errors.New("expected type num")
		}

		node = &ast.BinOp{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
		return node, nil
	}
	return node, nil
}

// term : factor ((mul|div) factor)*
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

		// We expect left to be of type num.
		left, ok := node.(*ast.Num)
		if !ok {
			return nil, errors.New("expected type num")
		}

		// Get the value on the right
		node, err = p.Factor()
		if err != nil {
			return nil, err
		}

		right, ok := node.(*ast.Num)
		if !ok {
			return nil, errors.New("expected type num")
		}

		node = &ast.BinOp{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
		return node, nil
	}
	return node, nil
}

// factor : integer
//			| add factor
//			| sub factor
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
		child, ok := factor.(*ast.Num)
		if !ok {
			return nil, unaryNumError
		}
		node.Expression = child
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
		child, ok := factor.(*ast.Num)
		if !ok {
			return nil, unaryNumError
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
	return nil, endReachedError
}

// New creates the struct Parser.
func New(lexer *scanner.Scanner) *Parser {
	parser := &Parser{}
	parser.lexer = lexer
	parser.currentToken = parser.lexer.Next()
	return parser
}
