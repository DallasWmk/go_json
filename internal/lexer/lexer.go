// Package lexer provides a simple lexer for JSON-like structures.
package lexer

import (
	"bufio"
	"io"
)

type TokenType int

const (
	EOF = iota
	CurlyOpen
	CurlyClose
	BracketOpen
	BracketClose
	Colon
	Comma
)

var Tokens = []string{
	EOF:          "EOF",
	CurlyOpen:    "{",
	CurlyClose:   "}",
	BracketOpen:  "[",
	BracketClose: "]",
	Colon:        ":",
	Comma:        ",",
}

func (t TokenType) String() string {
	return Tokens[t]
}

type Position struct {
	Line   int
	Column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{Line: 1, Column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, TokenType, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}
		}
		l.pos.Column++

		switch r {
		case '\n':
			l.resetPosition()
		case '{':
			return l.pos, CurlyOpen, "{"
		case '}':
			return l.pos, CurlyClose, "}"
		case '[':
			return l.pos, BracketOpen, "["
		case ']':
			return l.pos, BracketClose, "]"
		case ':':
			return l.pos, Colon, ":"
		case ',':
			return l.pos, Comma, ","
		default:
			continue
		}
	}
}

func (l *Lexer) resetPosition() {
	l.pos.Line++
	l.pos.Column = 0
}
