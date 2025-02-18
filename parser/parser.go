package parser

import (
	"Interpreter_in_Go/ast"
	"Interpreter_in_Go/lexer"
	"Interpreter_in_Go/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	psr := &Parser{l: l}

	// Read two tokens, so that curToken and peekToken are set
	psr.nextToken()
	psr.nextToken()

	return psr
}

func (psr *Parser) nextToken() {
	psr.curToken = psr.peekToken
	psr.peekToken = psr.l.NextToken()
}

func (psr *Parser) ParseProgram() *ast.Program {
	return nil
}
