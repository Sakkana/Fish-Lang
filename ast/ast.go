package ast

type Node interface {
	TokenLiteral() string
	String() string
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
