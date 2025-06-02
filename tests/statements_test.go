package tests

import (
	"testing"

	"github.com/mylang/interpreter/ast"
	"github.com/mylang/interpreter/interpreter"
	"github.com/stretchr/testify/assert"
)

// TestVariableStatements tests variable-related statements
func TestVariableStatements(t *testing.T) {
	// Test variable declaration and assignment
	t.Run("Variable Declaration and Assignment", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program with variable declaration and assignment
		program := &ast.Program{
			Statements: []ast.Statement{
				// Declare variable x
				&ast.VariableDeclaration{
					Name:        "x",
					Initializer: &ast.LiteralExpression{Value: float64(10)},
				},
				// Print x
				&ast.PrintStatement{
					Expression: &ast.VariableExpression{Name: "x"},
				},
				// Assign new value to x
				&ast.AssignmentStatement{
					Name:  "x",
					Value: &ast.LiteralExpression{Value: float64(20)},
				},
				// Print x again
				&ast.PrintStatement{
					Expression: &ast.VariableExpression{Name: "x"},
				},
			},
		}

		// Interpret the program
		output, err := interp.Interpret(program)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 2, len(output))
		assert.Equal(t, "10", output[0])
		assert.Equal(t, "20", output[1])
	})

	// Test variable scope
	t.Run("Variable Scope", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program with nested scopes
		program := &ast.Program{
			Statements: []ast.Statement{
				// Declare variable x in global scope
				&ast.VariableDeclaration{
					Name:        "x",
					Initializer: &ast.LiteralExpression{Value: float64(10)},
				},
				// Print x
				&ast.PrintStatement{
					Expression: &ast.VariableExpression{Name: "x"},
				},
				// Create a block with local variable x
				&ast.Block{
					Statements: []ast.Statement{
						&ast.VariableDeclaration{
							Name:        "x",
							Initializer: &ast.LiteralExpression{Value: float64(20)},
						},
						&ast.PrintStatement{
							Expression: &ast.VariableExpression{Name: "x"},
						},
					},
				},
				// Print global x again
				&ast.PrintStatement{
					Expression: &ast.VariableExpression{Name: "x"},
				},
			},
		}

		// Interpret the program
		output, err := interp.Interpret(program)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 3, len(output))
		assert.Equal(t, "10", output[0])
		assert.Equal(t, "20", output[1])
		assert.Equal(t, "10", output[2])
	})
}

// TestControlFlowStatements tests control flow statements
func TestControlFlowStatements(t *testing.T) {
	// Test if-else statement
	t.Run("If-Else Statement", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program with if-else statements
		program := &ast.Program{
			Statements: []ast.Statement{
				// Declare variable x
				&ast.VariableDeclaration{
					Name:        "x",
					Initializer: &ast.LiteralExpression{Value: float64(5)},
				},
				// If x > 10, print "Greater", else print "Less or equal"
				&ast.IfStatement{
					Condition: &ast.BinaryExpression{
						Left:     &ast.VariableExpression{Name: "x"},
						Operator: ">",
						Right:    &ast.LiteralExpression{Value: float64(10)},
					},
					ThenBranch: &ast.PrintStatement{
						Expression: &ast.LiteralExpression{Value: "Greater"},
					},
					ElseBranch: &ast.PrintStatement{
						Expression: &ast.LiteralExpression{Value: "Less or equal"},
					},
				},
				// Assign new value to x
				&ast.AssignmentStatement{
					Name:  "x",
					Value: &ast.LiteralExpression{Value: float64(15)},
				},
				// Test if statement again
				&ast.IfStatement{
					Condition: &ast.BinaryExpression{
						Left:     &ast.VariableExpression{Name: "x"},
						Operator: ">",
						Right:    &ast.LiteralExpression{Value: float64(10)},
					},
					ThenBranch: &ast.PrintStatement{
						Expression: &ast.LiteralExpression{Value: "Greater"},
					},
					ElseBranch: &ast.PrintStatement{
						Expression: &ast.LiteralExpression{Value: "Less or equal"},
					},
				},
			},
		}

		// Interpret the program
		output, err := interp.Interpret(program)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 2, len(output))
		assert.Equal(t, "\"Less or equal\"", output[0])
		assert.Equal(t, "\"Greater\"", output[1])
	})

	// Test for loop
	t.Run("For Loop", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program with a for loop
		program := &ast.Program{
			Statements: []ast.Statement{
				&ast.ForStatement{
					Initializer: &ast.VariableDeclaration{
						Name:        "i",
						Initializer: &ast.LiteralExpression{Value: float64(0)},
					},
					Condition: &ast.BinaryExpression{
						Left:     &ast.VariableExpression{Name: "i"},
						Operator: "<",
						Right:    &ast.LiteralExpression{Value: float64(3)},
					},
					Increment: &ast.BinaryExpression{
						Left:     &ast.VariableExpression{Name: "i"},
						Operator: "+",
						Right:    &ast.LiteralExpression{Value: float64(1)},
					},
					Body: &ast.PrintStatement{
						Expression: &ast.VariableExpression{Name: "i"},
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

	// Test break statement
	t.Run("Break Statement", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program with a while loop and break statement
		program := &ast.Program{
			Statements: []ast.Statement{
				&ast.VariableDeclaration{
					Name:        "i",
					Initializer: &ast.LiteralExpression{Value: float64(0)},
				},
				&ast.WhileStatement{
					Condition: &ast.LiteralExpression{Value: true},
					Body: &ast.Block{
						Statements: []ast.Statement{
							&ast.PrintStatement{
								Expression: &ast.VariableExpression{Name: "i"},
							},
							&ast.AssignmentStatement{
								Name: "i",
								Value: &ast.BinaryExpression{
									Left:     &ast.VariableExpression{Name: "i"},
									Operator: "+",
									Right:    &ast.LiteralExpression{Value: float64(1)},
								},
							},
							&ast.IfStatement{
								Condition: &ast.BinaryExpression{
									Left:     &ast.VariableExpression{Name: "i"},
									Operator: "==",
									Right:    &ast.LiteralExpression{Value: float64(3)},
								},
								ThenBranch: &ast.BreakStatement{},
								ElseBranch: nil,
							},
						},
					},
				},
				&ast.PrintStatement{
					Expression: &ast.LiteralExpression{Value: "Done"},
				},
			},
		}

		// Interpret the program
		output, err := interp.Interpret(program)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 4, len(output))
		assert.Equal(t, "0", output[0])
		assert.Equal(t, "1", output[1])
		assert.Equal(t, "2", output[2])
		assert.Equal(t, "\"Done\"", output[3])
	})

	// Test continue statement
	t.Run("Continue Statement", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a program with a for loop and continue statement
		program := &ast.Program{
			Statements: []ast.Statement{
				&ast.ForStatement{
					Initializer: &ast.VariableDeclaration{
						Name:        "i",
						Initializer: &ast.LiteralExpression{Value: float64(0)},
					},
					Condition: &ast.BinaryExpression{
						Left:     &ast.VariableExpression{Name: "i"},
						Operator: "<",
						Right:    &ast.LiteralExpression{Value: float64(5)},
					},
					Increment: &ast.BinaryExpression{
						Left:     &ast.VariableExpression{Name: "i"},
						Operator: "+",
						Right:    &ast.LiteralExpression{Value: float64(1)},
					},
					Body: &ast.Block{
						Statements: []ast.Statement{
							&ast.IfStatement{
								Condition: &ast.BinaryExpression{
									Left: &ast.BinaryExpression{
										Left:     &ast.VariableExpression{Name: "i"},
										Operator: "%",
										Right:    &ast.LiteralExpression{Value: float64(2)},
									},
									Operator: "==",
									Right:    &ast.LiteralExpression{Value: float64(0)},
								},
								ThenBranch: &ast.ContinueStatement{},
								ElseBranch: nil,
							},
							&ast.PrintStatement{
								Expression: &ast.VariableExpression{Name: "i"},
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
		assert.Equal(t, "1", output[0]) // Only odd numbers are printed
		assert.Equal(t, "3", output[1])
		assert.Equal(t, "5", output[2])
	})
}

// TestFunctionStatements tests function declaration and calls
func TestFunctionStatements(t *testing.T) {
	// Test function declaration and call
	t.Run("Function Declaration and Call", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create a function block
		functionBody := &ast.Block{
			Statements: []ast.Statement{
				&ast.ReturnStatement{
					Value: &ast.BinaryExpression{
						Left:     &ast.VariableExpression{Name: "a"},
						Operator: "+",
						Right:    &ast.VariableExpression{Name: "b"},
					},
				},
			},
		}

		// Create a program with function declaration and call
		program := &ast.Program{
			Statements: []ast.Statement{
				// Declare add function
				&ast.FunctionDeclaration{
					Name:       "add",
					Parameters: []string{"a", "b"},
					Body:       functionBody,
				},
				// Call add function
				&ast.PrintStatement{
					Expression: &ast.FunctionCallExpression{
						Callee: "add",
						Arguments: []ast.Expression{
							&ast.LiteralExpression{Value: float64(5)},
							&ast.LiteralExpression{Value: float64(3)},
						},
					},
				},
			},
		}

		// Interpret the program
		output, err := interp.Interpret(program)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 1, len(output))
		assert.Equal(t, "8", output[0])
	})

	// Test recursive function
	t.Run("Recursive Function", func(t *testing.T) {
		interp := interpreter.NewInterpreter()

		// Create factorial function body
		factorialBody := &ast.Block{
			Statements: []ast.Statement{
				&ast.IfStatement{
					Condition: &ast.BinaryExpression{
						Left:     &ast.VariableExpression{Name: "n"},
						Operator: "<=",
						Right:    &ast.LiteralExpression{Value: float64(1)},
					},
					ThenBranch: &ast.ReturnStatement{
						Value: &ast.LiteralExpression{Value: float64(1)},
					},
					ElseBranch: nil,
				},
				&ast.ReturnStatement{
					Value: &ast.BinaryExpression{
						Left: &ast.VariableExpression{Name: "n"},
						Operator: "*",
						Right: &ast.FunctionCallExpression{
							Callee: "factorial",
							Arguments: []ast.Expression{
								&ast.BinaryExpression{
									Left:     &ast.VariableExpression{Name: "n"},
									Operator: "-",
									Right:    &ast.LiteralExpression{Value: float64(1)},
								},
							},
						},
					},
				},
			},
		}

		// Create a program with recursive function
		program := &ast.Program{
			Statements: []ast.Statement{
				// Declare factorial function
				&ast.FunctionDeclaration{
					Name:       "factorial",
					Parameters: []string{"n"},
					Body:       factorialBody,
				},
				// Call factorial function
				&ast.PrintStatement{
					Expression: &ast.FunctionCallExpression{
						Callee: "factorial",
						Arguments: []ast.Expression{
							&ast.LiteralExpression{Value: float64(5)},
						},
					},
				},
			},
		}

		// Interpret the program
		output, err := interp.Interpret(program)

		// Check results
		assert.NoError(t, err)
		assert.Equal(t, 1, len(output))
		assert.Equal(t, "120", output[0]) // 5! = 120
	})
}
