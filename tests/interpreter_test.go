package tests

import (
	"testing"

	"github.com/mylang/interpreter/ast"
	"github.com/mylang/interpreter/interpreter"
	"github.com/stretchr/testify/assert"
)

// TestInterpreterExecution tests basic interpreter execution
func TestInterpreterExecution(t *testing.T) {
	interp := interpreter.NewInterpreter()

	// Create a simple program that prints "Hello, World!"
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.PrintStatement{
				Expression: &ast.LiteralExpression{
					Value: "Hello, World!",
				},
			},
		},
	}

	// Interpret the program
	output, err := interp.Interpret(program)

	// Check results
	assert.NoError(t, err)
	assert.Equal(t, 1, len(output))
	assert.Equal(t, "\"Hello, World!\"", output[0])
}

// TestVariableDeclaration tests variable declaration and access
func TestVariableDeclaration(t *testing.T) {
	interp := interpreter.NewInterpreter()

	// Create a program that declares a variable and prints it
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.VariableDeclaration{
				Name: "x",
				Initializer: &ast.LiteralExpression{
					Value: float64(42),
				},
			},
			&ast.PrintStatement{
				Expression: &ast.VariableExpression{
					Name: "x",
				},
			},
		},
	}

	// Interpret the program
	output, err := interp.Interpret(program)

	// Check results
	assert.NoError(t, err)
	assert.Equal(t, 1, len(output))
	assert.Equal(t, "42", output[0])
}

// TestArithmeticExpressions tests arithmetic expression evaluation
func TestArithmeticExpressions(t *testing.T) {
	testCases := []struct {
		name     string
		expr     ast.Expression
		expected string
	}{
		{
			"Addition",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "+",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			"8",
		},
		{
			"Subtraction",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "-",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			"2",
		},
		{
			"Multiplication",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "*",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			"15",
		},
		{
			"Division",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(6)},
				Operator: "/",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			"2",
		},
		{
			"Modulo",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(7)},
				Operator: "%",
				Right:    &ast.LiteralExpression{Value: float64(4)},
			},
			"3",
		},
		{
			"Complex Expression",
			&ast.BinaryExpression{
				Left: &ast.BinaryExpression{
					Left:     &ast.LiteralExpression{Value: float64(2)},
					Operator: "+",
					Right:    &ast.LiteralExpression{Value: float64(3)},
				},
				Operator: "*",
				Right:    &ast.LiteralExpression{Value: float64(4)},
			},
			"20",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			interp := interpreter.NewInterpreter()

			// Create a program that evaluates the expression
			program := &ast.Program{
				Statements: []ast.Statement{
					&ast.PrintStatement{
						Expression: tc.expr,
					},
				},
			}

			// Interpret the program
			output, err := interp.Interpret(program)

			// Check results
			assert.NoError(t, err)
			assert.Equal(t, 1, len(output))
			assert.Equal(t, tc.expected, output[0])
		})
	}
}

// TestComparisonExpressions tests comparison expression evaluation
func TestComparisonExpressions(t *testing.T) {
	testCases := []struct {
		name     string
		expr     ast.Expression
		expected string
	}{
		{
			"Equal",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "==",
				Right:    &ast.LiteralExpression{Value: float64(5)},
			},
			"true",
		},
		{
			"Not Equal",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "!=",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			"true",
		},
		{
			"Less Than",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(3)},
				Operator: "<",
				Right:    &ast.LiteralExpression{Value: float64(5)},
			},
			"true",
		},
		{
			"Greater Than",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: ">",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			"true",
		},
		{
			"Less Than or Equal",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "<=",
				Right:    &ast.LiteralExpression{Value: float64(5)},
			},
			"true",
		},
		{
			"Greater Than or Equal",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: ">=",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			"true",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			interp := interpreter.NewInterpreter()

			// Create a program that evaluates the expression
			program := &ast.Program{
				Statements: []ast.Statement{
					&ast.PrintStatement{
						Expression: tc.expr,
					},
				},
			}

			// Interpret the program
			output, err := interp.Interpret(program)

			// Check results
			assert.NoError(t, err)
			assert.Equal(t, 1, len(output))
			assert.Equal(t, tc.expected, output[0])
		})
	}
}

// TestLogicalExpressions tests logical expression evaluation
func TestLogicalExpressions(t *testing.T) {
	testCases := []struct {
		name     string
		expr     ast.Expression
		expected string
	}{
		{
			"Logical AND",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: true},
				Operator: "&&",
				Right:    &ast.LiteralExpression{Value: true},
			},
			"true",
		},
		{
			"Logical OR",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: false},
				Operator: "||",
				Right:    &ast.LiteralExpression{Value: true},
			},
			"true",
		},
		{
			"Logical NOT",
			&ast.UnaryExpression{
				Operator: "!",
				Right:    &ast.LiteralExpression{Value: false},
			},
			"true",
		},
		{
			"Complex Logical Expression",
			&ast.BinaryExpression{
				Left: &ast.BinaryExpression{
					Left:     &ast.LiteralExpression{Value: true},
					Operator: "&&",
					Right:    &ast.LiteralExpression{Value: false},
				},
				Operator: "||",
				Right: &ast.UnaryExpression{
					Operator: "!",
					Right:    &ast.LiteralExpression{Value: false},
				},
			},
			"true",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			interp := interpreter.NewInterpreter()

			// Create a program that evaluates the expression
			program := &ast.Program{
				Statements: []ast.Statement{
					&ast.PrintStatement{
						Expression: tc.expr,
					},
				},
			}

			// Interpret the program
			output, err := interp.Interpret(program)

			// Check results
			assert.NoError(t, err)
			assert.Equal(t, 1, len(output))
			assert.Equal(t, tc.expected, output[0])
		})
	}
}

// TestControlFlow tests control flow statements
func TestControlFlow(t *testing.T) {
	// Test if statement
	t.Run("If Statement", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program with an if statement
		program := &ast.Program{
			Statements: []ast.Statement{
				&ast.IfStatement{
					Condition: &ast.LiteralExpression{Value: true},
					ThenBranch: &ast.PrintStatement{
						Expression: &ast.LiteralExpression{Value: "Then"},
					},
					ElseBranch: &ast.PrintStatement{
						Expression: &ast.LiteralExpression{Value: "Else"},
					},
				},
			},
		}

		// Interpret the program
		output, err := interp.Interpret(program)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 1, len(output))
		assert.Equal(t, "\"Then\"", output[0])
	})

	// Test while loop
	t.Run("While Loop", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program with a while loop
		program := &ast.Program{
			Statements: []ast.Statement{
				&ast.VariableDeclaration{
					Name:        "counter",
					Initializer: &ast.LiteralExpression{Value: float64(0)},
				},
				&ast.WhileStatement{
					Condition: &ast.BinaryExpression{
						Left:     &ast.VariableExpression{Name: "counter"},
						Operator: "<",
						Right:    &ast.LiteralExpression{Value: float64(3)},
					},
					Body: &ast.Block{
						Statements: []ast.Statement{
							&ast.PrintStatement{
								Expression: &ast.VariableExpression{Name: "counter"},
							},
							&ast.AssignmentStatement{
								Name: "counter",
								Value: &ast.BinaryExpression{
									Left:     &ast.VariableExpression{Name: "counter"},
									Operator: "+",
									Right:    &ast.LiteralExpression{Value: float64(1)},
								},
							},
						},
					},
				},
			},
		}

		// Interpret the program
		output, err := interp.Interpret(program)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 3, len(output))
		assert.Equal(t, "0", output[0])
		assert.Equal(t, "1", output[1])
		assert.Equal(t, "2", output[2])
	})
}
