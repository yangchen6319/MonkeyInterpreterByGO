package lexer

import (
	"MonkeyInterpreterByGO/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "let five = 5;" +
		"let ten = 10;" +
		"let add = fn(x, y){" +
		"x + y;" +
		"};" +
		"" +
		"let result = add(five, ten);" +
		"- / * < >;" +
		"true; false; if; else; return;" +
		"== ! !=;"

	type expected struct {
		expectedType    token.TokenType
		expectedLiteral string
	}

	tests := []expected{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.SEMICOLON, ";"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.SEMICOLON, ";"},
		{token.ELSE, "else"},
		{token.SEMICOLON, ";"},
		{token.RETURN, "return"},
		{token.SEMICOLON, ";"},
		{token.EQ, "=="},
		{token.BANG, "!"},
		{token.NOT_EQ, "!="},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	L := New(input)

	for i, tt := range tests {
		tok := L.NextToken()
		if tok.Type != tt.expectedType {
			println(tt.expectedLiteral, tok.Literal)
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
