package ast

import (
	"MonkeyInterpreterByGO/token"
	"bytes"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.Token.Literal + " ")
	out.WriteString(ls.Name.Value + " ")
	out.WriteString("= ")
	out.WriteString(ls.Value)
	out.WriteString(";")
	return out.String()
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

func (i *Identifier) String() string {
	var out bytes.Buffer
	out.WriteString(i.Value)
	return out.String()
}

type Integer struct {
	Token token.Token
	Value int64
}

func (i *Integer) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Integer) String() string {
	var out bytes.Buffer
	out.WriteString(i.Token.Literal)
	return out.String()
}
func (i *Integer) expressionNode() {
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

func (r *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(r.Token.Literal + " ")
	out.WriteString(r.ReturnValue.String())
	out.WriteString(";")
	return out.String()
}

// 表达式语句结构体
type ExpressionStatement struct {
	Token      token.Token
	Expression ExpressionNode
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	var out bytes.Buffer
	out.WriteString(es.Expression.String())
	return out.String()
}

func (es *ExpressionStatement) statementNode() {

}
