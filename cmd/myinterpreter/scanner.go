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

	return &Scanner{fileContents: file, currentLine: 1, exitCode: 0}
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
		return newToken(RIGHT_BRACE, "}", nil), ""
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
	case '=':
		if s.current < len(s.fileContents) && s.fileContents[s.current] == '=' {
			s.current++
			return newToken(EQUAL_EQUAL, "==", nil), ""
		}
		return newToken(EQUAL, "=", nil), ""
	case '!':
		if s.current < len(s.fileContents) && s.fileContents[s.current] == '=' {
			s.current++
			return newToken(BANG_EQUAL, "!=", nil), ""
		}
		return newToken(BANG, "!", nil), ""
	case '>':
		if s.current < len(s.fileContents) && s.fileContents[s.current] == '=' {
			s.current++
			return newToken(GREATER_EQUAL, ">=", nil), ""
		}
		return newToken(GREATER, ">", nil), ""
	case '<':
		if s.current < len(s.fileContents) && s.fileContents[s.current] == '=' {
			s.current++
			return newToken(LESS_EQUAL, "<=", nil), ""
		}
		return newToken(LESS, "<", nil), ""
	case '/':
		if s.current < len(s.fileContents) && s.fileContents[s.current] == '/' {
			s.current++ //we have found a comment and should consume until the line ends
			s.readComment()
			return nil, ""
		} else {
			return newToken(SLASH, "/", nil), ""
		}
	case '"':
		if literal, err := s.readString(); err {
			errMsg := fmt.Sprintf("[line %d] Error: Unterminated string.\n", s.currentLine)
			s.exitCode = 65
			return nil, errMsg
		} else {
			return newToken(STRING, literal, literal), ""
		}

	case '\n':
		s.currentLine++
		return nil, ""
	case ' ', '\r', '\t':
		return nil, ""
	default:
		err := fmt.Sprintf("[line %d] Error: Unexpected character: %c\n", s.currentLine, currToken)
		s.exitCode = 65
		return nil, err
	}
}

func (s *Scanner) readString() (literal string, err bool) {

	terminated := false
	literal += "\""

	for s.current < len(s.fileContents) && !terminated {
		currentChar := s.fileContents[s.current]
		if currentChar == '"' { //we have reached the end of the string literal
			terminated = true
		} else {
			if currentChar == '\n' {
				s.currentLine++
			}
			literal += (string)(currentChar) //add the character to the end of the string and proceed
		}
		s.current++
	}

	if !terminated { //don't return a string at all and report the error
		return "", true
	}

	literal += "\""
	return literal, false
}

func (s *Scanner) readComment() {
	for ; s.current < len(s.fileContents) && s.fileContents[s.current] != '\n'; s.current++ {
		//do nothing until comment is done
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
		"GREATER", "GREATER_EQUAL", "SLASH", "STRING"}[t]
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
