package ast

import "Fish-Lang/token"

type Node interface {
	TokenLiteral() string
}

// Statement 语句 -> let x = 5;
// 不会产生值
type Statement interface {
	Node
	statementNode()
}

// Expression 表达式 -> 5, add(2, 3)
// 会产生值
type Expression interface {
	Node
	expressionNode()
}

// Program 每个 AST 的根节点
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

// LetStatement let 语句
type LetStatement struct {
	Token token.Token  // let -> token.LET 词法单元
	Name  *Indentifier // x -> 标识符
	Value Expression   // add(2, 3) - 1 -> 产生值的表达式
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Indentifier 标识符实体
type Indentifier struct {
	Token token.Token // x -> token.IDENT 词法单元
	Value string
}

func (i *Indentifier) expressionNode()      {}
func (i *Indentifier) TokenLiteral() string { return i.Token.Literal }
