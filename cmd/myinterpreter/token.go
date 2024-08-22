package main

import "fmt"

type TokenType int

const (
	EOF TokenType = iota
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	PLUS
	MINUS
	SEMICOLON
	STAR
	EQUAL
	EQUAL_EQUAL
	BANG
	BANG_EQUAL
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
	SLASH
	STRING
	NUMBER
	IDENTIFIER
	AND
	CLASS
	ELSE
	FALSE
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
)

type Token struct {
	Type    TokenType
	lexeme  string
	literal interface{}
}

func newToken(t TokenType, lexeme string, literal interface{}) *Token {
	return &Token{Type: t, lexeme: lexeme, literal: literal}
}

func (t TokenType) String() string {
	return [...]string{"EOF", "LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACE", "RIGHT_BRACE", "COMMA", "DOT", "PLUS", "MINUS",
		"SEMICOLON", "STAR", "EQUAL", "EQUAL_EQUAL", "BANG", "BANG_EQUAL", "LESS", "LESS_EQUAL",
		"GREATER", "GREATER_EQUAL", "SLASH", "STRING", "NUMBER", "IDENTIFIER", "AND", "CLASS", "ELSE",
		"FALSE", "FOR", "IF", "NIL", "OR", "PRINT", "RETURN", "SUPER", "THIS", "TRUE", "VAR", "WHILE"}[t]
}

func (t *Token) String() string {
	s := fmt.Sprintf("%s %s ", t.Type, t.lexeme)
	if t.literal != nil {
		s += fmt.Sprintf("%v", t.literal)
	} else {
		s += "null"
	}
	return s
}
