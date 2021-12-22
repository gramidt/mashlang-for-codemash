package grammar

import "strconv"

type TokenType int

type Token struct {
	Type TokenType
	Lit  string
}

// The tokens of the "Mash" programming language.
const (
	ILLEGAL TokenType = iota
	EOF

	// Literals (identifiers and basic types)
	IDENT
	STRING

	// Operators
	ASSIGN
	ADD
	EQ

	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUN
	VAR
	TRUE
	FALSE
	IF
	ELSE
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	IDENT:  "IDENT",
	STRING: "STRING",

	ASSIGN: "=",
	ADD:    "+",
	EQ:     "==",

	COMMA:     ",",
	SEMICOLON: ";",
	LPAREN:    "(",
	RPAREN:    ")",
	LBRACE:    "{",
	RBRACE:    "}",

	FUN:   "fun",
	VAR:   "var",
	TRUE:  "true",
	FALSE: "false",
	IF:    "if",
	ELSE:  "else",
}

func (tt TokenType) String() string {
	s := ""
	if 0 <= tt && tt < TokenType(len(tokens)) {
		s = tokens[tt]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tt)) + ")"
	}
	return s
}

const (
	LowestPrecedence = 1
)

func (tok Token) Precedence() int {
	switch tok.Type {
	case EQ:
		return 2
	case ADD:
		return 3
	case LPAREN:
		return 4
	}
	return LowestPrecedence
}

var keywords = map[string]TokenType{
	tokens[FUN]:   FUN,
	tokens[VAR]:   VAR,
	tokens[TRUE]:  TRUE,
	tokens[FALSE]: FALSE,
	tokens[IF]:    IF,
	tokens[ELSE]:  ELSE,
}

func Lookup(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
