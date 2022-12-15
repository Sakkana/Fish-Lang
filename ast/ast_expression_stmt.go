package ast

import (
	"Fish-Lang/token"
)

// ExpressionStatement 表达式
type ExpressionStatement struct {
	Token      token.Token // 该表达式中的第一个语法单元
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
