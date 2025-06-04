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
	// Default to standard mode
	languageMode := "standard"
	
	if len(os.Args) > 1 {
		// Check for language mode flag
		if os.Args[1] == "--v" || os.Args[1] == "-v" {
			languageMode = "v"
			os.Args = append(os.Args[:1], os.Args[2:]...) // Remove the flag
		}
		
		// Check for debug flag
		if len(os.Args) > 1 && os.Args[1] == "--debug" {
			debug = true
			
			if len(os.Args) > 2 {
				// Run a file in debug mode
				runFile(os.Args[2], languageMode)
			} else {
				// Start REPL in debug mode
				runPrompt(languageMode)
			}
		} else if len(os.Args) > 1 {
			// Run a file normally
			runFile(os.Args[1], languageMode)
		} else {
			// No arguments after flags, start REPL
			printWelcome(languageMode)
			runPrompt(languageMode)
		}
	} else {
		// No arguments at all, start standard REPL
		printWelcome(languageMode)
		runPrompt(languageMode)
	}
}

// printWelcome displays a welcome message based on the language mode
func printWelcome(mode string) {
	if mode == "v" {
		fmt.Println("V Language Interpreter REPL")
	} else {
		fmt.Println("MyLang Interpreter REPL")
	}
	fmt.Println("Type 'exit' to quit")
}

// runFile runs the interpreter on a file
func runFile(path string, mode string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Error resolving path: %v\n", err)
		os.Exit(1)
	}

	// Read the file
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Run the source
	run(string(bytes), mode)
}

// runPrompt starts a REPL
func runPrompt(mode string) {
	reader := bufio.NewReader(os.Stdin)
	
	// Create the appropriate interpreter based on mode
	var standardInterp *interpreter.Interpreter
	var vInterp *interpreter.VInterpreter
	
	if mode == "v" {
		vInterp = interpreter.NewVInterpreter()
	} else {
		standardInterp = interpreter.NewInterpreter()
	}
	
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
		
		// Add a semicolon if missing for simple expressions/statements in standard mode
		if mode != "v" && !strings.HasSuffix(line, ";") && !strings.HasSuffix(line, "}") {
			line += ";"
		}
		
		var output []string
		var err error
		
		if mode == "v" {
			output, err = runWithVInterpreter(line, vInterp)
		} else {
			output, err = runWithInterpreter(line, standardInterp)
		}
		
		if err != nil {
			fmt.Println(err)
		} else if len(output) > 0 {
			for _, line := range output {
				fmt.Println(line)
			}
		}
	}
}

// run executes the source code based on language mode
func run(source string, mode string) {
	var output []string
	var err error
	
	if mode == "v" {
		vInterp := interpreter.NewVInterpreter()
		output, err = runWithVInterpreter(source, vInterp)
	} else {
		interp := interpreter.NewInterpreter()
		output, err = runWithInterpreter(source, interp)
	}
	
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

// runWithVInterpreter executes the V language source code with the given V interpreter
func runWithVInterpreter(source string, interp *interpreter.VInterpreter) ([]string, error) {
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
	
	// Interpret the program with the V interpreter
	output, err := interp.InterpretV(astProgram)
	
	return output, err
}
