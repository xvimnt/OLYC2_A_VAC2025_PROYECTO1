package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/mylang/interpreter/ast"
	"github.com/mylang/interpreter/interpreter"
	"github.com/mylang/interpreter/lexer"
	"github.com/mylang/interpreter/parser"
)

var (debug bool = false)

func main() {
	if len(os.Args) > 1 {
		// Check for debug flag
		if os.Args[1] == "--debug" {
			debug = true
			
			if len(os.Args) > 2 {
				// Run a file in debug mode
				runFile(os.Args[2])
			} else {
				// Start REPL in debug mode
				runPrompt()
			}
		} else {
			// Run a file normally
			runFile(os.Args[1])
		}
	} else {
		// Start REPL
		
		fmt.Println("MyLang Interpreter REPL")
		fmt.Println("Type 'exit' to quit")
		runPrompt()
	}
}

// runFile runs the interpreter on a file
func runFile(path string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Error resolving path: %v\n", err)
		os.Exit(1)
	}
	
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	
	source := string(bytes)
	run(source)
}

// runPrompt starts a REPL
func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	interp := interpreter.NewInterpreter()
	
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		
		line = strings.TrimSpace(line)
		if line == "exit" {
			break
		}
		
		if line == "" {
			continue
		}
		
		// Add a semicolon if missing for simple expressions/statements
		if !strings.HasSuffix(line, ";") && !strings.HasSuffix(line, "}") {
			line += ";"
		}
		
		output, err := runWithInterpreter(line, interp)
		if err != nil {
			fmt.Println(err)
		} else if len(output) > 0 {
			for _, line := range output {
				fmt.Println(line)
			}
		}
	}
}

// run executes the source code
func run(source string) {
	interp := interpreter.NewInterpreter()
	output, err := runWithInterpreter(source, interp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	for _, line := range output {
		fmt.Println(line)
	}
}

// runWithInterpreter executes the source code with the given interpreter
func runWithInterpreter(source string, interp *interpreter.Interpreter) ([]string, error) {
	// Create the input stream
	input := antlr.NewInputStream(source)
	
	// Create the lexer
	lexerInst := lexer.NewCustomLexer(input)
	
	// Create the token stream
	tokens := antlr.NewCommonTokenStream(lexerInst, antlr.TokenDefaultChannel)
	
	// Create the parser
	parserInst := parser.NewMyLangParser(tokens)
	
	// Set custom error handler
	errorListener := lexer.NewCustomErrorListener()
	parserInst.RemoveErrorListeners()
	parserInst.AddErrorListener(errorListener)
	
	// Parse the input
	program := parserInst.Program()
	
	// Check for syntax errors
	if len(lexerInst.Errors) > 0 || len(errorListener.Errors) > 0 {
		errors := append(lexerInst.Errors, errorListener.Errors...)
		return nil, fmt.Errorf("Syntax error: %s", strings.Join(errors, ", "))
	}
	
	// Debug output
	if debug {
		fmt.Println("Parse tree:")
		fmt.Println(program.ToStringTree(nil, parserInst))
	}
	
	// Build the AST from the parse tree
	astBuilder := ast.NewASTBuilder()
	astProgram := astBuilder.BuildAST(program)
	
	// Interpret the program
	output, err := interp.Interpret(astProgram)
	
	return output, err
}
