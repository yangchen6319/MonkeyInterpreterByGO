package parser

import (
	"MonkeyInterpreterByGO/ast"
	"MonkeyInterpreterByGO/lexer"
	"MonkeyInterpreterByGO/token"
	"fmt"
)

type Parser struct {
	l lexer.Lexer
	// 添加一个字段用于错误收集
	errors    []string
	curToken  token.Token
	nextToken token.Token
}

func New(l lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.parseNext()
	p.parseNext()
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseNext() {
	p.curToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.StatementNode{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.parseNext()
	}
	return program
}

// 注意： 一个包含nil指针的接口不是nil接口
// 这里存在一个接口值的问题，接口值包含接口类型和类型对应的值两部分
// nil接口：类型为nil，值为nil，所以nil接口==nil
// 包含nil指针的接口：类型不为nil，但值为nil，包含nil指针的接口 != nil

// 这里对parseLetStatement()的返回值进行判断，如果是一个空指针，则返回一个nil接口
// 不能写成如下格式：
// return p.parseLetStatement()
// 如果parseLetStatement()返回一个空指针，会返回一个包含nil指针，类型为*LetStatement的值
func (p *Parser) parseStatement() ast.StatementNode {
	switch p.curToken.Type {
	case token.LET:
		stmt := p.parseLetStatement()
		if stmt != nil {
			return stmt
		} else {
			return nil
		}
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	letStmt := &ast.LetStatement{Token: p.curToken}
	// 指定当前Statement的token为let token
	if !p.peekExpect(token.IDENT) {
		return nil
	}
	letStmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	// TODO： “=”号后的部分
	if !p.peekExpect(token.ASSIGN) {
		return nil
	}
	for !p.curTokenType(token.SEMICOLON) {
		p.parseNext()
	}
	return letStmt
}

func (p *Parser) curTokenType(t token.TokenType) bool {
	return t == p.curToken.Type
}

func (p *Parser) peekTokenType(t token.TokenType) bool {
	return t == p.nextToken.Type
}

func (p *Parser) peekExpect(t token.TokenType) bool {
	if t == p.nextToken.Type {
		p.parseNext()
		return true
	} else {
		errorMsg := fmt.Sprintf("expected type: %s, got type: %s", t, p.nextToken.Type)
		p.errors = append(p.errors, errorMsg)
		return false
	}
}
