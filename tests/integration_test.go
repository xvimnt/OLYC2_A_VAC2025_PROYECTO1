package tests

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/mylang/interpreter/ast"
	"github.com/mylang/interpreter/interpreter"
	"github.com/mylang/interpreter/lexer"
	"github.com/mylang/interpreter/parser"
	"github.com/stretchr/testify/assert"
)

// TestIntegrationWithExamplePrograms tests complete example programs to ensure the whole system works together
func TestIntegrationWithExamplePrograms(t *testing.T) {
	// Test a program with variable declarations, arithmetic, and conditionals
	t.Run("Basic Program", func(t *testing.T) {
		code := `
			var x = 10;
			var y = 5;
			var sum = x + y;
			var diff = x - y;
			print(sum);
			print(diff);
			if (sum > 10) {
				print("Sum is greater than 10");
			} else {
				print("Sum is not greater than 10");
			}
		`

		// Tokenize and parse the input
		ast, err := parseProgram(code)
		assert.NoError(t, err)

		// Run the interpreter
		interp := interpreter.NewInterpreter()
		output, err := interp.Interpret(ast)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 3, len(output))
		assert.Equal(t, "15", output[0]) // sum = 10 + 5 = 15
		assert.Equal(t, "5", output[1])  // diff = 10 - 5 = 5
		assert.Equal(t, "\"Sum is greater than 10\"", output[2])
	})

	// Test a program with loops
	t.Run("Loop Program", func(t *testing.T) {
		code := `
			var sum = 0;
			for (var i = 1; i <= 5; i = i + 1) {
				sum = sum + i;
			}
			print(sum);

			var factorial = 1;
			var j = 5;
			while (j > 0) {
				factorial = factorial * j;
				j = j - 1;
			}
			print(factorial);
		`

		// Tokenize and parse the input
		ast, err := parseProgram(code)
		assert.NoError(t, err)

		// Run the interpreter
		interp := interpreter.NewInterpreter()
		output, err := interp.Interpret(ast)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 2, len(output))
		assert.Equal(t, "15", output[0])  // sum = 1+2+3+4+5 = 15
		assert.Equal(t, "120", output[1]) // factorial = 5! = 120
	})

	// Test a program with functions
	t.Run("Function Program", func(t *testing.T) {
		code := `
			function add(a, b) {
				return a + b;
			}

			function factorial(n) {
				if (n <= 1) {
					return 1;
				}
				return n * factorial(n - 1);
			}

			var result1 = add(5, 3);
			print(result1);
			
			var result2 = factorial(5);
			print(result2);
		`

		// Tokenize and parse the input
		ast, err := parseProgram(code)
		assert.NoError(t, err)

		// Run the interpreter
		interp := interpreter.NewInterpreter()
		output, err := interp.Interpret(ast)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 2, len(output))
		assert.Equal(t, "8", output[0])   // add(5, 3) = 8
		assert.Equal(t, "120", output[1]) // factorial(5) = 120
	})

	// Test a program with arrays
	t.Run("Array Program", func(t *testing.T) {
		code := `
			var arr = [1, 2, 3, 4, 5];
			print(arr[0]);
			print(arr[4]);
			
			arr[2] = 99;
			print(arr[2]);
			
			var sum = 0;
			for (var i = 0; i < 5; i = i + 1) {
				sum = sum + arr[i];
			}
			print(sum);
		`

		// Tokenize and parse the input
		ast, err := parseProgram(code)
		assert.NoError(t, err)

		// Run the interpreter
		interp := interpreter.NewInterpreter()
		output, err := interp.Interpret(ast)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 4, len(output))
		assert.Equal(t, "1", output[0])
		assert.Equal(t, "5", output[1])
		assert.Equal(t, "99", output[2])
		assert.Equal(t, "111", output[3]) // sum = 1+2+99+4+5 = 111
	})
}

// Helper function to parse a program string into an AST
func parseProgram(code string) (*ast.Program, error) {
	// Create the input stream
	inputStream := antlr.NewInputStream(code)

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
	tree := parserInst.Program()

	// Check for errors
	errors := append(lexerInst.Errors, errorListener.Errors...)
	if len(errors) > 0 {
		return nil, &interpreter.RuntimeError{Message: errors[0]}
	}

	// Use the AST builder to convert the parse tree to an AST
	astBuilder := ast.NewASTBuilder()
	program := astBuilder.BuildAST(tree)

	return program, nil
}



// TestErrorHandling tests error handling capabilities of the interpreter
func TestErrorHandling(t *testing.T) {
	// Test runtime error - division by zero
	t.Run("Division by Zero", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program that divides by zero
		program := &ast.Program{
			Statements: []ast.Statement{
				&ast.VariableDeclaration{
					Name: "result",
					Initializer: &ast.BinaryExpression{
						Left:     &ast.LiteralExpression{Value: float64(10)},
						Operator: "/",
						Right:    &ast.LiteralExpression{Value: float64(0)},
					},
				},
			},
		}

		// Interpret the program - should fail with a division by zero error
		_, err := interp.Interpret(program)

		// Check that there was an error
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "division by zero") // Check the error message
	})

	// Test runtime error - undefined variable
	t.Run("Undefined Variable", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program that uses an undefined variable
		program := &ast.Program{
			Statements: []ast.Statement{
				&ast.PrintStatement{
					Expression: &ast.VariableExpression{Name: "undefinedVar"},
				},
			},
		}

		// Interpret the program - should fail with an undefined variable error
		_, err := interp.Interpret(program)

		// Check that there was an error
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "undefined variable") // Check the error message
	})

	// Test runtime error - invalid array index
	t.Run("Invalid Array Index", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program that uses an invalid array index
		program := &ast.Program{
			Statements: []ast.Statement{
				&ast.VariableDeclaration{
					Name: "arr",
					Initializer: &ast.ArrayLiteralExpression{
						Elements: []ast.Expression{
							&ast.LiteralExpression{Value: float64(1)},
							&ast.LiteralExpression{Value: float64(2)},
							&ast.LiteralExpression{Value: float64(3)},
						},
					},
				},
				&ast.PrintStatement{
					Expression: &ast.ArrayAccessExpression{
						Array: &ast.VariableExpression{Name: "arr"},
						Index: &ast.LiteralExpression{Value: float64(10)}, // Out of bounds
					},
				},
			},
		}

		// Interpret the program - should fail with an array index out of bounds error
		_, err := interp.Interpret(program)

		// Check that there was an error
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "index out of bounds") // Check the error message
	})
}
