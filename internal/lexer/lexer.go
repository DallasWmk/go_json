// Package lexer provides a simple lexer for JSON-like structures.
package lexer

import (
	"bufio"
	"io"
	"unicode"
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
	SingleQuote
	DoubleQuote
	Int
)

var Tokens = []string{
	EOF:          "EOF",
	CurlyOpen:    "{",
	CurlyClose:   "}",
	BracketOpen:  "[",
	BracketClose: "]",
	Colon:        ":",
	Comma:        ",",
	DoubleQuote:  "\"",
	SingleQuote:  "'",
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
		case '\'':
			startPos := l.pos
			l.backup()
			lit := l.lexQuote()
			return startPos, Int, lit
		case '"':
			startPos := l.pos
			l.backup()
			lit := l.lexQuote()
			return startPos, Int, lit
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				startPos := l.pos
				l.backup()
				lit := l.lexInt()
				return startPos, Int, lit
			} else if unicode.IsLetter(r) {
				switch r {
				case 't', 'f':
					startPos := l.pos
					l.backup()
					lit := l.lexBool()
					if lit == "true" || lit == "false" {
						return startPos, Int, lit
					} else {
						l.pos = startPos // Reset position if not a valid boolean
						continue
					}
				default:
					continue // Ignore other letters
				}
			}
		}
	}
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}
	l.pos.Column--
}

func (l *Lexer) lexQuote() string {
	var lit string
	singleQuote := 0
	dblQuote := 0
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// never reached end of string without closing quote
				return ""
			}
		}
		l.pos.Column++
		if r == '"' {
			dblQuote++
		}
		if r == '\'' {
			singleQuote++
		}
		lit += string(r)
		if singleQuote == 2 || dblQuote == 2 {
			break
		}
	}
	return lit
}

// lexInt reads a sequence of digits and returns the literal.
func (l *Lexer) lexInt() string {
	var lit string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}
		l.pos.Column++
		if unicode.IsDigit(r) {
			lit = lit + string(r)
		} else {
			// at the end of the integer, backtrack one rune
			l.backup()
			return lit
		}
	}
}

// lexBool reads a boolean value (true or false) and returns the literal.
func (l *Lexer) lexBool() string {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return ""
			}
		}
		l.pos.Column++
		switch r {
		case 't':
			if string(r)+l.peek(3) == "true" {
				return "true"
			}
		case 'f':
			if string(r)+l.peek(4) == "false" {
				return "false"
			}
		}
	}
}

func (l *Lexer) peek(n int) string {
	var lit string
	for range n {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			panic(err)
		}
		lit += string(r)
	}
	return lit
}

func (l *Lexer) resetPosition() {
	l.pos.Line++
	l.pos.Column = 0
}
