package ast

import "interpreter-in-go/token"

type Node interface {
	TokenLiteral() string
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

func (ls *LetStatement) statementNode() {
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token //token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	//TODO implement me
	return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() {
	//TODO implement me
	panic("implement me")
}
