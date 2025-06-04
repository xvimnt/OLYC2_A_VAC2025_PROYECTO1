package interpreter

import (
	"fmt"
	"github.com/mylang/interpreter/ast"
)

// VisitModuleDeclaration implements the module declaration visitor
func (v *VInterpreter) VisitModuleDeclaration(module *ast.ModuleDeclaration) interface{} {
	moduleName := module.Name
	
	// Create a new module or get existing
	if _, exists := v.modules[moduleName]; !exists {
		v.modules[moduleName] = &Module{
			Name:      moduleName,
			Types:     make(map[string]interface{}),
			Constants: make(map[string]*Value),
			Functions: make(map[string]*Function),
		}
	}
	
	v.currentModule = v.modules[moduleName]
	return nil
}

// VisitImportDeclaration implements the import declaration visitor
func (v *VInterpreter) VisitImportDeclaration(imp *ast.ImportDeclaration) interface{} {
	// In a real implementation, this would load the module
	// For now, we'll just track that it was imported
	importPath := imp.Path
	
	// Create a placeholder module if it doesn't exist
	if _, exists := v.modules[importPath]; !exists {
		v.modules[importPath] = &Module{
			Name:      importPath,
			Types:     make(map[string]interface{}),
			Constants: make(map[string]*Value),
			Functions: make(map[string]*Function),
		}
	}
	
	return nil
}

// VisitConstDeclaration implements the constant declaration visitor
func (v *VInterpreter) VisitConstDeclaration(constant *ast.ConstDeclaration) interface{} {
	// Evaluate the initializer
	value := v.evaluate(constant.Value)
	
	// Store in the current module's constants
	v.currentModule.Constants[constant.Name] = value
	
	// Also define in the current environment (for immediate use)
	v.environment.Define(constant.Name, value)
	
	// Mark as immutable
	v.immutableVars[constant.Name] = true
	
	return nil
}

// VisitTypedVariable implements the typed variable declaration visitor for V
func (v *VInterpreter) VisitTypedVariable(variable *ast.TypedVariable) interface{} {
	// Evaluate the initializer if provided
	var value *Value = NewNull()
	
	if variable.Initializer != nil {
		value = v.evaluate(variable.Initializer)
		
		// Perform type checking based on TypeName if necessary
		// This is simplified, you'd want more robust type checking
		switch variable.TypeName {
		case "int", "i8", "i16", "i32", "i64":
			if !value.IsNumber() {
				panic(RuntimeError{message: fmt.Sprintf("Cannot assign %v to %s type", value, variable.TypeName)})
			}
		case "string":
			if !value.IsString() {
				panic(RuntimeError{message: fmt.Sprintf("Cannot assign %v to string type", value)})
			}
		case "bool":
			if !value.IsBoolean() {
				panic(RuntimeError{message: fmt.Sprintf("Cannot assign %v to bool type", value)})
			}
		}
	}
	
	// Define in the current environment
	v.environment.Define(variable.Name, value)
	
	// Mark as immutable unless explicitly mutable
	if !variable.IsMutable {
		v.immutableVars[variable.Name] = true
	}
	
	return nil
}
