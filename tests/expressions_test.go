package tests

import (
	"testing"

	"github.com/mylang/interpreter/ast"
	"github.com/mylang/interpreter/interpreter"
	"github.com/stretchr/testify/assert"
)

// TestArithmeticOperations tests arithmetic operations
func TestArithmeticOperations(t *testing.T) {
	testCases := []struct {
		name     string
		expr     ast.Expression
		expected float64
	}{
		{
			"Addition",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "+",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			8,
		},
		{
			"Subtraction",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "-",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			2,
		},
		{
			"Multiplication",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "*",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			15,
		},
		{
			"Division",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(15)},
				Operator: "/",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			5,
		},
		{
			"Modulo",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(7)},
				Operator: "%",
				Right:    &ast.LiteralExpression{Value: float64(4)},
			},
			3,
		},
		{
			"Unary Minus",
			&ast.UnaryExpression{
				Operator: "-",
				Right:    &ast.LiteralExpression{Value: float64(5)},
			},
			-5,
		},
		{
			"Nested Operations",
			&ast.BinaryExpression{
				Left: &ast.BinaryExpression{
					Left:     &ast.LiteralExpression{Value: float64(2)},
					Operator: "+",
					Right:    &ast.LiteralExpression{Value: float64(3)},
				},
				Operator: "*",
				Right: &ast.BinaryExpression{
					Left:     &ast.LiteralExpression{Value: float64(4)},
					Operator: "-",
					Right:    &ast.LiteralExpression{Value: float64(1)},
				},
			},
			15, // (2 + 3) * (4 - 1) = 5 * 3 = 15
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			interp := interpreter.NewInterpreter()

			// Create a program that evaluates the expression
			program := &ast.Program{
				Statements: []ast.Statement{
					&ast.VariableDeclaration{
						Name:        "result",
						Initializer: tc.expr,
					},
					&ast.PrintStatement{
						Expression: &ast.VariableExpression{Name: "result"},
					},
				},
			}

			// Interpret the program
			output, err := interp.Interpret(program)

			// Check results
			assert.NoError(t, err)
			assert.Equal(t, 1, len(output))
			assert.Equal(t, tc.expected, interpreter.ParseNumber(output[0]))
		})
	}
}

// TestComparisonOperations tests comparison operations
func TestComparisonOperations(t *testing.T) {
	testCases := []struct {
		name     string
		expr     ast.Expression
		expected bool
	}{
		{
			"Equal (true)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "==",
				Right:    &ast.LiteralExpression{Value: float64(5)},
			},
			true,
		},
		{
			"Equal (false)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "==",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			false,
		},
		{
			"Not Equal (true)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "!=",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			true,
		},
		{
			"Not Equal (false)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "!=",
				Right:    &ast.LiteralExpression{Value: float64(5)},
			},
			false,
		},
		{
			"Less Than (true)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(3)},
				Operator: "<",
				Right:    &ast.LiteralExpression{Value: float64(5)},
			},
			true,
		},
		{
			"Less Than (false)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "<",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			false,
		},
		{
			"Greater Than (true)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: ">",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			true,
		},
		{
			"Greater Than (false)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(3)},
				Operator: ">",
				Right:    &ast.LiteralExpression{Value: float64(5)},
			},
			false,
		},
		{
			"Less Than or Equal (equal)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: "<=",
				Right:    &ast.LiteralExpression{Value: float64(5)},
			},
			true,
		},
		{
			"Less Than or Equal (less)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(3)},
				Operator: "<=",
				Right:    &ast.LiteralExpression{Value: float64(5)},
			},
			true,
		},
		{
			"Greater Than or Equal (equal)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: ">=",
				Right:    &ast.LiteralExpression{Value: float64(5)},
			},
			true,
		},
		{
			"Greater Than or Equal (greater)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: float64(5)},
				Operator: ">=",
				Right:    &ast.LiteralExpression{Value: float64(3)},
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			interp := interpreter.NewInterpreter()

			// Create a program that evaluates the expression
			program := &ast.Program{
				Statements: []ast.Statement{
					&ast.VariableDeclaration{
						Name:        "result",
						Initializer: tc.expr,
					},
					&ast.PrintStatement{
						Expression: &ast.VariableExpression{Name: "result"},
					},
				},
			}

			// Interpret the program
			output, err := interp.Interpret(program)

			// Check results
			assert.NoError(t, err)
			assert.Equal(t, 1, len(output))
			assert.Equal(t, tc.expected, interpreter.ParseBool(output[0]))
		})
	}
}

// TestLogicalOperations tests logical operations
func TestLogicalOperations(t *testing.T) {
	testCases := []struct {
		name     string
		expr     ast.Expression
		expected bool
	}{
		{
			"Logical AND (true && true)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: true},
				Operator: "&&",
				Right:    &ast.LiteralExpression{Value: true},
			},
			true,
		},
		{
			"Logical AND (true && false)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: true},
				Operator: "&&",
				Right:    &ast.LiteralExpression{Value: false},
			},
			false,
		},
		{
			"Logical OR (true || false)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: true},
				Operator: "||",
				Right:    &ast.LiteralExpression{Value: false},
			},
			true,
		},
		{
			"Logical OR (false || false)",
			&ast.BinaryExpression{
				Left:     &ast.LiteralExpression{Value: false},
				Operator: "||",
				Right:    &ast.LiteralExpression{Value: false},
			},
			false,
		},
		{
			"Logical NOT (!true)",
			&ast.UnaryExpression{
				Operator: "!",
				Right:    &ast.LiteralExpression{Value: true},
			},
			false,
		},
		{
			"Logical NOT (!false)",
			&ast.UnaryExpression{
				Operator: "!",
				Right:    &ast.LiteralExpression{Value: false},
			},
			true,
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
			true, // (true && false) || (!false) = false || true = true
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			interp := interpreter.NewInterpreter()

			// Create a program that evaluates the expression
			program := &ast.Program{
				Statements: []ast.Statement{
					&ast.VariableDeclaration{
						Name:        "result",
						Initializer: tc.expr,
					},
					&ast.PrintStatement{
						Expression: &ast.VariableExpression{Name: "result"},
					},
				},
			}

			// Interpret the program
			output, err := interp.Interpret(program)

			// Check results
			assert.NoError(t, err)
			assert.Equal(t, 1, len(output))
			assert.Equal(t, tc.expected, interpreter.ParseBool(output[0]))
		})
	}
}

// TestArrayExpressions tests array-related expressions
func TestArrayExpressions(t *testing.T) {
	// Test array literals and access
	t.Run("Array Literal and Access", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program with array literal and access
		program := &ast.Program{
			Statements: []ast.Statement{
				// Create an array
				&ast.VariableDeclaration{
					Name: "arr",
					Initializer: &ast.ArrayLiteralExpression{
						Elements: []ast.Expression{
							&ast.LiteralExpression{Value: float64(1)},
							&ast.LiteralExpression{Value: float64(2)},
							&ast.LiteralExpression{Value: float64(3)},
							&ast.LiteralExpression{Value: float64(4)},
							&ast.LiteralExpression{Value: float64(5)},
						},
					},
				},
				// Access array elements
				&ast.PrintStatement{
					Expression: &ast.ArrayAccessExpression{
						Array: &ast.VariableExpression{Name: "arr"},
						Index: &ast.LiteralExpression{Value: float64(0)},
					},
				},
				&ast.PrintStatement{
					Expression: &ast.ArrayAccessExpression{
						Array: &ast.VariableExpression{Name: "arr"},
						Index: &ast.LiteralExpression{Value: float64(2)},
					},
				},
				&ast.PrintStatement{
					Expression: &ast.ArrayAccessExpression{
						Array: &ast.VariableExpression{Name: "arr"},
						Index: &ast.LiteralExpression{Value: float64(4)},
					},
				},
			},
		}

		// Interpret the program
		output, err := interp.Interpret(program)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 3, len(output))
		assert.Equal(t, "1", output[0])
		assert.Equal(t, "3", output[1])
		assert.Equal(t, "5", output[2])
	})

	// Test array assignment
	t.Run("Array Assignment", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program with array assignment
		program := &ast.Program{
			Statements: []ast.Statement{
				// Create an array
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
				// Print initial value
				&ast.PrintStatement{
					Expression: &ast.ArrayAccessExpression{
						Array: &ast.VariableExpression{Name: "arr"},
						Index: &ast.LiteralExpression{Value: float64(1)},
					},
				},
				// Assign new value
				&ast.ArrayAssignmentStatement{
					Array: &ast.VariableExpression{Name: "arr"},
					Index: &ast.LiteralExpression{Value: float64(1)},
					Value: &ast.LiteralExpression{Value: float64(99)},
				},
				// Print new value
				&ast.PrintStatement{
					Expression: &ast.ArrayAccessExpression{
						Array: &ast.VariableExpression{Name: "arr"},
						Index: &ast.LiteralExpression{Value: float64(1)},
					},
				},
			},
		}

		// Interpret the program
		output, err := interp.Interpret(program)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 2, len(output))
		assert.Equal(t, "2", output[0])
		assert.Equal(t, "99", output[1])
	})
}

// Helper function for ParseNumber
func init() {
	// Add ParseNumber and ParseBool functions if they don't exist in the interpreter package
}
