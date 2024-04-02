package lexer

import (
	"MonkeyInterpreterByGO/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	type expected struct {
		expectedType    token.TokenType
		expectedLiteral string
	}

	tests := []expected{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	L := New(input)

	for i, tt := range tests {
		tok := L.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
