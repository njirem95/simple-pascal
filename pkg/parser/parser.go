package parser

import (
	"errors"
	"fmt"
	"github.com/njirem95/simple-pascal/pkg/ast"
	"github.com/njirem95/simple-pascal/pkg/scanner"
	"github.com/njirem95/simple-pascal/pkg/scanner/token"
)

var (
	endReachedError = errors.New("reached end of recursion")
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
	return fmt.Errorf("unable to consume token %s", p.currentToken.Lexeme)
}

func (p *Parser) Program() ([]ast.Statement, error) {
	statements, err := p.CompoundStmt()
	if err != nil {
		return nil, err
	}

	err = p.Consume(token.Dot)
	if err != nil {
		return nil, err
	}

	return statements, nil
}

func (p *Parser) CompoundStmt() ([]ast.Statement, error) {
	err := p.Consume(token.Begin)
	if err != nil {
		return nil, err
	}

	statements, err := p.StmtList()
	if err != nil {
		return nil, err
	}

	err = p.Consume(token.End)
	if err != nil {
		return nil, err
	}
	return statements, nil
}

func (p *Parser) StmtList() ([]ast.Statement, error) {
	node, err := p.Statement()
	if err != nil {
		return nil, err
	}

	var statements []ast.Statement
	statements = append(statements, node)

	// Add the next statement to the list if the current token is a semicolon.
	for p.currentToken.Type == token.Semi {
		err := p.Consume(token.Semi)
		if err != nil {
			return nil, err
		}
		node, err := p.Statement()
		if err != nil {
			return nil, err
		}

		statements = append(statements, node)
	}

	return statements, nil
}

func (p *Parser) Statement() (ast.Statement, error) {
	switch p.currentToken.Type {
	case token.Identifier:
		return p.AssignmentStmt()
	case token.Begin:
		return p.CompoundStmt()
	default:
		return p.Empty()
	}
}

func (p *Parser) Empty() (*ast.Empty, error) {
	return &ast.Empty{}, nil
}

func (p *Parser) AssignmentStmt() (*ast.Assign, error) {
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

func (p *Parser) Variable() (*ast.Variable, error) {
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

	for p.currentToken.Type == token.Add || p.currentToken.Type == token.Sub {
		operator := p.currentToken
		err = p.Consume(p.currentToken.Type)
		if err != nil {
			return nil, err
		}
		left := node

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

	for p.currentToken.Type == token.Mul || p.currentToken.Type == token.Div {
		operator := p.currentToken
		err = p.Consume(p.currentToken.Type)
		if err != nil {
			return nil, err
		}

		left := node

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
