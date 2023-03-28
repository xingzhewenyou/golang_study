package lexer

import "unicode"

package lexer

import (
"fmt"
"unicode"
"unicode/utf8"
)

// TokenType represents the type of a token.
type TokenType int

const (
	// ILLEGAL represents an illegal token.
	ILLEGAL TokenType = iota
	// EOF represents the end of file.
	EOF
	// INT represents an integer constant.
	INT
	// ASSIGN represents an assignment operator.
	ASSIGN
	// PLUS represents an addition operator.
	PLUS
	// MINUS represents a subtraction operator.
	MINUS
	// BANG represents a logical NOT operator.
	BANG
	// ASTERISK represents a multiplication operator.
	ASTERISK
	// SLASH represents a division operator.
	SLASH
	// LT represents a less-than operator.
	LT
	// GT represents a greater-than operator.
	GT
	// EQ represents an equality operator.
	EQ
	// NEQ represents a non-equality operator.
	NEQ
	// COMMA represents a comma separator.
	COMMA
	// SEMICOLON represents a semicolon separator.
	SEMICOLON
	// LPAREN represents a left parenthesis.
	LPAREN
	// RPAREN represents a right parenthesis.
	RPAREN
	// LBRACE represents a left brace.
	LBRACE
	// RBRACE represents a right brace.
	RBRACE
	// FUNCTION represents the function keyword.
	FUNCTION
	// LET represents the let keyword.
	LET
	// TRUE represents the true boolean constant.
	TRUE
	// FALSE represents the false boolean constant.
	FALSE
	// IF represents the if keyword.
	IF
	// ELSE represents the else keyword.
	ELSE
)

// Token represents a token in the source code.
type Token struct {
	Type    TokenType // The type of the token.
	Literal string    // The literal value of the token.
}

// Lexer represents a lexer for the Monkey programming language.
type Lexer struct {
	input        string  // The input string.
	position     int     // The current position in the input.
	readPosition int     // The current read position in the input.
	ch           rune    // The current character being read.
}

// New creates a new lexer for the given input string.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken reads the next token from the input and returns it.
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(PLUS, l.ch)
	case '-':
		tok = newToken(MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: NEQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(BANG, l.ch)
		}
	case '*':
		tok = newToken(ASTERISK, l.ch)
	case '/':
		tok = newToken(SLASH, l.ch)
	case '<':
		tok = newToken(LT, l.ch)
	case '>':
		tok = newToken(GT, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case '{':
		tok = newToken(LBRACE, l.ch)
	case '}':
		tok = newToken(RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = lookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok

}

// readChar reads the next character from the input.
// It sets l.ch to 0 if the end of the input has been reached.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.readPosition])
	}

	l.position = l.readPosition
	l.readPosition += 1

}

// peekChar returns the next character from the input without consuming it.
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return rune(l.input[l.readPosition])
}

// readIdentifier reads an identifier from the input.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads a number from the input.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace skips whitespace characters in the input.
func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

// newToken creates a new Token with the given TokenType and literal value.
func newToken(tokenType TokenType, ch rune) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// isLetter returns true if the given rune is a letter or underscore.
func isLetter(ch rune) bool {
	return unicode
}


