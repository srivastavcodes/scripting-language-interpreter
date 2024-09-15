# Interpreter-in-Go

This repository contains my implementation of an interpreter for a programming language as I follow the book *Writing an
Interpreter in Go* by Thorsten Ball. The project involves building a lexer, parser, and evaluator for a simple language.

## Project Overview

This project is an interpreter for the Monkey programming language. The implementation consists of several key components:

* **Lexical Analysis (Lexer)**: Breaks the input source code into tokens.
* **Parsing**: Converts the tokens into an Abstract Syntax Tree (AST).
* **Evaluation**: Interprets and executes the AST.

## Directory Structure
*subject to change*
```
├── ast/        # Abstract Syntax Tree implementation
├── evaluator/  # Code for evaluating the AST
├── lexer/      # Lexer to tokenize the source code
├── object/     # Definitions of Monkey language objects
├── parser/     # Parser to generate AST from tokens
├── repl/       # Read-Eval-Print Loop for interacting with the interpreter
├── token/      # Definitions of tokens
├── main.go     # Entry point for running the interpreter
└── README.md   # Project information and documentation
```

## How to Run the Interpreter

Ensure you have Go installed on your system. Then clone the repository and run the interpreter as follows:

```bash
git clone https://github.com/GitUser-3-2-3/Interpreter-in-Go.git
cd Interpreter-in-Go
go run main.go
```

This will start the REPL (Read-Eval-Print Loop), where you can enter Monkey code and see the interpreter's response.

## Example Usage

Here's an example of code written in the Monkey language:

```monkey
let factorial = fn(n) {
  if (n == 0) {
    return 1;
  } else {
    return n * factorial(n - 1);
  }
};

factorial(5); // Outputs 120
```

## Resources

* Book: *Writing an Interpreter in Go* by Thorsten Ball
* Go Documentation: https://golang.org/doc/

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.