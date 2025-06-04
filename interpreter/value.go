package interpreter

import (
	"fmt"
	"reflect"
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
	// V-specific types
	OPTION   // Option<T>
	MAP      // map[K]V
	STRUCT   // struct
	ENUM     // enum
	INTERFACE // interface
	MODULE   // module reference
	ANY      // any other type
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
		num := v.Value.(float64)
		// Check if the number is an integer
		if num == float64(int(num)) {
			return fmt.Sprintf("%d", int(num))
		}
		return fmt.Sprintf("%g", num)
	case STRING:
		return fmt.Sprintf("%q", v.Value.(string))
	case BOOLEAN:
		return fmt.Sprintf("%t", v.Value.(bool))
	case ARRAY:
		elements := v.Value.([]*Value)
		var builder strings.Builder
		builder.WriteString("[")
		for i, element := range elements {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(element.String())
		}
		builder.WriteString("]")
		return builder.String()
	case FUNCTION:
		function := v.Value.(*Function)
		return function.String()
	// V-specific types
	case OPTION:
		option, ok := v.Value.(*OptionValue)
		if !ok {
			return "<invalid option>"
		}
		if !option.HasValue {
			return "None"
		}
		return fmt.Sprintf("Some(%s)", option.Value.String())
	case MAP:
		mapVal := v.Value.(*MapValue)
		pairs := make([]string, 0, len(mapVal.Entries))
		for k, val := range mapVal.Entries {
			keyStr := fmt.Sprintf("%v", k)
			pairs = append(pairs, fmt.Sprintf("%s: %s", keyStr, val.String()))
		}
		return fmt.Sprintf("{%s}", strings.Join(pairs, ", "))
	case STRUCT:
		structVal := v.Value.(*StructValue)
		fields := make([]string, 0, len(structVal.Fields))
		for name, val := range structVal.Fields {
			fields = append(fields, fmt.Sprintf("%s: %s", name, val.String()))
		}
		return fmt.Sprintf("%s{%s}", structVal.Name, strings.Join(fields, ", "))
	case ENUM:
		enumVal := v.Value.(*EnumValue)
		if enumVal.NumericVal != 0 {
			return fmt.Sprintf("%s.%s(%g)", enumVal.EnumType, enumVal.Name, enumVal.NumericVal)
		}
		return fmt.Sprintf("%s.%s", enumVal.EnumType, enumVal.Name)
	case INTERFACE:
		interfaceVal := v.Value.(*InterfaceValue)
		return fmt.Sprintf("<interface %s>", interfaceVal.Name)
	case MODULE:
		moduleName, ok := v.Value.(string)
		if !ok {
			return "<invalid module>"
		}
		return fmt.Sprintf("<module %s>", moduleName)
	case ANY:
		return fmt.Sprintf("<any %v>", v.Value)
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
		return v.Value.(bool)
	case NUMBER:
		return v.Value.(float64) != 0
	case STRING:
		return v.Value.(string) != ""
	case ARRAY:
		return len(v.Value.([]*Value)) > 0
	// V-specific types
	case OPTION:
		option := v.Value.(*OptionValue)
		return option.HasValue // Some is truthy, None is falsy
	case MAP:
		mapVal := v.Value.(*MapValue)
		return len(mapVal.Entries) > 0
	case STRUCT:
		return true // structs are always truthy in V
	case ENUM:
		enumVal := v.Value.(*EnumValue)
		return enumVal.NumericVal != 0 // 0 is falsy, other values are truthy
	case INTERFACE:
		return true // interfaces are always truthy
	case MODULE:
		return true // modules are always truthy
	case ANY:
		// If we can determine the value's truthiness, use that; otherwise true
		switch val := v.Value.(type) {
		case bool:
			return val
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			// Using reflection to handle all numeric types
			return reflect.ValueOf(val).Float() != 0
		case string:
			return val != ""
		default:
			return true
		}
	default:
		return true
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

// V Language Specific Value Types

// OptionValue represents V's option type which can be Some(value) or None
type OptionValue struct {
	HasValue bool
	Value    *Value
}

// MapValue represents a V map with keys and values
type MapValue struct {
	KeyType   string
	ValueType string
	Entries   map[interface{}]*Value
}

// StructValue represents a V struct instance
type StructValue struct {
	Name   string
	Fields map[string]*Value
}

// EnumValue represents a V enum instance
type EnumValue struct {
	Name       string
	EnumType   string
	NumericVal float64 // If it has an associated value
}

// Method represents a function with a receiver (method)
type Method struct {
	Function     *Function
	ReceiverType string
}

// InterfaceValue represents a V interface
type InterfaceValue struct {
	Name    string
	Methods map[string]*Method
}

// NewOption creates a new option value with a value
func NewOption(value *Value) *Value {
	return &Value{
		Type: OPTION,
		Value: &OptionValue{
			HasValue: true,
			Value:    value,
		},
	}
}

// NewNone creates a new option value with no value (None)
func NewNone() *Value {
	return &Value{
		Type: OPTION,
		Value: &OptionValue{
			HasValue: false,
			Value:    nil,
		},
	}
}

// NewMap creates a new map value
func NewMap(keyType, valueType string) *Value {
	return &Value{
		Type: MAP,
		Value: &MapValue{
			KeyType:   keyType,
			ValueType: valueType,
			Entries:   make(map[interface{}]*Value),
		},
	}
}

// NewStruct creates a new struct value
func NewStruct(name string) *Value {
	return &Value{
		Type: STRUCT,
		Value: &StructValue{
			Name:   name,
			Fields: make(map[string]*Value),
		},
	}
}

// NewEnum creates a new enum value
func NewEnum(name, enumType string, value float64) *Value {
	return &Value{
		Type: ENUM,
		Value: &EnumValue{
			Name:       name,
			EnumType:   enumType,
			NumericVal: value,
		},
	}
}

// IsSome checks if an option has a value
func (v *Value) IsSome() bool {
	if v.Type != OPTION {
		return false
	}
	
	option, ok := v.Value.(*OptionValue)
	if !ok {
		return false
	}
	
	return option.HasValue
}

// IsNone checks if an option has no value
func (v *Value) IsNone() bool {
	if v.Type != OPTION {
		return false
	}
	
	option, ok := v.Value.(*OptionValue)
	if !ok {
		return false
	}
	
	return !option.HasValue
}

// GetOptionValue gets the value of an option
func (v *Value) GetOptionValue() (*Value, error) {
	if v.Type != OPTION {
		return nil, fmt.Errorf("not an option value")
	}
	
	option, ok := v.Value.(*OptionValue)
	if !ok {
		return nil, fmt.Errorf("invalid option value")
	}
	
	if !option.HasValue {
		return nil, fmt.Errorf("option has no value (None)")
	}
	
	return option.Value, nil
}

// IsMap checks if a value is a map
func (v *Value) IsMap() bool {
	return v.Type == MAP
}

// GetMapValue gets the value for a key in a map
func (v *Value) GetMapValue(key interface{}) (*Value, error) {
	if v.Type != MAP {
		return nil, fmt.Errorf("not a map value")
	}
	
	mpValue, ok := v.Value.(*MapValue)
	if !ok {
		return nil, fmt.Errorf("invalid map value")
	}
	
	result, exists := mpValue.Entries[key]
	if !exists {
		return nil, fmt.Errorf("key not found in map")
	}
	
	return result, nil
}

// SetMapValue sets a key-value pair in a map
func (v *Value) SetMapValue(key interface{}, value *Value) error {
	if v.Type != MAP {
		return fmt.Errorf("not a map value")
	}
	
	mpValue, ok := v.Value.(*MapValue)
	if !ok {
		return fmt.Errorf("invalid map value")
	}
	
	mpValue.Entries[key] = value
	return nil
}

// IsStruct checks if a value is a struct
func (v *Value) IsStruct() bool {
	return v.Type == STRUCT
}

// GetField gets a field from a struct
func (v *Value) GetField(field string) (*Value, error) {
	if v.Type != STRUCT {
		return nil, fmt.Errorf("not a struct value")
	}
	
	structValue, ok := v.Value.(*StructValue)
	if !ok {
		return nil, fmt.Errorf("invalid struct value")
	}
	
	result, exists := structValue.Fields[field]
	if !exists {
		return nil, fmt.Errorf("field '%s' not found in struct '%s'", field, structValue.Name)
	}
	
	return result, nil
}

// SetField sets a field in a struct
func (v *Value) SetField(field string, value *Value) error {
	if v.Type != STRUCT {
		return fmt.Errorf("not a struct value")
	}
	
	structValue, ok := v.Value.(*StructValue)
	if !ok {
		return fmt.Errorf("invalid struct value")
	}
	
	structValue.Fields[field] = value
	return nil
}

// IsEnum checks if a value is an enum
func (v *Value) IsEnum() bool {
	return v.Type == ENUM
}

// GetEnumValue gets the enum value
func (v *Value) GetEnumValue() (string, string, float64, error) {
	if v.Type != ENUM {
		return "", "", 0, fmt.Errorf("not an enum value")
	}
	
	enum, ok := v.Value.(*EnumValue)
	if !ok {
		return "", "", 0, fmt.Errorf("invalid enum value")
	}
	
	return enum.Name, enum.EnumType, enum.NumericVal, nil
}

// IsModule checks if a value is a module reference
func (v *Value) IsModule() bool {
	return v.Type == MODULE
}

// NewModule creates a new module reference value
func NewModule(name string) *Value {
	return &Value{
		Type: MODULE,
		Value: name,
	}
}

// GetModuleName gets the module name
func (v *Value) GetModuleName() (string, error) {
	if v.Type != MODULE {
		return "", fmt.Errorf("not a module reference")
	}
	
	name, ok := v.Value.(string)
	if !ok {
		return "", fmt.Errorf("invalid module reference")
	}
	
	return name, nil
}
