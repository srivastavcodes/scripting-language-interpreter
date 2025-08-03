package repl

import (
	"Interpreter_in_Go/evaluator"
	"Interpreter_in_Go/parser"
	"bufio"
	"fmt"
	"io"

	"Interpreter_in_Go/lexer"
)

const PROMPT = ">>"

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		lex := lexer.NewLexer(line)
		psr := parser.NewParser(lex)

		program := psr.ParseProgram()
		if len(psr.Errors()) != 0 {
			printParserErrors(output, psr.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			_, _ = io.WriteString(output, evaluated.Inspect())
			_, _ = io.WriteString(output, "\n")
		}
	}
}

func printParserErrors(output io.Writer, errors []string) {
	_, _ = io.WriteString(output, "Parser Errors:\n")

	for _, err := range errors {
		_, _ = io.WriteString(output, "\t"+err+"\n")
	}
}
