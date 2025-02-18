package repl

import (
	"bufio"
	"fmt"
	"io"

	"Interpreter_in_Go/lexer"
	"Interpreter_in_Go/token"
)

const PROMPT = ">>"

//goland:noinspection GoUnusedParameter
func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}
		line := scanner.Text()
		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%v\n", tok)
		}
	}
}
