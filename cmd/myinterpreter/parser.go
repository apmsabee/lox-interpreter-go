package main

import "fmt"

type Parser struct {
	tokens   []Token
	current  int
	exitCode int
}

type ParseError struct {
	token   Token
	message string
}

// constructor
func newParser(tokens []Token) *Parser {
	return &Parser{
		tokens:   tokens,
		current:  0,
		exitCode: 0,
	}
}

func (p *Parser) Parse() (*Expr, error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(ParseError); ok {
				return
			}
			panic(r)
		}
	}()

	return p.expression(), nil
}

// expression type navigation
func (p *Parser) expression() *Expr {
	return p.equality()
}

func (p *Parser) equality() *Expr {
	expr := p.comparison()
	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = &Expr{
			exprType: BINARY,
			left:     expr,
			right:    right,
			operator: operator,
		}
	}
	return expr
}

func (p *Parser) comparison() *Expr {
	expr := p.term()

	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = &Expr{
			exprType: BINARY,
			operator: operator,
			left:     expr,
			right:    right,
		}
	}
	return expr
}

func (p *Parser) term() *Expr {
	expr := p.factor()

	for p.match(MINUS, PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = &Expr{
			exprType: BINARY,
			left:     expr,
			operator: operator,
			right:    right,
		}
	}
	return expr
}

func (p *Parser) factor() *Expr {
	expr := p.unary()
	for p.match(SLASH, STAR) {
		operator := p.previous()
		right := p.unary()
		expr = &Expr{
			exprType: BINARY,
			left:     expr,
			operator: operator,
			right:    right,
		}
	}

	return expr
}

func (p *Parser) unary() *Expr {
	for p.match(BANG, MINUS) {
		operator := p.previous()
		right := p.unary()
		return &Expr{
			exprType: UNARY,
			right:    right,
			operator: operator,
		}
	}
	return p.primary()
}

func (p *Parser) primary() *Expr {
	if p.match(FALSE) {
		return &Expr{
			exprType: LITERAL,
			value:    false,
		}
	}
	if p.match(TRUE) {
		return &Expr{
			exprType: LITERAL,
			value:    true,
		}
	}
	if p.match(NIL) {
		return &Expr{
			exprType: LITERAL,
			value:    p.previous().literal,
		}
	}
	if p.match(NUMBER) {

		return &Expr{
			exprType: LITERAL,
			value:    p.previous().literal,
		}
	}
	if p.match(STRING) {
		return &Expr{
			exprType: LITERAL,
			value:    p.previous().literal,
		}
	}
	if p.match(LEFT_PAREN) {
		expr := p.expression()
		fmt.Println(expr)
		fmt.Println(expr.left)
		fmt.Println(expr.right)
		fmt.Println(expr.value)
		p.consume(RIGHT_PAREN, "Expect ')' after expression.")
		return &Expr{
			exprType: GROUPING,
			left:     expr,
		}
	}
	panic(p.error(p.peek(), "Expect expression."))
}

// Token Traversal methods
func (p *Parser) previous() Token {
	return p.tokens[p.current-1]
}
func (p *Parser) peek() Token {
	return p.tokens[p.current]
}
func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}
func (p *Parser) isAtEnd() bool {
	return p.peek().Type == EOF
}

// AST Helper Methods
func (p *Parser) match(tokens ...TokenType) bool {
	for _, token := range tokens {
		if p.check(token) {
			p.current++
			return true
		}
	}
	return false
}
func (p *Parser) check(token TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == token
}

func (p *Parser) consume(token TokenType, message string) Token {
	if p.check(token) {
		return p.advance()
	}
	panic(p.error(p.peek(), message))
}

func (p *Parser) error(t Token, message string) ParseError {
	return ParseError{token: t, message: message}
}
