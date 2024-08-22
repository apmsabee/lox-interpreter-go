package main

// type Parser struct {
// 	tokens  []Token
// 	current int
// }
// Parsing different expressions
// func (p *Parser) expression() Expr {
// 	return p.equality()
// }

// func (p *Parser) equality() Expr {
// 	expr := p.comparison()
// 	var newExpr Binary
// 	tokens := [2]TokenType{BANG_EQUAL, EQUAL_EQUAL}
// 	for p.match(tokens[:]) {
// 		operator := p.previous()
// 		right := p.comparison()
// 		newExpr.left = expr
// 		newExpr.operator = operator
// 		newExpr.right = right
// 	}
// 	return newExpr
// }

// func (p *Parser) comparison() Expr {
// 	panic("unimplemented")
// }

// // Token Traversal methods
// func (p *Parser) previous() Token {
// 	return p.tokens[p.current-1]
// }
// func (p *Parser) peek() Token {
// 	return p.tokens[p.current]
// }
// func (p *Parser) advance() Token {
// 	if !p.isAtEnd() {
// 		p.current++
// 	}
// 	return p.previous()
// }
// func (p *Parser) isAtEnd() bool {
// 	return p.peek().Type == EOF
// }

// // AST Helper Methods
// func (p *Parser) match(tokens []TokenType) bool {
// 	for _, token := range tokens {
// 		if p.check(token) {
// 			p.current++
// 			return true
// 		}
// 	}
// 	return false
// }
// func (p *Parser) check(token TokenType) bool {
// 	if p.isAtEnd() {
// 		return false
// 	}
// 	return p.peek().Type == token
// }
