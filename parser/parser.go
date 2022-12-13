package parser

import (
	"Fish-Lang/ast"
	"Fish-Lang/lexser"
	"Fish-Lang/token"
	"fmt"
	"log"
)

type Parser struct {
	lexer  *lexser.Lexer // 词法分析器
	errors []string

	curToken  token.Token
	peekToken token.Token
}

// New 构造一个语法分析器
func New(l *lexser.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

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

func (p *Parser) ParseStatement() ast.Statement {
	log.Printf("Parsing %q now!\n", p.curToken.Literal)
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
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
