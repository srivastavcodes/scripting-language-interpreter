package parser

import (
	"fmt"
	"strconv"

	"Interpreter_in_Go/ast"
	"Interpreter_in_Go/lexer"
	"Interpreter_in_Go/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -x or !x
	CALL        // myFunc(x)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lxr    *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func NewParser(lxr *lexer.Lexer) *Parser {
	psr := &Parser{lxr: lxr, errors: []string{}}

	// Read two tokens, so that curToken and peekToken are set
	psr.nextToken()
	psr.nextToken()

	psr.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	psr.registerPrefix(token.IDENT, psr.parseIdentifier)

	psr.registerPrefix(token.INT, psr.parseIntegerLiteral)
	psr.registerPrefix(token.BANG, psr.parsePrefixExpression)

	psr.registerPrefix(token.MINUS, psr.parsePrefixExpression)
	return psr
}

func (psr *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !psr.currentTokenIs(token.EOF) {
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
	case token.RETURN:
		return psr.parseReturnStatement()
	default:
		return psr.parseExpressionStatement()
	}
}

func (psr *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: psr.curToken}
	if !psr.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{
		Token: psr.curToken, Value: psr.curToken.Literal}

	if !psr.expectPeek(token.ASSIGN) {
		return nil
	}
	/*
		todo: add Expression handling,
		we are skipping the expressions until we encounter a semicolon
	*/
	for !psr.currentTokenIs(token.SEMICOLON) {
		psr.nextToken()
	}
	return stmt
}

func (psr *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: psr.curToken}
	psr.nextToken()
	/*
		todo: add Expression handling,
		we are skipping the expression until we encounter a semicolon
	*/
	for !psr.currentTokenIs(token.SEMICOLON) {
		psr.nextToken()
	}
	return stmt
}

func (psr *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: psr.curToken}
	stmt.Expression = psr.parseExpression(LOWEST)

	if psr.peekTokenIs(token.SEMICOLON) {
		psr.nextToken()
	}
	return stmt
}

func (psr *Parser) parseExpression(precedence int) ast.Expression {
	prefix := psr.prefixParseFns[psr.curToken.Type]
	if nil == prefix {
		psr.noPrefixParseFnError(psr.curToken.Type)
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (psr *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: psr.curToken, Value: psr.curToken.Literal}
}

func (psr *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: psr.curToken}

	value, err := strconv.ParseInt(psr.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", psr.curToken.Literal)
		psr.errors = append(psr.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (psr *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    psr.curToken,
		Operator: psr.curToken.Literal,
	}
	psr.nextToken()
	expression.Right = psr.parseExpression(PREFIX)
	return expression
}

func (psr *Parser) Errors() []string {
	return psr.errors
}

func (psr *Parser) peekError(tokn token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		tokn, psr.peekToken.Type)
	psr.errors = append(psr.errors, msg)
}

func (psr *Parser) noPrefixParseFnError(tokn token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", tokn)
	psr.errors = append(psr.errors, msg)
}

func (psr *Parser) nextToken() {
	psr.curToken = psr.peekToken
	psr.peekToken = psr.lxr.NextToken()
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
		psr.peekError(tokn)
		return false
	}
}

func (psr *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	psr.prefixParseFns[tokenType] = fn
}

func (psr *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	psr.infixParseFns[tokenType] = fn
}
