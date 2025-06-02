package interpreter

import (
	"fmt"
	"strings"
)

// RuntimeError represents an error that occurs during execution
type RuntimeError struct {
	Message string
	Line    int
}

func (e *RuntimeError) Error() string {
	return fmt.Sprintf("[line %d] Runtime Error: %s", e.Line, e.Message)
}

// Value types
const (
	TypeNull = iota
	TypeNumber
	TypeString
	TypeBoolean
	TypeArray
	TypeFunction
)

// Value represents a runtime value in the interpreter
type Value struct {
	Type  int
	Value interface{}
}

// NewNull creates a new null value
func NewNull() *Value {
	return &Value{Type: TypeNull, Value: nil}
}

// NewNumber creates a new number value
func NewNumber(value float64) *Value {
	return &Value{Type: TypeNumber, Value: value}
}

// NewString creates a new string value
func NewString(value string) *Value {
	return &Value{Type: TypeString, Value: value}
}

// NewBoolean creates a new boolean value
func NewBoolean(value bool) *Value {
	return &Value{Type: TypeBoolean, Value: value}
}

// NewArray creates a new array value
func NewArray(elements []*Value) *Value {
	return &Value{Type: TypeArray, Value: elements}
}

// Function represents a callable function
type Function struct {
	Name       string
	Parameters []string
	Body       interface{} // This will be an *ast.Block
	Closure    *Environment
}

// NewFunction creates a new function value
func NewFunction(name string, parameters []string, body interface{}, closure *Environment) *Value {
	return &Value{Type: TypeFunction, Value: &Function{
		Name:       name,
		Parameters: parameters,
		Body:       body,
		Closure:    closure,
	}}
}

// IsNull checks if the value is null
func (v *Value) IsNull() bool {
	return v.Type == TypeNull
}

// IsNumber checks if the value is a number
func (v *Value) IsNumber() bool {
	return v.Type == TypeNumber
}

// IsString checks if the value is a string
func (v *Value) IsString() bool {
	return v.Type == TypeString
}

// IsBoolean checks if the value is a boolean
func (v *Value) IsBoolean() bool {
	return v.Type == TypeBoolean
}

// IsArray checks if the value is an array
func (v *Value) IsArray() bool {
	return v.Type == TypeArray
}

// IsFunction checks if the value is a function
func (v *Value) IsFunction() bool {
	return v.Type == TypeFunction
}

// AsNumber converts the value to a number
func (v *Value) AsNumber() (float64, error) {
	if !v.IsNumber() {
		return 0, fmt.Errorf("expected a number, got %s", v.TypeName())
	}
	return v.Value.(float64), nil
}

// AsString converts the value to a string
func (v *Value) AsString() (string, error) {
	if !v.IsString() {
		return "", fmt.Errorf("expected a string, got %s", v.TypeName())
	}
	return v.Value.(string), nil
}

// AsBoolean converts the value to a boolean
func (v *Value) AsBoolean() (bool, error) {
	if !v.IsBoolean() {
		return false, fmt.Errorf("expected a boolean, got %s", v.TypeName())
	}
	return v.Value.(bool), nil
}

// AsArray converts the value to an array
func (v *Value) AsArray() ([]*Value, error) {
	if !v.IsArray() {
		return nil, fmt.Errorf("expected an array, got %s", v.TypeName())
	}
	return v.Value.([]*Value), nil
}

// AsFunction converts the value to a function
func (v *Value) AsFunction() (*Function, error) {
	if !v.IsFunction() {
		return nil, fmt.Errorf("expected a function, got %s", v.TypeName())
	}
	return v.Value.(*Function), nil
}

// TypeName returns the name of the value's type
func (v *Value) TypeName() string {
	switch v.Type {
	case TypeNull:
		return "null"
	case TypeNumber:
		return "number"
	case TypeString:
		return "string"
	case TypeBoolean:
		return "boolean"
	case TypeArray:
		return "array"
	case TypeFunction:
		return "function"
	default:
		return "unknown"
	}
}

// String returns a string representation of the value
func (v *Value) String() string {
	switch v.Type {
	case TypeNull:
		return "null"
	case TypeNumber:
		num := v.Value.(float64)
		// If it's an integer, don't show the decimal point
		if num == float64(int(num)) {
			return fmt.Sprintf("%d", int(num))
		}
		return fmt.Sprintf("%g", num)
	case TypeString:
		return fmt.Sprintf("\"%s\"", v.Value.(string))
	case TypeBoolean:
		if v.Value.(bool) {
			return "true"
		}
		return "false"
	case TypeArray:
		elements := v.Value.([]*Value)
		strs := make([]string, len(elements))
		for i, elem := range elements {
			strs[i] = elem.String()
		}
		return "[" + strings.Join(strs, ", ") + "]"
	case TypeFunction:
		fn := v.Value.(*Function)
		return fmt.Sprintf("<function %s>", fn.Name)
	default:
		return "<unknown>"
	}
}

// IsTruthy returns true if the value is considered truthy in conditionals
func (v *Value) IsTruthy() bool {
	switch v.Type {
	case TypeNull:
		return false
	case TypeBoolean:
		return v.Value.(bool)
	case TypeNumber:
		return v.Value.(float64) != 0
	case TypeString:
		return v.Value.(string) != ""
	case TypeArray:
		elements := v.Value.([]*Value)
		return len(elements) > 0
	default:
		return true
	}
}
