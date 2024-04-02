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
	switch L.ch {
	case '=':
		tok = newToken(token.ASSIGN, L.ch)
	case '+':
		tok = newToken(token.PLUS, L.ch)
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
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
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
