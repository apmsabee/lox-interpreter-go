package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
			lexeme := "\"" + literal + "\""
			return newToken(STRING, lexeme, literal), ""
		}

	case '\n':
		s.currentLine++
		return nil, ""
	case ' ', '\r', '\t':
		return nil, ""
	default:
		if isDigit(currToken) {
			//need to be able to add 1 point of precision to whole numbers and truncate floating numbers that are .00... to .0
			//val will be the base return(the entire number, whether it has a decimal place or not)
			val := s.readNumber()
			interfaceVal := val

			if strings.Contains(val, ".") { //potentially truncate floating-point values
				floatVal, _ := strconv.ParseFloat(val, 64) //convert to float
				if math.Trunc(floatVal) == floatVal {      //dropping the decimal does nothing to the value, therefore all decimal places are 0
					//reduce to 1 point of precision
					interfaceVal, _, _ = strings.Cut(interfaceVal, ".") //take only the integer portion of the number
					interfaceVal += ".0"                                //append one degree of precision
				}
			} else {
				interfaceVal += ".0"
			}

			return newToken(NUMBER, val, interfaceVal), ""
		} else if isAlpha(currToken) {
			val := s.identifier()
			return newToken(IDENTIFIER, val, nil), ""
		} else {
			err := fmt.Sprintf("[line %d] Error: Unexpected character: %c\n", s.currentLine, currToken)
			s.exitCode = 65
			return nil, err
		}
	}
}

func isAlpha(currToken byte) bool {
	return (currToken >= 'a' && currToken <= 'z') ||
		(currToken >= 'A' && currToken <= 'Z') ||
		currToken == '_'
}

func isDigit(b byte) bool {
	return (b >= '0' && b <= '9')
}

func isAlphaNumeric(b byte) bool {
	return isAlpha(b) || isDigit(b)
}

func (s *Scanner) identifier() string {
	start := s.current - 1
	for isAlphaNumeric(s.peek()) {
		s.current++
	}
	return (string)(s.fileContents[start:s.current])
}

func (s *Scanner) readString() (literal string, err bool) {
	start := s.current
	terminated := false

	for s.current < len(s.fileContents) && !terminated {
		if s.peek() == '"' { //we have reached the end of the string literal
			terminated = true
		} else {
			if s.peek() == '\n' {
				s.currentLine++
			}
		}
		s.current++
	}

	if !terminated { //don't return a string at all and report the error
		return "", true
	}

	return (string)(s.fileContents[start : s.current-1]), false
}

func (s *Scanner) readNumber() (literal string) {
	start := s.current - 1

	for s.current < len(s.fileContents) && isDigit(s.peek()) {
		s.current++
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.current++ //consume the decimal place
		for s.current < len(s.fileContents) && isDigit(s.peek()) {
			s.current++ //move to next token
		}
	}

	return (string)(s.fileContents[start:s.current])
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.fileContents) {
		return 0
	} else {
		return s.fileContents[s.current+1]
	}
}

func (s *Scanner) peek() byte {
	if s.current >= len(s.fileContents) {
		return 0
	} else {
		return s.fileContents[s.current]
	}
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
	NUMBER
	IDENTIFIER
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
		"GREATER", "GREATER_EQUAL", "SLASH", "STRING", "NUMBER", "IDENTIFIER"}[t]
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
