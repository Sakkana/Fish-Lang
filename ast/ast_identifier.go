package ast

import "Fish-Lang/token"

// Indentifier 标识符实体
type Indentifier struct {
	Token token.Token // x -> token.IDENT 词法单元
	Value string
}

func (i *Indentifier) expressionNode()      {}
func (i *Indentifier) TokenLiteral() string { return i.Token.Literal }
func (i *Indentifier) String() string       { return i.Value }
