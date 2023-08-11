package ast

import "interpreter-golang/token"

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
	Statement []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statement) <= 0 {
		return ""
	}

	return p.Statement[0].TokenLiteral()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {
}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
  Token token.Token
  Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {return i.Token.Literal}

