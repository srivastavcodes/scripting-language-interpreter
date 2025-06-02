package parser

import (
	"fmt"
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
	lxr := lexer.NewLexer(input)
	psr := NewParser(lxr)

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

func checkParserErrors(t *testing.T, psr *Parser) {
	t.Helper()
	errors := psr.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, err := range errors {
		t.Errorf("parser error %q", err)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	t.Helper()
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral not 'let'. got=%q", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt not *ast.LetStatement. got=%T", stmt)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("stmt.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}

func TestReturnStatement(tst *testing.T) {
	input := `
return 5;
return 10;
return 992233;
`
	lxr := lexer.NewLexer(input)
	psr := NewParser(lxr)

	program := psr.ParseProgram()
	checkParserErrors(tst, psr)

	if len(program.Statements) != 3 {
		tst.Fatalf("program.Statements does not contain 3 statement. got=%d",
			len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			tst.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			tst.Errorf("returnStmt.TokenLiteral() not 'return'. got=%q",
				returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := `foobar;`

	lxr := lexer.NewLexer(input)
	psr := NewParser(lxr)
	program := psr.ParseProgram()
	checkParserErrors(t, psr)

	if len(program.Statements) != 1 {
		t.Fatalf("program does not have 1 statement. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", stmt)
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("Expression is not *ast.Identifier. got=%T", stmt.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not '%s'. got=%s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not '%s'. got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := `5;`

	lxr := lexer.NewLexer(input)
	psr := NewParser(lxr)
	program := psr.ParseProgram()
	checkParserErrors(t, psr)

	if len(program.Statements) != 1 {
		t.Fatalf("program does not have 1 length statement. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not *ast.ExpressionStatement. got=%T", stmt)
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("Expression is not *ast.IntegerLiteral. got=%T", literal)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not '%s'. got=%s", "5", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}
	for _, pt := range prefixTests {
		lxr := lexer.NewLexer(pt.input)
		psr := NewParser(lxr)
		program := psr.ParseProgram()
		checkParserErrors(t, psr)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statement. got=%d",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Program.Statements[0] is not ast.ExpressionStatement. got=%T", stmt)
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not ast.PrefixExpression. got=%T", exp)
		}
		if exp.Operator != pt.operator {
			t.Fatalf("exp.Operator not '%s'. got=%s", pt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, pt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integer, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integer.Value != value {
		t.Errorf("integer.Value not %d. got=%d", value, integer.Value)
		return false
	}
	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integer.TokenLiteral not %d. got=%s", value, integer.TokenLiteral())
		return false
	}
	return true
}
