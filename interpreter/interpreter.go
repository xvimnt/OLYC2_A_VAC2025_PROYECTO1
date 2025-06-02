package interpreter

import (
	"fmt"
	"github.com/mylang/interpreter/ast"
)

// Interpreter implements the Visitor interface to execute the AST
type Interpreter struct {
	ast.BaseVisitor
	globals     *Environment
	environment *Environment
	locals      map[ast.Expression]int
	output      []string

	// For control flow
	breakFlag    bool
	continueFlag bool
	returnValue  *Value
	
	// Context tracking for control flow validation
	inLoop    bool // Whether we're currently in a loop context
	inSwitch  bool // Whether we're currently in a switch context
}

// NewInterpreter creates a new interpreter instance
func NewInterpreter() *Interpreter {
	globals := NewEnvironment()
	return &Interpreter{
		globals:     globals,
		environment: globals,
		locals:      make(map[ast.Expression]int),
		output:      make([]string, 0),
	}
}

// Interpret interprets the given AST
func (i *Interpreter) Interpret(program *ast.Program) ([]string, error) {
	defer func() {
		if r := recover(); r != nil {
			switch err := r.(type) {
			case *RuntimeError:
				// It's our specific runtime error type
				fmt.Printf("Runtime error: %v\n", err.Message)
			default:
				// It's some other panic
				fmt.Printf("Runtime error: %v\n", r)
			}
		}
	}()

	i.output = make([]string, 0)
	i.breakFlag = false
	i.continueFlag = false
	i.returnValue = nil
	i.VisitProgram(program)
	return i.output, nil
}

// GetOutput returns the output buffer
func (i *Interpreter) GetOutput() []string {
	return i.output
}

// Print adds a string to the output buffer
func (i *Interpreter) Print(value string) {
	i.output = append(i.output, value)
}

// VisitProgram interprets a program node
func (i *Interpreter) VisitProgram(program *ast.Program) interface{} {
	for _, stmt := range program.Statements {
		i.execute(stmt)
		
		if i.breakFlag || i.continueFlag || i.returnValue != nil {
			break
		}
	}
	return nil
}

// Execute executes a statement with error handling
func (i *Interpreter) execute(stmt ast.Statement) interface{} {
	try := func() interface{} {
		defer func() {
			if r := recover(); r != nil {
				// Already a RuntimeError, just re-panic
				if _, ok := r.(*RuntimeError); ok {
					panic(r)
				}
				
				// Convert generic panics to RuntimeError
				panic(&RuntimeError{Message: fmt.Sprintf("Error executing statement: %v", r)})
			}
		}()
		
		return stmt.Accept(i)
	}
	
	return try()
}

// evaluate evaluates an expression with proper error handling
func (i *Interpreter) evaluate(expr ast.Expression) *Value {
	try := func() interface{} {
		defer func() {
			if r := recover(); r != nil {
				// If already a RuntimeError, just re-panic
				if _, ok := r.(*RuntimeError); ok {
					panic(r)
				}
				
				// Convert generic panics to RuntimeError
				panic(&RuntimeError{Message: fmt.Sprintf("Error evaluating expression: %v", r)})
			}
		}()
		
		return expr.Accept(i)
	}
	
	value := try()
	if value == nil {
		return NewNull()
	}
	
	// Ensure we return a Value type
	result, ok := value.(*Value)
	if !ok {
		panic(&RuntimeError{Message: fmt.Sprintf("Expression did not evaluate to a Value: %v", value)})
	}
	
	return result
}

// VisitBlock interprets a block node
func (i *Interpreter) VisitBlock(block *ast.Block) interface{} {
	prevEnv := i.environment
	defer func() { i.environment = prevEnv }()
	
	i.environment = NewEnclosedEnvironment(i.environment)
	
	for _, stmt := range block.Statements {
		i.execute(stmt)
		
		if i.breakFlag || i.continueFlag || i.returnValue != nil {
			break
		}
	}
	
	return nil
}

// VisitVariableDeclaration interprets a variable declaration with error handling
func (i *Interpreter) VisitVariableDeclaration(decl *ast.VariableDeclaration) interface{} {
	// Default to null if no initializer
	var value *Value = NewNull()
	
	// Evaluate initializer if present
	if decl.Initializer != nil {
		try := func() *Value {
			defer func() {
				if r := recover(); r != nil {
					// If it's already a RuntimeError, add context about the variable
					if err, ok := r.(*RuntimeError); ok {
						panic(&RuntimeError{
							Message: fmt.Sprintf("Error initializing variable '%s': %s", decl.Name, err.Message),
						})
					}
					
					// Otherwise wrap the error
					panic(&RuntimeError{
						Message: fmt.Sprintf("Error initializing variable '%s': %v", decl.Name, r),
					})
				}
			}()
			
			return i.evaluate(decl.Initializer)
		}
		
		value = try()
	}
	
	// Define the variable in the current environment
	try := func() {
		defer func() {
			if r := recover(); r != nil {
				// If we get an error from the environment (like redefinition)
				panic(&RuntimeError{
					Message: fmt.Sprintf("Error defining variable '%s': %v", decl.Name, r),
				})
			}
		}()
		
		i.environment.Define(decl.Name, value)
	}
	
	try()
	return nil
}

// VisitAssignmentStatement interprets an assignment statement with error handling
func (i *Interpreter) VisitAssignmentStatement(assign *ast.AssignmentStatement) interface{} {
	// Evaluate the value to assign
	value := i.evaluate(assign.Value)
	
	if assign.Index != nil {
		// Array assignment
		arrayVal, err := i.environment.Get(assign.Name)
		if err != nil {
			panic(&RuntimeError{
				Message: fmt.Sprintf("Cannot assign to undefined variable '%s'", assign.Name),
			})
		}
		
		// Verify that the target is an array
		array, ok := arrayVal.(*Value)
		if !ok || !array.IsArray() {
			panic(&RuntimeError{
				Message: fmt.Sprintf("Cannot use index on non-array variable '%s'", assign.Name),
			})
		}
		
		// Evaluate and validate the index
		indexValue := i.evaluate(assign.Index)
		if !indexValue.IsNumber() {
			panic(&RuntimeError{
				Message: fmt.Sprintf("Array index must be a number, got %s", indexValue.Type),
			})
		}
		
		indexNum, err := indexValue.AsNumber()
		if err != nil {
			panic(err) // This should already be a RuntimeError from AsNumber
		}
		
		// Convert to integer index and get array elements
		index := int(indexNum)
		elements, err := array.AsArray()
		if err != nil {
			panic(err) // This should already be a RuntimeError from AsArray
		}
		
		// Bounds checking
		if index < 0 || index >= len(elements) {
			panic(&RuntimeError{
				Message: fmt.Sprintf("Array index %d out of bounds (array length: %d)", index, len(elements)),
			})
		}
		
		// Perform the assignment
		elements[index] = value
	} else {
		// Regular variable assignment
		try := func() {
			defer func() {
				if r := recover(); r != nil {
					// If it's already a RuntimeError, just re-panic
					if err, ok := r.(*RuntimeError); ok {
						panic(err)
					}
					
					// Otherwise wrap the error
					panic(&RuntimeError{
						Message: fmt.Sprintf("Error assigning to variable '%s': %v", assign.Name, r),
					})
				}
			}()
			
			// Check if it's a local variable
			if distance, ok := i.locals[assign]; ok {
				i.environment.AssignAt(distance, assign.Name, value)
			} else {
				// Otherwise, try assigning in the global environment
				err := i.globals.Assign(assign.Name, value)
				if err != nil {
					panic(&RuntimeError{
						Message: fmt.Sprintf("Cannot assign to undefined variable '%s'", assign.Name),
					})
				}
			}
		}
		
		try()
	}
	
	return nil
}

// VisitExpressionStatement interprets an expression statement
func (i *Interpreter) VisitExpressionStatement(stmt *ast.ExpressionStatement) interface{} {
	i.evaluate(stmt.Expression)
	return nil
}

// VisitPrintStatement interprets a print statement
func (i *Interpreter) VisitPrintStatement(stmt *ast.PrintStatement) interface{} {
	value := i.evaluate(stmt.Expression)
	i.Print(value.String())
	return nil
}

// VisitIfStatement interprets an if statement with proper error handling
func (i *Interpreter) VisitIfStatement(stmt *ast.IfStatement) interface{} {
	// Evaluate the condition with proper error handling
	condition := i.evaluate(stmt.Condition)
	
	// Check condition truthiness and execute appropriate branch
	if condition.IsTruthy() {
		// Execute the 'then' branch
		return i.execute(stmt.ThenBranch)
	} else if stmt.ElseBranch != nil {
		// Execute the 'else' branch if it exists
		return i.execute(stmt.ElseBranch)
	}
	
	return nil
}

// VisitWhileStatement interprets a while statement with proper error handling
func (i *Interpreter) VisitWhileStatement(stmt *ast.WhileStatement) interface{} {
	// Track that we're in a loop context
	prevInLoop := i.inLoop
	i.inLoop = true
	defer func() { i.inLoop = prevInLoop }()
	
	for {
		// Evaluate condition first (unlike do-while)
		conditionValue := i.evaluate(stmt.Condition)
		if !conditionValue.IsTruthy() {
			break
		}
		
		// Execute the body if condition is true
		i.execute(stmt.Body)
		
		// Handle break statement
		if i.breakFlag {
			i.breakFlag = false // Clear the flag for outer loops
			break
		}
		
		// Handle continue statement
		if i.continueFlag {
			i.continueFlag = false // Clear the flag for next iteration
			continue // Go back to condition check
		}
		
		// Handle return statement (we don't clear the flag as it needs to propagate)
		if i.returnValue != nil {
			break
		}
	}
	
	return nil
}

// VisitFunctionCallExpression interprets a function call expression with comprehensive error handling
func (i *Interpreter) VisitFunctionCallExpression(expr *ast.FunctionCallExpression) interface{} {
	// Look up function in the current environment first, then globals if not found
	functionValue, err := i.lookUpVariable(expr.Callee, expr)
	if err != nil {
		// Already a properly formatted error from lookUpVariable
		panic(err)
	}
	
	// Check if it's actually a function
	if functionValue == nil || !functionValue.IsFunction() {
		panic(&RuntimeError{Message: fmt.Sprintf("'%s' is not a function", expr.Callee)})
	}
	
	// Evaluate all arguments with error handling for each argument
	args := make([]*Value, len(expr.Arguments))
	for j, arg := range expr.Arguments {
		try := func(argIdx int, argExpr ast.Expression) *Value {
			defer func() {
				if r := recover(); r != nil {
					// If it's already a RuntimeError, add context about the argument
					if err, ok := r.(*RuntimeError); ok {
						panic(&RuntimeError{
							Message: fmt.Sprintf("Error in argument %d of call to '%s': %s", argIdx+1, expr.Callee, err.Message),
						})
					}
					
					// Otherwise wrap the error
					panic(&RuntimeError{
						Message: fmt.Sprintf("Error in argument %d of call to '%s': %v", argIdx+1, expr.Callee, r),
					})
				}
			}()
			
			return i.evaluate(argExpr)
		}
		
		// Safely evaluate each argument
		args[j] = try(j, arg)
	}
	
	// Get the actual function object
	function, err := functionValue.(*Value).AsFunction()
	if err != nil {
		panic(err)
	}
	
	// Check if argument count matches parameter count
	if len(args) != len(function.Params) {
		panic(&RuntimeError{
			Message: fmt.Sprintf("Expected %d arguments but got %d", len(function.Params), len(args)),
		})
	}
	
	// Create a new environment for function execution, with the global environment as parent
	functionEnv := NewEnvironment(i.globals)
	
	// Bind parameters to arguments
	for j, param := range function.Params {
		functionEnv.Define(param, args[j])
	}
	
	// Save current environment and restore it after function execution
	prevEnv := i.environment
	i.environment = functionEnv
	
	// Save current control flags and return value
	prevBreakFlag := i.breakFlag
	prevContinueFlag := i.continueFlag
	prevReturnValue := i.returnValue
	
	// Reset control flags and return value for this function call
	i.breakFlag = false
	i.continueFlag = false
	i.returnValue = nil
	
	defer func() {
		// Restore environment and control flags
		i.environment = prevEnv
		i.breakFlag = prevBreakFlag
		i.continueFlag = prevContinueFlag
		
		// If no return statement was executed, returnValue will be nil
		if i.returnValue == nil {
			i.returnValue = prevReturnValue
		}
	}()
	
	// Execute function body
	i.executeBlock(function.Body.Statements, functionEnv)
	
	// Return the function's return value, or null if no return statement was executed
	if i.returnValue != nil {
		returnVal := i.returnValue
		i.returnValue = prevReturnValue // Restore previous return value
		return returnVal
	}
	
	return NewNull()
}

// VisitDoWhileStatement interprets a do-while statement with proper error handling
func (i *Interpreter) VisitDoWhileStatement(stmt *ast.DoWhileStatement) interface{} {
	// Track that we're in a loop context
	prevInLoop := i.inLoop
	i.inLoop = true
	defer func() { i.inLoop = prevInLoop }()
	
	// Execute body at least once
	for {
		// Execute the body
		i.execute(stmt.Body)
		
		// Handle break statement
		if i.breakFlag {
			i.breakFlag = false // Clear the flag for outer loops
			break
		}
		
		// Handle continue statement
		if i.continueFlag {
			i.continueFlag = false // Clear the flag for next iteration
			// Skip to condition check
		}
		
		// Handle return statement (we don't clear the flag as it needs to propagate)
		if i.returnValue != nil {
			break
		}
		
		// Check condition after executing body (do-while behavior)
		conditionValue := i.evaluate(stmt.Condition)
		if !conditionValue.IsTruthy() {
			break
		}
	}
	
	return nil
}

// VisitForStatement interprets a for statement with proper scoping and control flow
func (i *Interpreter) VisitForStatement(stmt *ast.ForStatement) interface{} {
	// Save current environment to restore later
	prevEnv := i.environment
	
	// Track that we're in a loop context
	prevInLoop := i.inLoop
	i.inLoop = true
	
	// Restore previous environment and loop state when done
	defer func() { 
		i.environment = prevEnv 
		i.inLoop = prevInLoop
	}()
	
	// Create a new environment for the for loop scope
	i.environment = NewEnclosedEnvironment(i.environment)
	
	// Execute initializer if present
	if stmt.Initializer != nil {
		i.execute(stmt.Initializer)
	}
	
	// Main loop execution
	for {
		// Check condition (if nil, it's an infinite loop)
		if stmt.Condition != nil {
			conditionValue := i.evaluate(stmt.Condition)
			if !conditionValue.IsTruthy() {
				break
			}
		}
		
		// Execute loop body
		i.execute(stmt.Body)
		
		// Handle break statement
		if i.breakFlag {
			i.breakFlag = false // Clear the flag for outer loops
			break
		}
		
		// Handle continue statement
		if i.continueFlag {
			i.continueFlag = false // Clear the flag for next iteration
			// Continue directly to the increment and condition check
			goto increment
		}
		
		// Handle return statement (we don't clear the flag here as it needs to propagate)
		if i.returnValue != nil {
			break
		}
		
		// Increment step
		increment:
		if stmt.Increment != nil {
			try := func() {
				defer func() {
					if r := recover(); r != nil {
						// Convert generic panics to RuntimeError
						switch err := r.(type) {
						case *RuntimeError:
							panic(err) // Re-panic with the same error
						default:
							// If it's not already a RuntimeError, wrap it
							panic(&RuntimeError{Message: fmt.Sprintf("Error in for loop increment: %v", r)})
						}
					}
				}()
				
				// Evaluate the increment expression
				i.evaluate(stmt.Increment)
			}
			
			// Execute the increment with error handling
			try()
		}
	}
	
	return nil
}

// VisitSwitchStatement interprets a switch statement with proper case matching and fallthrough behavior
func (i *Interpreter) VisitSwitchStatement(stmt *ast.SwitchStatement) interface{} {
	// Track that we're in a switch context
	prevInSwitch := i.inSwitch
	i.inSwitch = true
	defer func() { i.inSwitch = prevInSwitch }()
	
	// Evaluate the switch expression
	value := i.evaluate(stmt.Expression)
	
	// Track if any case was matched
	matched := false
	
	// Try each case
	for _, switchCase := range stmt.Cases {
		// Evaluate the case value
		caseValue := i.evaluate(switchCase.Value)
		
		// Check if the case matches
		if i.valuesEqual(value, caseValue) {
			matched = true
			
			// Execute all statements in the case block
			i.execute(switchCase.Statements)
			
			// If a break was encountered, exit the switch
			if i.breakFlag {
				i.breakFlag = false
				return nil
			}
			
			// If a return was encountered, exit the function
			if i.returnValue != nil {
				return nil
			}
			
			// In this language, we fall through to the next case without a break,
			// so we continue checking other cases
		}
	}
	
	// If no case matched and there's a default case, execute it
	if !matched && stmt.DefaultCase != nil {
		i.execute(stmt.DefaultCase)
		
		// If a break was encountered, clear the flag
		if i.breakFlag {
			i.breakFlag = false
		}
	}
	
	return nil
}

// VisitArrayAssignmentStatement interprets an array assignment statement with bounds checking
func (i *Interpreter) VisitArrayAssignmentStatement(stmt *ast.ArrayAssignmentStatement) interface{} {
	// Evaluate array expression
	arrayVal := i.evaluate(stmt.Array)
	if !arrayVal.IsArray() {
		panic(&RuntimeError{Message: "Target must be an array"})
	}
	
	// Evaluate index expression
	indexVal := i.evaluate(stmt.Index)
	if !indexVal.IsNumber() {
		panic(&RuntimeError{Message: "Array index must be a number"})
	}
	
	// Get the numeric index
	index, err := indexVal.AsNumber()
	if err != nil {
		panic(err)
	}
	indexInt := int(index)
	
	// Get the array elements
	elements, err := arrayVal.AsArray()
	if err != nil {
		panic(err)
	}
	
	// Bounds check
	if indexInt < 0 || indexInt >= len(elements) {
		panic(&RuntimeError{
			Message: fmt.Sprintf("Array index %d out of bounds (array length: %d)", indexInt, len(elements)),
		})
	}
	
	// Evaluate the value to be assigned
	value := i.evaluate(stmt.Value)
	
	// Perform the assignment
	elements[indexInt] = value
	
	return nil
}

// VisitBreakStatement interprets a break statement with validation
func (i *Interpreter) VisitBreakStatement(stmt *ast.BreakStatement) interface{} {
	// Set the break flag to signal to enclosing loop
	if !i.inLoop && !i.inSwitch {
		panic(&RuntimeError{Message: "'break' statement used outside of loop or switch"})
	}
	
	i.breakFlag = true
	return nil
}

// VisitContinueStatement interprets a continue statement with validation
func (i *Interpreter) VisitContinueStatement(stmt *ast.ContinueStatement) interface{} {
	// Validate that we're inside a loop (continue cannot be used in switch like break can)
	if !i.inLoop {
		panic(&RuntimeError{Message: "'continue' statement used outside of loop"})
	}
	
	i.continueFlag = true
	return nil
}

// VisitFunctionDeclaration interprets a function declaration with error handling
func (i *Interpreter) VisitFunctionDeclaration(decl *ast.FunctionDeclaration) interface{} {
	// Create function closure with current environment
	function := NewFunction(decl.Name, decl.Parameters, decl.Body, i.environment)
	
	// Define the function in the current environment with error handling
	try := func() {
		defer func() {
			if r := recover(); r != nil {
				// If it's already a RuntimeError, just re-panic
				if err, ok := r.(*RuntimeError); ok {
					panic(err)
				}
				
				// Otherwise wrap the error
				panic(&RuntimeError{
					Message: fmt.Sprintf("Error defining function '%s': %v", decl.Name, r),
				})
			}
		}()
		
		i.environment.Define(decl.Name, function)
	}
	
	try()
	return nil
}

// VisitSwitchCase interprets a switch case
func (i *Interpreter) VisitSwitchCase(c *ast.SwitchCase) interface{} {
	// Execute all statements in the case block
	return i.execute(c.Statements)
}

// VisitReturnStatement interprets a return statement with error handling
func (i *Interpreter) VisitReturnStatement(stmt *ast.ReturnStatement) interface{} {
	var value *Value = NewNull()
	
	// Evaluate the return value if provided
	if stmt.Value != nil {
		try := func() *Value {
			defer func() {
				if r := recover(); r != nil {
					// If it's already a RuntimeError, add context about the return statement
					if err, ok := r.(*RuntimeError); ok {
						panic(&RuntimeError{
							Message: fmt.Sprintf("Error evaluating return value: %s", err.Message),
						})
					}
					
					// Otherwise wrap the error
					panic(&RuntimeError{
						Message: fmt.Sprintf("Error evaluating return value: %v", r),
					})
				}
			}()
			
			return i.evaluate(stmt.Value)
		}
		
		value = try()
	}
	
	// Set the return value and let the executeBlock or function call handler take care of it
	i.returnValue = value
	return nil
}

// VisitBinaryExpression interprets a binary expression with proper error handling for division by zero and other errors
func (i *Interpreter) VisitBinaryExpression(expr *ast.BinaryExpression) interface{} {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)
	
	switch expr.Operator {
	// Arithmetic
	case "+":
		if left.IsNumber() && right.IsNumber() {
			leftVal, err := left.AsNumber()
			if err != nil {
				panic(err)
			}
			rightVal, err := right.AsNumber()
			if err != nil {
				panic(err)
			}
			return NewNumber(leftVal + rightVal)
		} else if left.IsString() && right.IsString() {
			leftVal, err := left.AsString()
			if err != nil {
				panic(err)
			}
			rightVal, err := right.AsString()
			if err != nil {
				panic(err)
			}
			return NewString(leftVal + rightVal)
		}
		panic(&RuntimeError{Message: "Operands must be two numbers or two strings"})
	case "-":
		i.checkNumberOperands(expr.Operator, left, right)
		leftVal, err := left.AsNumber()
		if err != nil {
			panic(err)
		}
		rightVal, err := right.AsNumber()
		if err != nil {
			panic(err)
		}
		return NewNumber(leftVal - rightVal)
	case "*":
		i.checkNumberOperands(expr.Operator, left, right)
		leftVal, err := left.AsNumber()
		if err != nil {
			panic(err)
		}
		rightVal, err := right.AsNumber()
		if err != nil {
			panic(err)
		}
		return NewNumber(leftVal * rightVal)
	case "/":
		i.checkNumberOperands(expr.Operator, left, right)
		leftVal, err := left.AsNumber()
		if err != nil {
			panic(err)
		}
		rightVal, err := right.AsNumber()
		if err != nil {
			panic(err)
		}
		if rightVal == 0 {
			panic(&RuntimeError{Message: "Division by zero"})
		}
		return NewNumber(leftVal / rightVal)
	case "%":
		i.checkNumberOperands(expr.Operator, left, right)
		leftVal, err := left.AsNumber()
		if err != nil {
			panic(err)
		}
		rightVal, err := right.AsNumber()
		if err != nil {
			panic(err)
		}
		if rightVal == 0 {
			panic(&RuntimeError{Message: "Modulo by zero"})
		}
		return NewNumber(float64(int(leftVal) % int(rightVal)))
	
	// Comparison
	case "<":
		i.checkNumberOperands(expr.Operator, left, right)
		leftVal, err := left.AsNumber()
		if err != nil {
			panic(err)
		}
		rightVal, err := right.AsNumber()
		if err != nil {
			panic(err)
		}
		return NewBoolean(leftVal < rightVal)
	case ">":
		i.checkNumberOperands(expr.Operator, left, right)
		leftVal, err := left.AsNumber()
		if err != nil {
			panic(err)
		}
		rightVal, err := right.AsNumber()
		if err != nil {
			panic(err)
		}
		return NewBoolean(leftVal > rightVal)
	case "<=":
		i.checkNumberOperands(expr.Operator, left, right)
		leftVal, err := left.AsNumber()
		if err != nil {
			panic(err)
		}
		rightVal, err := right.AsNumber()
		if err != nil {
			panic(err)
		}
		return NewBoolean(leftVal <= rightVal)
	case ">=":
		i.checkNumberOperands(expr.Operator, left, right)
		leftVal, err := left.AsNumber()
		if err != nil {
			panic(err)
		}
		rightVal, err := right.AsNumber()
		if err != nil {
			panic(err)
		}
		return NewBoolean(leftVal >= rightVal)
	
	// Equality
	case "==":
		return NewBoolean(i.valuesEqual(left, right))
	case "!=":
		return NewBoolean(!i.valuesEqual(left, right))
	
	// Logical
	case "&&":
		return NewBoolean(left.IsTruthy() && right.IsTruthy())
	case "||":
		return NewBoolean(left.IsTruthy() || right.IsTruthy())
	}
	
	return NewNull()
}

// VisitUnaryExpression interprets a unary expression with proper error handling
func (i *Interpreter) VisitUnaryExpression(expr *ast.UnaryExpression) interface{} {
	right := i.evaluate(expr.Right)
	
	switch expr.Operator {
	case "-":
		// Ensure operand is a number
		i.checkNumberOperand(expr.Operator, right)
		
		// Get the numeric value with error handling
		val, err := right.AsNumber()
		if err != nil {
			panic(err)
		}
		
		// Return the negated value
		return NewNumber(-val)
		
	case "!":
		// Logical NOT works on any type (uses truthiness)
		return NewBoolean(!right.IsTruthy())
		
	default:
		// This should never happen if the parser is correct, but handle it anyway
		panic(&RuntimeError{Message: fmt.Sprintf("Unknown unary operator '%s'", expr.Operator)})
	}
}

// VisitLiteralExpression interprets a literal expression
func (i *Interpreter) VisitLiteralExpression(expr *ast.LiteralExpression) interface{} {
	switch v := expr.Value.(type) {
	case float64:
		return NewNumber(v)
	case string:
		return NewString(v)
	case bool:
		return NewBoolean(v)
	default:
		return NewNull()
	}
}

// VisitVariableExpression interprets a variable expression with error handling
func (i *Interpreter) VisitVariableExpression(expr *ast.VariableExpression) interface{} {
	value, err := i.lookUpVariable(expr.Name, expr)
	if err != nil {
		panic(err)
	}
	return value
}

// VisitArrayAccessExpression interprets an array access expression with bounds checking and error handling
func (i *Interpreter) VisitArrayAccessExpression(expr *ast.ArrayAccessExpression) interface{} {
	// Lookup the array variable
	array, err := i.lookUpVariable(expr.Name, expr)
	if err != nil {
		panic(err)
	}
	
	// Verify it's an array type
	if !array.IsArray() {
		panic(&RuntimeError{Message: fmt.Sprintf("Variable '%s' is not an array", expr.Name)})
	}
	
	index := i.evaluate(expr.Index)
	if !index.IsNumber() {
		panic(&RuntimeError{Message: "Array index must be a number"})
	}
	
	indexVal, err := index.AsNumber()
	if err != nil {
		panic(err)
	}
	indexInt := int(indexVal)
	
	elements, err := array.AsArray()
	if err != nil {
		panic(err)
	}
	
	if indexInt < 0 || indexInt >= len(elements) {
		panic(&RuntimeError{
			Message: fmt.Sprintf("Array index %d out of bounds (array length: %d)", indexInt, len(elements)),
		})
	}
	
	return elements[indexInt]
}

// VisitArrayLiteralExpression interprets an array literal expression with error handling
func (i *Interpreter) VisitArrayLiteralExpression(expr *ast.ArrayLiteralExpression) interface{} {
	elements := make([]*Value, len(expr.Elements))
	
	// Evaluate each array element with error handling
	for j, element := range expr.Elements {
		try := func(idx int, elementExpr ast.Expression) *Value {
			defer func() {
				if r := recover(); r != nil {
					// If it's already a RuntimeError, add context about the array element
					if err, ok := r.(*RuntimeError); ok {
						panic(&RuntimeError{
							Message: fmt.Sprintf("Error evaluating array element at index %d: %s", idx, err.Message),
						})
					}
					
					// Otherwise wrap the error
					panic(&RuntimeError{
						Message: fmt.Sprintf("Error evaluating array element at index %d: %v", idx, r),
					})
				}
			}()
			
			return i.evaluate(elementExpr)
		}
		
		// Safely evaluate each element
		elements[j] = try(j, element)
	}
	
	// Create and return the array value
	return NewArray(elements)
}

// VisitFunctionCallExpression interprets a function call expression
func (i *Interpreter) VisitFunctionCallExpression(expr *ast.FunctionCallExpression) interface{} {
	callee := i.lookUpVariable(expr.Callee, expr)
	if !callee.IsFunction() {
		panic(fmt.Sprintf("'%s' is not a function", expr.Callee))
	}
	
	function, _ := callee.AsFunction()
	
	// Check arity
	if len(expr.Arguments) != len(function.Parameters) {
		panic(fmt.Sprintf("Expected %d arguments but got %d", len(function.Parameters), len(expr.Arguments)))
	}
	
	// Evaluate arguments
	args := make([]*Value, len(expr.Arguments))
	for j, arg := range expr.Arguments {
		args[j] = i.evaluate(arg)
	}
	
	// Create a new environment for the function call
	prevEnv := i.environment
	defer func() { i.environment = prevEnv }()
	
	environment := NewEnclosedEnvironment(function.Closure)
	
	// Bind parameters to arguments
	for j, param := range function.Parameters {
		environment.Define(param, args[j])
	}
	
	i.environment = environment
	
	// Execute the function body
	body := function.Body.(*ast.Block)
	i.VisitBlock(body)
	
	// Return value
	result := NewNull()
	if i.returnValue != nil {
		result = i.returnValue
		i.returnValue = nil
	}
	
	return result
}

// lookUpVariable looks up a variable in the environment with enhanced error handling
func (i *Interpreter) lookUpVariable(name string, expr ast.Expression) (*Value, error) {
	// First check for local variables (based on scope resolution)
	if distance, ok := i.locals[expr]; ok {
		value := i.environment.GetAt(distance, name)
		if value == nil {
			return nil, &RuntimeError{
				Message: fmt.Sprintf("Undefined variable '%s' in local scope", name),
			}
		}
		
		// Type assertion with safety check
		result, ok := value.(*Value)
		if !ok {
			return nil, &RuntimeError{
				Message: fmt.Sprintf("Internal error: variable '%s' does not contain a valid value", name),
			}
		}
		
		return result, nil
	}
	
	// If not found in locals, check globals
	value, err := i.globals.Get(name)
	if err != nil {
		return nil, &RuntimeError{
			Message: fmt.Sprintf("Undefined variable '%s'", name),
		}
	}
	
	// Type assertion with safety check
	result, ok := value.(*Value)
	if !ok {
		return nil, &RuntimeError{
			Message: fmt.Sprintf("Internal error: variable '%s' does not contain a valid value", name),
		}
	}
	
	return result, nil
}

// valuesEqual checks if two values are equal
func (i *Interpreter) valuesEqual(a, b *Value) bool {
	// If types are different, they're not equal
	if a.Type != b.Type {
		return false
	}
	
	// Check based on type
	switch a.Type {
	case TypeNull:
		return true
	case TypeNumber:
		a, _ := a.AsNumber()
		b, _ := b.AsNumber()
		return a == b
	case TypeString:
		a, _ := a.AsString()
		b, _ := b.AsString()
		return a == b
	case TypeBoolean:
		a, _ := a.AsBoolean()
		b, _ := b.AsBoolean()
		return a == b
	case TypeArray:
		// Arrays are equal if they have the same elements
		aElements, _ := a.AsArray()
		bElements, _ := b.AsArray()
		
		if len(aElements) != len(bElements) {
			return false
		}
		
		for j := range aElements {
			if !i.valuesEqual(aElements[j], bElements[j]) {
				return false
			}
		}
		
		return true
	case TypeFunction:
		// Functions are only equal if they're the same object
		return a.Value == b.Value
	}
	
	return false
}

// checkNumberOperand checks if an operand is a number
func (i *Interpreter) checkNumberOperand(operator string, operand *Value) {
	if !operand.IsNumber() {
		panic(&RuntimeError{Message: fmt.Sprintf("Operand for '%s' must be a number", operator)})
	}
}

// checkNumberOperands checks if both operands are numbers
func (i *Interpreter) checkNumberOperands(operator string, left, right *Value) {
	if !left.IsNumber() || !right.IsNumber() {
		panic(&RuntimeError{Message: fmt.Sprintf("Operands for '%s' must be numbers", operator)})
	}
}

// executeBlock executes a block of statements in the given environment with error handling
func (i *Interpreter) executeBlock(statements []ast.Statement, environment *Environment) {
	// Save the previous environment
	prevEnv := i.environment
	
	// Use defer to ensure environment is restored even if an error occurs
	defer func() {
		i.environment = prevEnv
	}()
	
	// Set the current environment to the new environment
	i.environment = environment
	
	// Execute each statement in the block with proper error handling
	for _, stmt := range statements {
		// Use the execute method which already has error handling
		i.execute(stmt)
		
		// If a control flow statement was executed, exit the block early
		if i.breakFlag || i.continueFlag || i.returnValue != nil {
			break
		}
	}
}
