package repl

import (
	"Interpreter_in_Go/evaluator"
	"Interpreter_in_Go/object"
	"Interpreter_in_Go/parser"
	"bufio"
	"fmt"
	"io"

	"Interpreter_in_Go/lexer"
)

const PROMPT = ">>"

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		ok := scanner.Scan()
		if !ok {
			return
		}
		scanned := scanner.Text()
		lxr := lexer.NewLexer(scanned)
		psr := parser.NewParser(lxr)

		root := psr.ParseRootStatement()
		if len(psr.Errors()) != 0 {
			printParserErrors(output, psr.Errors())
			continue
		}
		evaluated := evaluator.Evaluate(root, env)
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
