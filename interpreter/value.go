package interpreter

import (
	"fmt"
	"strings"
)

// ValueType represents the type of a runtime value
type ValueType int

const (
	NULL ValueType = iota
	NUMBER
	STRING
	BOOLEAN
	ARRAY
	FUNCTION
)

// Value represents a runtime value in the interpreter
type Value struct {
	Type  ValueType
	Value interface{}
}

// NewNullValue creates a new null value
func NewNullValue() *Value {
	return &Value{Type: NULL, Value: nil}
}

// NewNull is an alias for NewNullValue for easier use
func NewNull() *Value {
	return NewNullValue()
}

// NewNumberValue creates a new number value
func NewNumberValue(value float64) *Value {
	return &Value{Type: NUMBER, Value: value}
}

// NewNumber is an alias for NewNumberValue for easier use
func NewNumber(value float64) *Value {
	return NewNumberValue(value)
}

// NewStringValue creates a new string value
func NewStringValue(value string) *Value {
	return &Value{Type: STRING, Value: value}
}

// NewString is an alias for NewStringValue for easier use
func NewString(value string) *Value {
	return NewStringValue(value)
}

// NewBooleanValue creates a new boolean value
func NewBooleanValue(value bool) *Value {
	return &Value{Type: BOOLEAN, Value: value}
}

// NewBoolean is an alias for NewBooleanValue for easier use
func NewBoolean(value bool) *Value {
	return NewBooleanValue(value)
}

// NewArrayValue creates a new array value
func NewArrayValue(elements []*Value) *Value {
	return &Value{Type: ARRAY, Value: elements}
}

// NewArray is an alias for NewArrayValue for easier use
func NewArray(elements []*Value) *Value {
	return NewArrayValue(elements)
}

// NewFunctionValue creates a new function value
func NewFunctionValue(fn *Function) *Value {
	return &Value{Type: FUNCTION, Value: fn}
}

// NewFunction is an alias for NewFunctionValue for easier use
func NewFunction(fn *Function) *Value {
	return NewFunctionValue(fn)
}

// IsNull checks if the value is null
func (v *Value) IsNull() bool {
	return v.Type == NULL
}

// IsNumber checks if the value is a number
func (v *Value) IsNumber() bool {
	return v.Type == NUMBER
}

// IsString checks if the value is a string
func (v *Value) IsString() bool {
	return v.Type == STRING
}

// IsBoolean checks if the value is a boolean
func (v *Value) IsBoolean() bool {
	return v.Type == BOOLEAN
}

// IsArray checks if the value is an array
func (v *Value) IsArray() bool {
	return v.Type == ARRAY
}

// IsFunction checks if the value is a function
func (v *Value) IsFunction() bool {
	return v.Type == FUNCTION
}

// AsNumber returns the value as a number
func (v *Value) AsNumber() (float64, error) {
	if !v.IsNumber() {
		return 0, &RuntimeError{Message: "Value is not a number"}
	}
	return v.Value.(float64), nil
}

// AsString returns the value as a string
func (v *Value) AsString() (string, error) {
	if !v.IsString() {
		return "", &RuntimeError{Message: "Value is not a string"}
	}
	return v.Value.(string), nil
}

// AsBoolean returns the value as a boolean
func (v *Value) AsBoolean() (bool, error) {
	if !v.IsBoolean() {
		return false, &RuntimeError{Message: "Value is not a boolean"}
	}
	return v.Value.(bool), nil
}

// AsArray returns the value as an array
func (v *Value) AsArray() ([]*Value, error) {
	if !v.IsArray() {
		return nil, &RuntimeError{Message: "Value is not an array"}
	}
	return v.Value.([]*Value), nil
}

// AsFunction returns the value as a function
func (v *Value) AsFunction() (*Function, error) {
	if !v.IsFunction() {
		return nil, &RuntimeError{Message: "Value is not a function"}
	}
	return v.Value.(*Function), nil
}

// String returns a string representation of the value
func (v *Value) String() string {
	switch v.Type {
	case NULL:
		return "null"
	case NUMBER:
		num, _ := v.AsNumber()
		// If the number is an integer, print it as such
		if num == float64(int(num)) {
			return fmt.Sprintf("%d", int(num))
		}
		return fmt.Sprintf("%g", num)
	case STRING:
		s, _ := v.AsString()
		return fmt.Sprintf("\"%s\"", s)
	case BOOLEAN:
		b, _ := v.AsBoolean()
		return fmt.Sprintf("%t", b)
	case ARRAY:
		elements, _ := v.AsArray()
		strings := make([]string, len(elements))
		for i, elem := range elements {
			strings[i] = elem.String()
		}
		return "[" + strings.Join(strings, ", ") + "]"
	case FUNCTION:
		f, _ := v.AsFunction()
		return f.String()
	default:
		return "<unknown>"
	}
}

// IsTruthy returns whether the value is truthy
func (v *Value) IsTruthy() bool {
	switch v.Type {
	case NULL:
		return false
	case BOOLEAN:
		return v.AsBoolean()
	case NUMBER:
		return v.AsNumber() != 0
	case STRING:
		return v.AsString() != ""
	case ARRAY:
		return len(v.AsArray()) > 0
	case FUNCTION:
		return true
	default:
		return false
	}
}

// Function represents a function in the interpreter
type Function struct {
	Name       string
	Parameters []string
	Body       interface{} // This would be an ast.Block in a real implementation
	Closure    *Environment
}

// String returns a string representation of the function
func (f *Function) String() string {
	return fmt.Sprintf("<function %s>", f.Name)
}
