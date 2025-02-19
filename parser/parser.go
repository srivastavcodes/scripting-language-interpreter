package parser

import (
	"Interpreter_in_Go/ast"
	"Interpreter_in_Go/lexer"
	"Interpreter_in_Go/token"
)

type Parser struct {
	lxr *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(lxr *lexer.Lexer) *Parser {
	psr := &Parser{lxr: lxr}

	// Read two tokens, so that curToken and peekToken are set
	psr.nextToken()
	psr.nextToken()

	return psr
}

func (psr *Parser) nextToken() {
	psr.curToken = psr.peekToken
	psr.peekToken = psr.lxr.NextToken()
}

func (psr *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for psr.curToken.Type != token.EOF {
		stmt := psr.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		psr.nextToken()
	}
	return program
}

func (psr *Parser) parseStatement() ast.Statement {
	switch psr.curToken.Type {
	case token.LET:
		return psr.parseLetStatement()
	default:
		return nil
	}
}

func (psr *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: psr.curToken}
	if !psr.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: psr.curToken, Value: psr.curToken.Literal}
	if !psr.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO -> we are skipping the expressions until we encounter a semicolon
	for !psr.currentTokenIs(token.SEMICOLON) {
		psr.nextToken()
	}
	return stmt
}

func (psr *Parser) currentTokenIs(tokn token.TokenType) bool {
	return psr.curToken.Type == tokn
}

func (psr *Parser) peekTokenIs(tokn token.TokenType) bool {
	return psr.peekToken.Type == tokn
}

func (psr *Parser) expectPeek(tokn token.TokenType) bool {
	if psr.peekTokenIs(tokn) {
		psr.nextToken()
		return true
	} else {
		return false
	}
}
