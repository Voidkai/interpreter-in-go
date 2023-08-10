package lexer

import "interpreter-in-go/token"

type Lexer struct {
	input        string
	position     int
	readposition int
	ch           byte
}

var keyword = map[string]token.TokenType{
	"fn":     token.FUNCTION,
	"let":    token.LET,
	"return": token.RETURN,
	"if":     token.IF,
	"else":   token.ELSE,
	"true":   token.TRUE,
	"false":  token.FALSE,
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readposition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readposition]
	}

	l.position = l.readposition
	l.readposition++
}

func isLetter(input byte) bool {
	return 'a' <= input && input <= 'z' || 'A' <= input && input <= 'Z' || input == '_'
}

func isDigit(input byte) bool {
	return '0' <= input && input <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readposition >= len(l.input) {
		return 0
	}
	return l.input[l.readposition]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.New(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '(':
		tok = token.New(token.LPAREN, l.ch)
	case ')':
		tok = token.New(token.RPAREN, l.ch)
	case '{':
		tok = token.New(token.LBRACE, l.ch)
	case '}':
		tok = token.New(token.RBRACE, l.ch)
	case '[':
		tok = token.New(token.LBRACKET, l.ch)
	case ']':
		tok = token.New(token.RBRACKET, l.ch)
	case ',':
		tok = token.New(token.COMMA, l.ch)
	case ';':
		tok = token.New(token.SEMICOLON, l.ch)
	case '-':
		tok = token.New(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = token.New(token.BANG, l.ch)
		}
	case '/':
		tok = token.New(token.SLASH, l.ch)
	case '*':
		tok = token.New(token.ASTERISK, l.ch)
	case '<':
		tok = token.New(token.LT, l.ch)
	case '>':
		tok = token.New(token.GT, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.characterIdentifying(isLetter)
			tok.Type = lookupIndent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.characterIdentifying(isDigit)
			tok.Type = token.INT
			return tok
		} else {
			tok = token.New(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// characterIdentifying composed of readNumber and readIdentifier.
//
//	func (l *Lexer) readNumber() string {
//		position := l.position
//		for isDigit(l.ch) {
//			l.readChar()
//		}
//		return l.input[position:l.position]
//	}
//
//	func (l *Lexer) readIdentifier() string {
//		position := l.position
//		for isLetter(l.ch) {
//			l.readChar()
//		}
//		return l.input[position:l.position]
//	}
func (l *Lexer) characterIdentifying(f func(byte) bool) string {
	position := l.position
	for f(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Check if the identifier is a keyword.
func lookupIndent(ident string) token.TokenType {
	if tok, ok := keyword[ident]; ok {
		return tok
	}
	return token.IDENT
}

// Skip the white space and other characters that we don't care about.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}
