package ast

import (
	"bytes"
	"interpreter-in-go/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Identifier is a structure to formalize the identifier in a let assignment syntax
type Identifier struct {
	Token token.Token //token.IDENT
	Value string
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// LetStatement is a structure for formalizing the let assignment syntax
type LetStatement struct {
	Token token.Token //token.LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()

}

func (ls *LetStatement) statementNode() {
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token //token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) TokenLiteral() string {
	//TODO implement me
	return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() {
	//TODO implement me
	panic("implement me")
}

type ExpressionStatement struct {
	Token      token.Token //the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	var out bytes.Buffer
	if es.Expression != nil {
		out.WriteString(es.Expression.String())
		return out.String()
	}

	return ""
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) statementNode() {
}

type IntegerLiteral struct {
	Token token.Token //token.INT
	Value int64
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) expressionNode() {
}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token //the prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

func (pe *PrefixExpression) expressionNode() {
}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

type InfixExpression struct {
	Token    token.Token //the operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	return "(" + ie.Left.String() + " " + ie.Operator + " " + ie.Right.String() + ")"
}

func (ie *InfixExpression) expressionNode() {

}
