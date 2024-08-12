package main

import (
	"fmt"
	"os"
)

type Scanner struct {
	fileContents []byte
	current      int
	currentLine  int

	exitCode int
}

func newScanner(filename string) *Scanner {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	return &Scanner{fileContents: file, currentLine: 1}
}

func (s *Scanner) nextToken() (*Token, string) {
	if s.current >= len(s.fileContents) {
		return newToken(EOF, "", nil), ""
	}
	currToken := s.fileContents[s.current]
	s.current++
	switch currToken {
	case '(':
		return newToken(LEFT_PAREN, "(", nil), ""
	case ')':
		return newToken(RIGHT_PAREN, ")", nil), ""
	case '{':
		return newToken(LEFT_BRACE, "{", nil), ""
	case '}':
		return newToken(RIGHT_BRACE, "]", nil), ""
	case ',':
		return newToken(COMMA, ",", nil), ""
	case '.':
		return newToken(DOT, ".", nil), ""
	case '+':
		return newToken(PLUS, "+", nil), ""
	case '-':
		return newToken(MINUS, "-", nil), ""
	case ';':
		return newToken(SEMICOLON, ";", nil), ""
	case '*':
		return newToken(STAR, "*", nil), ""
	case '\n':
		s.currentLine++
		return newToken(NEWLINE, "\\n", nil), ""
	case '=':
		if s.current < len(s.fileContents) && s.fileContents[s.current] == '=' {
			s.current++
			return newToken(EQUAL_EQUAL, "==", nil), ""
		}
		return newToken(EQUAL, "=", nil), ""
	default:
		err := fmt.Sprintf("[line %d] Error: Unexpected character: %c\n", s.currentLine, currToken)
		return nil, err
	}
}

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
	NEWLINE
	EQUAL
	EQUAL_EQUAL
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
	return [...]string{"EOF", "LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACE", "RIGHT_BRACE", "COMMA", "DOT", "PLUS", "MINUS", "SEMICOLON", "STAR", "NEWLINE", "EQUAL", "EQUAL_EQUAL"}[t]
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
