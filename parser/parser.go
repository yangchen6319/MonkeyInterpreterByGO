package parser

import (
	"MonkeyInterpreterByGO/ast"
	"MonkeyInterpreterByGO/lexer"
	"MonkeyInterpreterByGO/token"
	"fmt"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESS
	SUM
	PRODUCT
	PREFIX
	CALL
)

// 定义两种类型的表达式解析函数
type (
	prefixParseFn func() ast.ExpressionNode
	infixParseFn  func(node ast.ExpressionNode) ast.ExpressionNode
)

type Parser struct {
	l lexer.Lexer
	// 添加一个字段用于错误收集
	errors    []string
	curToken  token.Token
	nextToken token.Token
	// 定义表达式解析函数映射
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.parseNext()
	p.parseNext()
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseExpressionIdentifier)
	p.registerPrefix(token.INT, p.parseExpressionInt)
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseNext() {
	p.curToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

// 注册表达式解析函数映射
func (p *Parser) registerPrefix(tokType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokType] = fn
}

func (p *Parser) registerInfix(tokType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokType] = fn
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
	case token.RETURN:
		stmt := p.parseReturnStatement()
		if stmt != nil {
			return stmt
		} else {
			return nil
		}
	default:
		stmt := p.parseExpressionStatement()
		if stmt != nil {
			return stmt
		} else {
			return nil
		}
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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	reStmt := &ast.ReturnStatement{Token: p.curToken}
	p.parseNext()
	// 跳过对表达式的处理
	for !p.curTokenType(token.SEMICOLON) {
		p.parseNext()
	}
	return reStmt
}

// 解析表达式的函数
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	expStmt := &ast.ExpressionStatement{Token: p.curToken}
	expStmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenType(token.SEMICOLON) {
		p.parseNext()
	}
	return expStmt
}

func (p *Parser) parseExpression(priority int) ast.ExpressionNode {
	parseCall := p.prefixParseFns[p.curToken.Type]
	if parseCall == nil {
		return nil
	}
	return parseCall()
}

// 具体的表达式解析函数
func (p *Parser) parseExpressionIdentifier() ast.ExpressionNode {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseExpressionInt() ast.ExpressionNode {
	exp := &ast.Integer{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, err.Error())
		return nil
	}
	exp.Value = value
	return exp
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
