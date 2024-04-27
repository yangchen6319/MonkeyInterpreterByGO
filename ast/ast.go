package ast

import "MonkeyInterpreterByGO/token"

type Node interface {
	TokenLiteral() string
}

type StatementNode interface {
	Node
	statementNode()
}

type ExpressionNode interface {
	Node
	expressionNode()
}

// Program 根节点Node
type Program struct {
	Statements []StatementNode
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement Let语句结构体
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value string
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {
}

// return 语句结构体
type ReturnStatement struct {
	Token       token.Token
	ReturnValue ExpressionNode
}

func (r *ReturnStatement) statementNode() {

}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}
