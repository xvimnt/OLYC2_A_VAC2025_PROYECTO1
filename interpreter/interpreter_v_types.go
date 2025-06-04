package interpreter

import (
	"fmt"
	"github.com/mylang/interpreter/ast"
)

// VisitStructDeclaration implements the struct declaration visitor
func (v *VInterpreter) VisitStructDeclaration(structDecl *ast.StructDeclaration) interface{} {
	structName := structDecl.Name
	fields := make(map[string]StructField)
	
	// Process all fields
	for _, field := range structDecl.Fields {
		fields[field.Name] = StructField{
			Type:     field.Type,
			IsMutable: field.IsMutable,
		}
	}
	
	// Create struct type definition
	structType := &StructType{
		Name:   structName,
		Fields: fields,
	}
	
	// Store in the current module's types
	v.currentModule.Types[structName] = structType
	
	return nil
}

// VisitEnumDeclaration implements the enum declaration visitor
func (v *VInterpreter) VisitEnumDeclaration(enumDecl *ast.EnumDeclaration) interface{} {
	enumName := enumDecl.Name
	values := make(map[string]float64)
	
	// Process all enum values
	nextValue := 0.0
	for _, val := range enumDecl.Values {
		if val.Value != nil {
			// Get explicit value
			valueExpr := v.evaluate(val.Value)
			if !valueExpr.IsNumber() {
				panic(RuntimeError{message: fmt.Sprintf("Enum value must be a number, got %v", valueExpr)})
			}
			
			numVal, _ := valueExpr.AsNumber()
			values[val.Name] = numVal
			nextValue = numVal + 1
		} else {
			// Assign implicit value
			values[val.Name] = nextValue
			nextValue++
		}
	}
	
	// Create enum type definition
	enumType := &EnumType{
		Name:   enumName,
		Values: values,
	}
	
	// Store in the current module's types
	v.currentModule.Types[enumName] = enumType
	
	return nil
}

// VisitInterfaceDeclaration implements the interface declaration visitor
func (v *VInterpreter) VisitInterfaceDeclaration(interfaceDecl *ast.InterfaceDeclaration) interface{} {
	interfaceName := interfaceDecl.Name
	methods := make(map[string]MethodSignature)
	
	// Process all method signatures
	for _, method := range interfaceDecl.Methods {
		paramTypes := make([]string, len(method.Params))
		paramNames := make([]string, len(method.Params))
		
		for i, param := range method.Params {
			paramTypes[i] = param.Type
			paramNames[i] = param.Name
		}
		
		methods[method.Name] = MethodSignature{
			Params:     paramNames,
			ParamTypes: paramTypes,
			ReturnType: method.ReturnType,
		}
	}
	
	// Create interface type definition
	interfaceType := &InterfaceType{
		Name:    interfaceName,
		Methods: methods,
	}
	
	// Store in the current module's types
	v.currentModule.Types[interfaceName] = interfaceType
	
	return nil
}

// VisitVFunctionDeclaration implements V function declaration
func (v *VInterpreter) VisitVFunctionDeclaration(funcDecl *ast.VFunctionDeclaration) interface{} {
	// Create function object
	function := &Function{
		Name: funcDecl.Name,
		Parameters: make([]string, len(funcDecl.Params)),
		Body: funcDecl.Body,
		Closure: v.environment,
	}
	
	// Extract parameter names
	for i, param := range funcDecl.Params {
		function.Parameters[i] = param.Name
	}
	
	// If this is a method (has a receiver), store it with the appropriate struct type
	if funcDecl.Receiver != nil {
		receiverType := funcDecl.Receiver.Type
		
		// Lookup the struct type
		structTypeObj, exists := v.currentModule.Types[receiverType]
		if !exists {
			panic(RuntimeError{message: fmt.Sprintf("Cannot declare method for unknown type %s", receiverType)})
		}
		
		// Make sure it's a struct
		structType, ok := structTypeObj.(*StructType)
		if !ok {
			panic(RuntimeError{message: fmt.Sprintf("Method receiver must be a struct type, %s is not a struct", receiverType)})
		}
		
		// Create a method with the function and receiver type
		method := &Method{
			Function: function,
			ReceiverType: receiverType,
		}
		
		// Store the method in the struct's method map (added to struct type)
		// In a real implementation, you'd have a methods field in StructType
		// But here we're just going to store it in the current module functions with a special name
		methodName := fmt.Sprintf("%s.%s", receiverType, funcDecl.Name)
		v.currentModule.Functions[methodName] = function
	} else {
		// Regular function, store in module
		v.currentModule.Functions[funcDecl.Name] = function
		
		// Also define in current environment for immediate use
		v.environment.Define(funcDecl.Name, NewFunction(function))
	}
	
	return nil
}
