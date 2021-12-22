package scanner

import (
	"github.com/gramidt/mash-lang-for-codemash/grammar"
)

type Scanner struct {
	input   string
	pos     int
	readPos int
	ch      byte
}

func NewScanner(input string) *Scanner {
	l := &Scanner{input: input}
	l.readChar()
	return l
}

func (l *Scanner) NextToken() grammar.Token {
	var tok grammar.Token

	l.eatWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok.Type = grammar.EQ
			tok.Lit = string(ch) + string(l.ch)
		} else {
			tok.Type = grammar.ASSIGN
			tok.Lit = string(l.ch)
		}
	case '+':
		tok.Type = grammar.ADD
		tok.Lit = string(l.ch)
	case '"':
		tok.Type = grammar.STRING
		tok.Lit = l.readString()
	case ';':
		tok.Type = grammar.SEMICOLON
		tok.Lit = string(l.ch)
	case '(':
		tok.Type = grammar.LPAREN
		tok.Lit = string(l.ch)
	case ')':
		tok.Type = grammar.RPAREN
		tok.Lit = string(l.ch)
	case ',':
		tok.Type = grammar.COMMA
		tok.Lit = string(l.ch)
	case '{':
		tok.Type = grammar.LBRACE
		tok.Lit = string(l.ch)
	case '}':
		tok.Type = grammar.RBRACE
		tok.Lit = string(l.ch)
	case 0:
		tok.Type = grammar.EOF
		tok.Lit = ""
	default:
		if isLetter(l.ch) {
			tok.Lit = l.readIdentifier()
			tok.Type = grammar.Lookup(tok.Lit)
			return tok
		} else {
			tok.Type = grammar.ILLEGAL
			tok.Lit = string(l.ch)
		}
	}

	l.readChar()

	return tok
}

func (l *Scanner) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos += 1
}

func (l *Scanner) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func (l *Scanner) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Scanner) readIdentifier() string {
	pos := l.pos

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func (l *Scanner) readString() string {
	pos := l.pos + 1

	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[pos:l.pos]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}
