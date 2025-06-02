package interpreter

import (
	"fmt"
)

// Environment represents a lexical scope for variables and functions
type Environment struct {
	Values    map[string]interface{}
	Enclosing *Environment
}

// NewEnvironment creates a new environment with no enclosing environment
func NewEnvironment() *Environment {
	return &Environment{
		Values:    make(map[string]interface{}),
		Enclosing: nil,
	}
}

// NewEnclosedEnvironment creates a new environment with the given enclosing environment
func NewEnclosedEnvironment(enclosing *Environment) *Environment {
	return &Environment{
		Values:    make(map[string]interface{}),
		Enclosing: enclosing,
	}
}

// Define defines a new variable in the current environment
func (e *Environment) Define(name string, value interface{}) {
	e.Values[name] = value
}

// Get retrieves a variable's value from the environment
func (e *Environment) Get(name string) (interface{}, error) {
	if value, ok := e.Values[name]; ok {
		return value, nil
	}

	if e.Enclosing != nil {
		return e.Enclosing.Get(name)
	}

	return nil, fmt.Errorf("undefined variable '%s'", name)
}

// Assign assigns a new value to an existing variable
func (e *Environment) Assign(name string, value interface{}) error {
	if _, ok := e.Values[name]; ok {
		e.Values[name] = value
		return nil
	}

	if e.Enclosing != nil {
		return e.Enclosing.Assign(name, value)
	}

	return fmt.Errorf("undefined variable '%s'", name)
}

// GetAt retrieves a variable from a specific environment at the given distance
func (e *Environment) GetAt(distance int, name string) interface{} {
	return e.ancestor(distance).Values[name]
}

// AssignAt assigns a value to a variable in a specific environment at the given distance
func (e *Environment) AssignAt(distance int, name string, value interface{}) {
	e.ancestor(distance).Values[name] = value
}

// ancestor returns the environment at the given distance
func (e *Environment) ancestor(distance int) *Environment {
	environment := e
	for i := 0; i < distance; i++ {
		environment = environment.Enclosing
	}
	return environment
}
