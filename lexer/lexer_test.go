package lexer

import (
	"Interpreter_in_Go/token"
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPARAN, "("},
		{token.RPARAN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	lex := New(input)
	var index = 0
	for i, test := range tests {
		tok := lex.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedLiteral, tok.Literal)
		}
		index++
		fmt.Printf("passed test times: %v\n", index)
	}
}
