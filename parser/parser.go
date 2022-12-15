package parser

import (
	"Fish-Lang/ast"
	"Fish-Lang/lexser"
	"Fish-Lang/token"
	"fmt"
)

type Parser struct {
	lexer  *lexser.Lexer // 词法分析器
	errors []string

	curToken  token.Token
	peekToken token.Token

	// map: <词法单元类型 -> 相应的解析函数>
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

// New 构造一个语法分析器
func New(l *lexser.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	// 建立 <此法单元 -> 解析函数> 映射
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIndentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupExpression)

	// 前缀运算符
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	// 中缀运算符
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)

	// 读取两个词法单元，设置 curToken 和 peekToken
	p.nextToken()
	p.nextToken()

	return p
}

// 获取下一个词法单元
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

//! ------------------------------------------

// ParseProgram 语法解析
func (p *Parser) ParseProgram() *ast.Program {
	// 构造 AST 根节点
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// 一次 Parse 解析一整句
	// for 循环 Parse 多句，每次返回一个 statement
	for p.curToken.Type != token.EOF {
		stmt := p.ParseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}
