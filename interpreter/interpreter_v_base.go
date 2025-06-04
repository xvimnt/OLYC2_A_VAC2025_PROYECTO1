package interpreter

import (
	"fmt"
	"github.com/mylang/interpreter/ast"
)

// VInterpreter extends the base Interpreter to support V language features
type VInterpreter struct {
	*Interpreter
	
	// V-specific state
	modules        map[string]*Module
	currentModule  *Module
	currentStruct  string
	immutableVars  map[string]bool // Track which variables are immutable
	deferStack     []ast.Statement // Stack of deferred statements
}

// Module represents a V module
type Module struct {
	Name      string
	Types     map[string]interface{} // structs, enums, interfaces
	Constants map[string]*Value
	Functions map[string]*Function
}

// StructType represents a V struct type definition
type StructType struct {
	Name   string
	Fields map[string]StructField
}

// StructField represents a field in a struct definition
type StructField struct {
	Type     string
	IsMutable bool
}

// EnumType represents a V enum type definition
type EnumType struct {
	Name   string
	Values map[string]float64
}

// InterfaceType represents a V interface type definition
type InterfaceType struct {
	Name    string
	Methods map[string]MethodSignature
}

// MethodSignature represents a method signature in an interface
type MethodSignature struct {
	Params     []string
	ParamTypes []string
	ReturnType string
}

// NewVInterpreter creates a new V interpreter
func NewVInterpreter() *VInterpreter {
	baseInterpreter := NewInterpreter()
	return &VInterpreter{
		Interpreter:   baseInterpreter,
		modules:       make(map[string]*Module),
		immutableVars: make(map[string]bool),
		deferStack:    make([]ast.Statement, 0),
	}
}

// InterpretV interprets a V program
func (v *VInterpreter) InterpretV(program *ast.Program) ([]string, error) {
	// Create the default module if no module was declared
	if v.currentModule == nil {
		v.currentModule = &Module{
			Name:      "main",
			Types:     make(map[string]interface{}),
			Constants: make(map[string]*Value),
			Functions: make(map[string]*Function),
		}
		v.modules["main"] = v.currentModule
	}
	
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(RuntimeError); ok {
				// Handle runtime error
				fmt.Printf("Runtime error: %v\n", err.message)
			} else {
				// Re-panic if it's not a RuntimeError
				panic(r)
			}
		}
	}()
	
	// Execute the program
	result := program.Accept(v)
	
	// Convert to string array if needed
	if result == nil {
		return []string{}, nil
	}
	
	if strResult, ok := result.([]string); ok {
		return strResult, nil
	}
	
	return []string{fmt.Sprintf("%v", result)}, nil
}
