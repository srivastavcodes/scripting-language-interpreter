package parser

import (
	"testing"

	"Interpreter_in_Go/ast"
	"Interpreter_in_Go/lexer"
)

func TestLetStatements(tst *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	lxr := lexer.New(input)
	psr := New(lxr)

	program := psr.ParseProgram()
	checkParserErrors(tst, psr)

	if program == nil {
		tst.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		tst.Fatalf("program.Statements does not contain 3 statements. go=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, t := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(tst, stmt, t.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(tst *testing.T, psr *Parser) {
	errors := psr.Errors()
	if len(errors) == 0 {
		return
	}
	tst.Errorf("parser has %d errors", len(errors))
	for _, err := range errors {
		tst.Errorf("parser error %q", err)
	}
	tst.FailNow()
}

func testLetStatement(tst *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		tst.Errorf("stmt.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		tst.Errorf("stmt not *ast.LetStatement. got=%T", stmt)
		return false
	}
	if letStmt.Name.Value != name {
		tst.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		tst.Errorf("stmt.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}
