package interpreter

import (
	"fmt"
	"github.com/mylang/interpreter/ast"
)

// VisitMatchStatement implements the V match statement visitor (similar to switch)
func (v *VInterpreter) VisitMatchStatement(match *ast.MatchStatement) interface{} {
	// Evaluate the subject expression
	subject := v.evaluate(match.Subject)
	
	// Track if we found a matching branch
	matched := false
	
	// Check each branch
	for _, branch := range match.Branches {
		// Handle 'else' branch
		if branch.IsElse {
			if !matched {
				// Execute the else branch statements
				for _, stmt := range branch.Body {
					result := stmt.Accept(v)
					
					// Check for control flow statements
					if v.returnValue != nil {
						return v.returnValue
					}
					if v.breakFlag {
						v.breakFlag = false
						return nil
					}
					if v.continueFlag {
						v.continueFlag = false
						return nil
					}
				}
			}
			continue
		}
		
		// For non-else branches, check if any pattern matches
		for _, pattern := range branch.Patterns {
			patternValue := v.evaluate(pattern)
			
			// Check for equality
			if equalValues(subject, patternValue) {
				matched = true
				
				// Execute the branch statements
				for _, stmt := range branch.Body {
					result := stmt.Accept(v)
					
					// Check for control flow statements
					if v.returnValue != nil {
						return v.returnValue
					}
					if v.breakFlag {
						v.breakFlag = false
						return nil
					}
					if v.continueFlag {
						v.continueFlag = false
						return nil
					}
				}
				
				// Don't check other patterns in this branch
				break
			}
		}
		
		// If we matched a branch, stop checking further branches
		if matched {
			break
		}
	}
	
	return nil
}

// VisitForInStatement implements V's for-in loop for arrays, strings, and maps
func (v *VInterpreter) VisitForInStatement(forIn *ast.ForInStatement) interface{} {
	// Evaluate the iterable expression
	iterable := v.evaluate(forIn.Iterable)
	
	// Create a new environment for the loop
	previous := v.environment
	v.environment = NewEnvironment(v.environment)
	
	if iterable.IsArray() {
		// Handle array iteration
		elements, _ := iterable.AsArray()
		
		// Iterate over the array
		for i, element := range elements {
			// Define loop variables
			if forIn.IndexVariable != "" {
				v.environment.Define(forIn.IndexVariable, NewNumber(float64(i)))
			}
			v.environment.Define(forIn.ValueVariable, element)
			
			// Execute the body
			for _, stmt := range forIn.Body {
				stmt.Accept(v)
				
				// Check for control flow
				if v.returnValue != nil || v.breakFlag {
					break
				}
				if v.continueFlag {
					v.continueFlag = false
					break
				}
			}
			
			// Check break flag
			if v.breakFlag {
				v.breakFlag = false
				break
			}
			
			// Check for return
			if v.returnValue != nil {
				break
			}
		}
	} else if iterable.IsString() {
		// Handle string iteration (by character)
		str, _ := iterable.AsString()
		runes := []rune(str)
		
		// Iterate over characters
		for i, r := range runes {
			// Define loop variables
			if forIn.IndexVariable != "" {
				v.environment.Define(forIn.IndexVariable, NewNumber(float64(i)))
			}
			v.environment.Define(forIn.ValueVariable, NewString(string(r)))
			
			// Execute the body
			for _, stmt := range forIn.Body {
				stmt.Accept(v)
				
				// Check for control flow
				if v.returnValue != nil || v.breakFlag {
					break
				}
				if v.continueFlag {
					v.continueFlag = false
					break
				}
			}
			
			// Check break flag
			if v.breakFlag {
				v.breakFlag = false
				break
			}
			
			// Check for return
			if v.returnValue != nil {
				break
			}
		}
	} else if iterable.IsMap() {
		// Handle map iteration
		mpValue, ok := iterable.Value.(*MapValue)
		if !ok {
			panic(RuntimeError{message: "Invalid map value"})
		}
		
		// Iterate over map entries
		for k, v := range mpValue.Entries {
			// Define loop variables
			if forIn.IndexVariable != "" {
				// Convert key to Value based on its type
				var keyValue *Value
				switch key := k.(type) {
				case string:
					keyValue = NewString(key)
				case float64:
					keyValue = NewNumber(key)
				case bool:
					keyValue = NewBoolean(key)
				default:
					keyValue = NewString(fmt.Sprintf("%v", key))
				}
				
				v.environment.Define(forIn.IndexVariable, keyValue)
			}
			v.environment.Define(forIn.ValueVariable, v)
			
			// Execute the body
			for _, stmt := range forIn.Body {
				stmt.Accept(v)
				
				// Check for control flow
				if v.returnValue != nil || v.breakFlag {
					break
				}
				if v.continueFlag {
					v.continueFlag = false
					break
				}
			}
			
			// Check break flag
			if v.breakFlag {
				v.breakFlag = false
				break
			}
			
			// Check for return
			if v.returnValue != nil {
				break
			}
		}
	} else {
		panic(RuntimeError{message: fmt.Sprintf("Cannot iterate over %v", iterable)})
	}
	
	// Restore the previous environment
	v.environment = previous
	
	return nil
}

// VisitForCStatement implements V's condition-only for loop (while loop)
func (v *VInterpreter) VisitForCStatement(forC *ast.ForCStatement) interface{} {
	// Create a new environment for the loop
	previous := v.environment
	v.environment = NewEnvironment(v.environment)
	
	// While loop implementation
	for {
		// Evaluate the condition
		condition := v.evaluate(forC.Condition)
		
		if !condition.IsTruthy() {
			break
		}
		
		// Execute the body
		for _, stmt := range forC.Body.Statements {
			stmt.Accept(v)
			
			// Check for control flow
			if v.returnValue != nil || v.breakFlag {
				break
			}
			if v.continueFlag {
				v.continueFlag = false
				break
			}
		}
		
		// Check break flag
		if v.breakFlag {
			v.breakFlag = false
			break
		}
		
		// Check for return
		if v.returnValue != nil {
			break
		}
	}
	
	// Restore the previous environment
	v.environment = previous
	
	return nil
}

// VisitDeferStatement implements V's defer statement
func (v *VInterpreter) VisitDeferStatement(defer_ *ast.DeferStatement) interface{} {
	// Add the statement to the defer stack to be executed later
	v.deferStack = append(v.deferStack, defer_.Statement)
	
	return nil
}

// ExecuteDeferredStatements executes all deferred statements in LIFO order
// This should be called at function return points
func (v *VInterpreter) ExecuteDeferredStatements() {
	// Store the original return value
	returnValue := v.returnValue
	
	// Execute deferred statements in LIFO order
	for i := len(v.deferStack) - 1; i >= 0; i-- {
		// Execute deferred statement
		v.deferStack[i].Accept(v)
		
		// Ignore break/continue in deferred statements
		v.breakFlag = false
		v.continueFlag = false
		
		// Restore the original return value if a new one was set
		if v.returnValue != nil && v.returnValue != returnValue {
			v.returnValue = returnValue
		}
	}
	
	// Clear the defer stack
	v.deferStack = make([]ast.Statement, 0)
}
