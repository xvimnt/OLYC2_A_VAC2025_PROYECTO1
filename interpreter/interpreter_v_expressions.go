package interpreter

import (
	"fmt"
	"github.com/mylang/interpreter/ast"
)

// VisitMethodCallExpression implements the V method call expression visitor
func (v *VInterpreter) VisitMethodCallExpression(call *ast.MethodCallExpression) interface{} {
	// Evaluate the object expression
	object := v.evaluate(call.Object)
	
	// Evaluate the argument expressions
	arguments := make([]*Value, len(call.Arguments))
	for i, arg := range call.Arguments {
		arguments[i] = v.evaluate(arg)
	}
	
	// Get the method name
	methodName := call.Method
	
	// Handle method calls on different types
	switch {
	case object.IsStruct():
		// Get the struct value and type
		structValue, ok := object.Value.(*StructValue)
		if !ok {
			panic(RuntimeError{message: "Invalid struct value"})
		}
		
		// Look up the method in the current module
		fullMethodName := fmt.Sprintf("%s.%s", structValue.Name, methodName)
		methodFunc, exists := v.currentModule.Functions[fullMethodName]
		if !exists {
			// Check imported modules
			for _, module := range v.modules {
				if module == v.currentModule {
					continue // Already checked
				}
				
				methodFunc, exists = module.Functions[fullMethodName]
				if exists {
					break
				}
			}
			
			if !exists {
				panic(RuntimeError{message: fmt.Sprintf("Method '%s' not found on type '%s'", methodName, structValue.Name)})
			}
		}
		
		// Create a new environment for the method call with the receiver as 'this' or 'self'
		environment := NewEnvironment(methodFunc.Closure)
		environment.Define("self", object) // V uses 'self' instead of 'this'
		
		// Add the arguments to the environment
		for i, param := range methodFunc.Parameters {
			if i < len(arguments) {
				environment.Define(param, arguments[i])
			} else {
				environment.Define(param, NewNull())
			}
		}
		
		// Execute the method body
		previous := v.environment
		v.environment = environment
		
		// Execute the method body
		methodFunc.Body.(*ast.BlockStatement).Accept(v)
		
		// Execute any deferred statements
		v.ExecuteDeferredStatements()
		
		// Save the return value and reset the flag
		result := v.returnValue
		v.returnValue = nil
		
		// Restore the previous environment
		v.environment = previous
		
		// Return the result
		if result == nil {
			return NewNull()
		}
		return result
		
	case object.IsArray() || object.IsString() || object.IsMap():
		// Handle built-in methods for arrays, strings and maps
		elements, isArray := object.AsArray()
		str, isString := object.AsString()
		
		switch methodName {
		case "len":
			if isArray {
				return NewNumber(float64(len(elements)))
			} else if isString {
				return NewNumber(float64(len(str)))
			} else if object.IsMap() {
				mapValue, _ := object.Value.(*MapValue)
				return NewNumber(float64(len(mapValue.Entries)))
			}
		case "contains":
			if len(arguments) < 1 {
				panic(RuntimeError{message: "contains() requires at least one argument"})
			}
			
			if isArray {
				searchValue := arguments[0]
				for _, elem := range elements {
					if equalValues(elem, searchValue) {
						return NewBoolean(true)
					}
				}
				return NewBoolean(false)
			} else if isString {
				searchStr, isStr := arguments[0].AsString()
				if !isStr {
					panic(RuntimeError{message: "contains() requires a string argument for string receiver"})
				}
				return NewBoolean(contains(str, searchStr))
			} else if object.IsMap() {
				mapValue, _ := object.Value.(*MapValue)
				key := arguments[0].Value
				_, exists := mapValue.Entries[key]
				return NewBoolean(exists)
			}
		case "keys":
			if object.IsMap() {
				mapValue, _ := object.Value.(*MapValue)
				keys := make([]*Value, 0, len(mapValue.Entries))
				
				for k := range mapValue.Entries {
					switch key := k.(type) {
					case string:
						keys = append(keys, NewString(key))
					case float64:
						keys = append(keys, NewNumber(key))
					case bool:
						keys = append(keys, NewBoolean(key))
					default:
						keys = append(keys, NewString(fmt.Sprintf("%v", key)))
					}
				}
				
				return NewArray(keys)
			}
		}
		
		panic(RuntimeError{message: fmt.Sprintf("Method '%s' not found", methodName)})
		
	default:
		panic(RuntimeError{message: fmt.Sprintf("Cannot call method on %v", object)})
	}
}

// VisitSelectorExpression implements field access in structs or module elements
func (v *VInterpreter) VisitSelectorExpression(selector *ast.SelectorExpression) interface{} {
	// Evaluate the object expression
	object := v.evaluate(selector.Object)
	field := selector.Field
	
	// Handle different object types
	switch {
	case object.IsStruct():
		value, err := object.GetField(field)
		if err != nil {
			panic(RuntimeError{message: err.Error()})
		}
		return value
		
	case object.IsModule():
		// This is not fully implemented, but would be used for module imports
		moduleName, ok := object.Value.(string)
		if !ok {
			panic(RuntimeError{message: "Invalid module value"})
		}
		
		// Look up the module
		module, exists := v.modules[moduleName]
		if !exists {
			panic(RuntimeError{message: fmt.Sprintf("Module '%s' not found", moduleName)})
		}
		
		// Check for constants first
		if value, exists := module.Constants[field]; exists {
			return value
		}
		
		// Check for functions
		if function, exists := module.Functions[field]; exists {
			return NewFunction(function)
		}
		
		// Check for types
		if _, exists := module.Types[field]; exists {
			// Return the type name as a string for now
			// In a real implementation, this would return a proper type value
			return NewString(field)
		}
		
		panic(RuntimeError{message: fmt.Sprintf("'%s' not found in module '%s'", field, moduleName)})
		
	default:
		panic(RuntimeError{message: fmt.Sprintf("Cannot access field '%s' on %v", field, object)})
	}
}

// VisitMapLiteral implements V map literals
func (v *VInterpreter) VisitMapLiteral(mapLiteral *ast.MapLiteral) interface{} {
	// Create a new map value
	mapValue := NewMap(mapLiteral.KeyType, mapLiteral.ValueType)
	
	// Add all entries
	for _, entry := range mapLiteral.Entries {
		key := v.evaluate(entry.Key)
		value := v.evaluate(entry.Value)
		
		// Get the actual key value
		var keyValue interface{}
		switch {
		case key.IsString():
			keyStr, _ := key.AsString()
			keyValue = keyStr
		case key.IsNumber():
			keyNum, _ := key.AsNumber()
			keyValue = keyNum
		case key.IsBoolean():
			keyBool, _ := key.AsBoolean()
			keyValue = keyBool
		default:
			panic(RuntimeError{message: fmt.Sprintf("Invalid map key type: %v", key)})
		}
		
		// Set the map entry
		mapValue.SetMapValue(keyValue, value)
	}
	
	return mapValue
}

// VisitOptionExpression implements V option expressions (Some/None)
func (v *VInterpreter) VisitOptionExpression(option *ast.OptionExpression) interface{} {
	// If there's a value, it's Some(value)
	if option.Value != nil {
		value := v.evaluate(option.Value)
		return NewOption(value)
	}
	
	// Otherwise it's None
	return NewNone()
}

// VisitOrBlockExpression implements V or block expressions for option types
func (v *VInterpreter) VisitOrBlockExpression(orBlock *ast.OrBlockExpression) interface{} {
	// Evaluate the option expression
	option := v.evaluate(orBlock.OptionExpr)
	
	// Check if it's an option type
	if option.Type != OPTION {
		panic(RuntimeError{message: fmt.Sprintf("Cannot use or block on non-option type %v", option)})
	}
	
	// If the option has a value, return it
	if option.IsSome() {
		value, _ := option.GetOptionValue()
		return value
	}
	
	// Otherwise, execute the or block
	previous := v.environment
	v.environment = NewEnvironment(v.environment)
	
	var result interface{}
	
	// Execute the or block
	for _, stmt := range orBlock.OrBlock {
		result = stmt.Accept(v)
		
		// Check for control flow statements
		if v.returnValue != nil || v.breakFlag || v.continueFlag {
			break
		}
	}
	
	// Restore the environment
	v.environment = previous
	
	// Convert the result to a Value if it's not already
	if result == nil {
		return NewNull()
	}
	
	if value, ok := result.(*Value); ok {
		return value
	}
	
	// Otherwise wrap the result in a Value
	return &Value{
		Type:  ANY,
		Value: result,
	}
}

// Helper function for string contains
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Helper function to compare values for equality
func equalValues(v1, v2 *Value) bool {
	if v1.Type != v2.Type {
		return false
	}
	
	switch v1.Type {
	case NUMBER:
		num1, _ := v1.AsNumber()
		num2, _ := v2.AsNumber()
		return num1 == num2
	case STRING:
		str1, _ := v1.AsString()
		str2, _ := v2.AsString()
		return str1 == str2
	case BOOLEAN:
		bool1, _ := v1.AsBoolean()
		bool2, _ := v2.AsBoolean()
		return bool1 == bool2
	case NULL:
		return true
	case OPTION:
		// Both are None
		if v1.IsNone() && v2.IsNone() {
			return true
		}
		// Both are Some
		if v1.IsSome() && v2.IsSome() {
			val1, _ := v1.GetOptionValue()
			val2, _ := v2.GetOptionValue()
			return equalValues(val1, val2)
		}
		return false
	case STRUCT:
		// For now, structs are equal if they have the same name
		// A more complete implementation would compare all fields
		s1, _ := v1.Value.(*StructValue)
		s2, _ := v2.Value.(*StructValue)
		return s1.Name == s2.Name
	case ENUM:
		e1, _ := v1.Value.(*EnumValue)
		e2, _ := v2.Value.(*EnumValue)
		return e1.EnumType == e2.EnumType && e1.Name == e2.Name
	default:
		// For other types, compare the raw values
		return v1.Value == v2.Value
	}
}
