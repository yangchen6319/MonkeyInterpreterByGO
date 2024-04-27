package ast

import (
	"MonkeyInterpreterByGO/token"
	"testing"
)

func TestExpressionStatement(t *testing.T) {
	program := &Program{Statements: []StatementNode{
		&LetStatement{
			Token: token.Token{
				Type:    token.LET,
				Literal: "let",
			},
			Name: &Identifier{
				Token: token.Token{Type: token.IDENT, Literal: "x"},
				Value: "x",
			},
			Value: "20",
		},
	}}
	if program.String() != "let x = 20;" {
		t.Errorf("program.String() wrong! got %s\n", program.String())
	}
}
