package parser

import (
	"Fish-Lang/ast"
	"Fish-Lang/token"
)

// ParseStatement 中轴
func (p *Parser) ParseStatement() ast.Statement {
	// log.Printf("Parsing %q now!\n", p.curToken.Literal)

	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// 解析 let 语句
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	// 	 let		  x		  =		  5;
	// curToken , peekToken
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Indentifier{Token: p.curToken, Value: p.curToken.Literal}

	// let		x 		 = 		 5;
	//		curToken  peekToken
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: 跳过对表达式的处理
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// 解析 return 语句
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: 表达式求值
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// 解析表达式
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	defer untrace(trace("parseExpressionStatement"))

	stmt := &ast.ExpressionStatement{Token: p.curToken}

	// 最低的优先级先传递给 parseExpression
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
