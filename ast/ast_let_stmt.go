package ast

import (
	"Fish-Lang/token"
	"bytes"
)

// LetStatement let 语句
type LetStatement struct {
	Token token.Token  // let -> token.LET 词法单元
	Name  *Indentifier // x -> 标识符
	Value Expression   // add(2, 3) - 1 -> 产生值的表达式
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
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
