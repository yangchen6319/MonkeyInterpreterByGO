package parser

import (
	"MonkeyInterpreterByGO/ast"
	"MonkeyInterpreterByGO/lexer"
	"MonkeyInterpreterByGO/token"
)

type Parser struct {
	l lexer.Lexer

	curToken  token.Token
	nextToken token.Token
}

func New(l lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.parseNext()
	p.parseNext()
	return p
}

func (p *Parser) parseNext() {
	p.curToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
