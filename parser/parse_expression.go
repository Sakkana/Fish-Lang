package parser

import (
	"Fish-Lang/ast"
	"Fish-Lang/token"
	"fmt"
	"strconv"
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
	defer untrace(trace("parseExpression"))

	// 寻找当前词法单元对应解析函数
	// 整数字面量，!，-
	prefix := p.prefixParseFns[p.curToken.Type]
	// 不存在该词法单元对应的解析函数
	if prefix == nil {
		p.noPrefixParseError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	// log.Println("now is " + p.curToken.Literal)

	// 找优先级更高的赶紧计算右侧
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

// 表达式解析函数
type (
	// [out]: 该表达式的 ast
	prefixParseFn func() ast.Expression

	// [in]: 中缀运算符左侧的内容
	// [out]: 该表达式的 ast
	infixParseFn func(ast.Expression) ast.Expression
)

// 注册哈希表
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
	// log.Printf("register %v succeed\n", tokenType)
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// 标识符解析函数
func (p *Parser) parseIndentifier() ast.Expression {
	defer untrace(trace("parseIndentifier"))

	return &ast.Indentifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

// 整型字面量解析函数
func (p *Parser) parseIntegerLiteral() ast.Expression {
	defer untrace(trace("parseIntegerExpression"))

	ilt := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer",
			p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	ilt.Value = value
	return ilt
}

// 逻辑 true, false 解析函数
func (p *Parser) parseBoolean() ast.Expression {
	defer untrace(trace("parseBooleanExpression"))
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}

// 左括号要更改优先级，解析函数
func (p *Parser) parseGroupExpression() ast.Expression {
	defer untrace(trace("parseGroupExpression"))

	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

// 前缀表达式解析函数
func (p *Parser) parsePrefixExpression() ast.Expression {
	defer untrace(trace("parsePrefixExpression"))

	pfx := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	// 取下一个操作数
	p.nextToken()
	// 传入当前运算符的优先级
	pfx.Right = p.parseExpression(PREFIX)

	return pfx
}

func (p *Parser) noPrefixParseError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

// 中缀表达式解析函数
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	defer untrace(trace("parseInfixExpression"))

	ifx := &ast.InfixExpression{
		Token:    token.Token{},
		Left:     left,
		Operator: p.curToken.Literal,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	ifx.Right = p.parseExpression(precedence)

	// 右结合
	//if ifx.Operator == "+" {
	//	ifx.Right = p.parseExpression(precedence - 1)
	//} else {
	//	ifx.Right = p.parseExpression(precedence)
	//}

	return ifx
}
