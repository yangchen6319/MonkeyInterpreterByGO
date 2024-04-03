package lexer

import "MonkeyInterpreterByGO/token"

type Lexer struct {
	input        string
	position     int
	nextPosition int
	ch           byte
}

func New(input string) *Lexer {
	L := &Lexer{
		input: input,
	}
	L.readChar()
	return L
}

func (L *Lexer) readChar() {
	if L.nextPosition >= len(L.input) {
		L.ch = 0
	} else {
		L.ch = L.input[L.nextPosition]
	}
	L.position = L.nextPosition
	L.nextPosition += 1
}

func (L *Lexer) NextToken() token.Token {
	var tok token.Token
	L.eatWhitespace()
	switch L.ch {
	case '=':
		if L.peekChar() == '=' {
			L.readChar()
			tok.Type = token.EQ
			tok.Literal = "=="
		} else {
			tok = newToken(token.ASSIGN, L.ch)
		}
	case '!':
		if L.peekChar() == '=' {
			L.readChar()
			tok.Type = token.NOT_EQ
			tok.Literal = "!="
		} else {
			tok = newToken(token.BANG, L.ch)
		}
	case '+':
		tok = newToken(token.PLUS, L.ch)
	case '-':
		tok = newToken(token.MINUS, L.ch)
	case '/':
		tok = newToken(token.SLASH, L.ch)
	case '*':
		tok = newToken(token.ASTERISK, L.ch)
	case ',':
		tok = newToken(token.COMMA, L.ch)
	case ';':
		tok = newToken(token.SEMICOLON, L.ch)
	case '(':
		tok = newToken(token.LPAREN, L.ch)
	case ')':
		tok = newToken(token.RPAREN, L.ch)
	case '{':
		tok = newToken(token.LBRACE, L.ch)
	case '}':
		tok = newToken(token.RBRACE, L.ch)
	case '<':
		tok = newToken(token.LT, L.ch)
	case '>':
		tok = newToken(token.GT, L.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(L.ch) {
			tok.Literal = L.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isNumber(L.ch) {
			tok.Literal = L.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			return newToken(token.ILLEGAL, L.ch)
		}
	}
	L.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (L *Lexer) readIdentifier() string {
	position := L.position
	for isLetter(L.ch) {
		L.readChar()
	}
	return L.input[position:L.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (L *Lexer) readNumber() string {
	position := L.position
	for isNumber(L.ch) {
		L.readChar()
	}
	return L.input[position:L.position]
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// 跳过空格
func (L *Lexer) eatWhitespace() {
	for L.ch == ' ' || L.ch == '\n' || L.ch == '\r' || L.ch == '\t' {
		L.readChar()
	}
}

// 查看下一字符
func (L *Lexer) peekChar() byte {
	if L.nextPosition >= len(L.input) {
		return 0
	} else {
		return L.input[L.nextPosition]
	}
}
