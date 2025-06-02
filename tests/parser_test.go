package tests

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/mylang/interpreter/lexer"
	"github.com/mylang/interpreter/parser"
	"github.com/stretchr/testify/assert"
)

// Test parsing valid basic statements
func TestParseBasicStatements(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Variable Declaration", "var x = 10;"},
		{"Assignment", "x = 20;"},
		{"Expression Statement", "1 + 2;"},
		{"Print Statement", "print(\"Hello\");"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := parseAndGetErrors(tc.input)
			assert.Empty(t, errors, "Expected no parse errors")
		})
	}
}

// Test parsing valid control flow statements
func TestParseControlFlow(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{
			"If Statement",
			"if (x > 10) { print(\"Greater\"); } else { print(\"Less or equal\"); }",
		},
		{
			"While Loop",
			"while (x < 10) { x = x + 1; }",
		},
		{
			"Do-While Loop",
			"do { x = x + 1; } while (x < 10);",
		},
		{
			"For Loop",
			"for (var i = 0; i < 10; i = i + 1) { print(i); }",
		},
		{
			"Switch Statement",
			"switch (x) { case 1: print(\"One\"); break; case 2: print(\"Two\"); break; default: print(\"Other\"); }",
		},
		{
			"Break Statement",
			"while (true) { break; }",
		},
		{
			"Continue Statement",
			"while (true) { continue; }",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := parseAndGetErrors(tc.input)
			assert.Empty(t, errors, "Expected no parse errors")
		})
	}
}

// Test parsing valid expressions
func TestParseExpressions(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Arithmetic", "var result = (1 + 2) * 3 / 4 % 5;"},
		{"Comparison", "var result = x > 10 && y < 20 || z == 30;"},
		{"Logical", "var result = !isTrue && (x > 0 || y < 0);"},
		{"Parenthesized", "var result = ((a + b) * (c - d)) / (e % f);"},
		{"Array Access", "var item = arr[0];"},
		{"Array Literal", "var arr = [1, 2, 3, 4, 5];"},
		{"Function Call", "var result = calculate(x, y, z);"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := parseAndGetErrors(tc.input)
			assert.Empty(t, errors, "Expected no parse errors")
		})
	}
}

// Test parsing valid function declarations
func TestParseFunctions(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{
			"No Parameters",
			"function greet() { print(\"Hello\"); }",
		},
		{
			"With Parameters",
			"function add(a, b) { return a + b; }",
		},
		{
			"Complex Function",
			`function fibonacci(n) {\n
				if (n <= 1) {\n
					return n;\n
				}\n
				return fibonacci(n-1) + fibonacci(n-2);\n
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := parseAndGetErrors(tc.input)
			assert.Empty(t, errors, "Expected no parse errors")
		})
	}
}

// Test invalid syntax cases
func TestInvalidSyntax(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Missing Semicolon", "var x = 10"},
		{"Invalid Variable Name", "var 1x = 10;"},
		{"Unbalanced Parentheses", "var result = (1 + 2;"},
		{"Unbalanced Braces", "if (x > 10) { print(\"Greater\"); "},
		{"Invalid Assignment Target", "1 = x;"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := parseAndGetErrors(tc.input)
			assert.NotEmpty(t, errors, "Expected parse errors")
		})
	}
}

// Helper function to parse input and get errors
func parseAndGetErrors(input string) []string {
	// Create the input stream
	inputStream := antlr.NewInputStream(input)

	// Create the lexer
	lexerInst := lexer.NewCustomLexer(inputStream)

	// Create the token stream
	tokens := antlr.NewCommonTokenStream(lexerInst, antlr.TokenDefaultChannel)

	// Create the parser
	parserInst := parser.NewMyLangParser(tokens)

	// Set custom error handler
	errorListener := lexer.NewCustomErrorListener()
	parserInst.RemoveErrorListeners()
	parserInst.AddErrorListener(errorListener)

	// Parse the input
	parserInst.Program()

	// Return errors
	return append(lexerInst.Errors, errorListener.Errors...)
}
