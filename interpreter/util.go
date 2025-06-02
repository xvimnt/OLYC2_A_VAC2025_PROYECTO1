package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseNumber parses a number from a string representation
func ParseNumber(s string) float64 {
	// Try to parse the number
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		// Return 0 if it's not a valid number
		return 0
	}
	return val
}

// ParseBool parses a boolean from a string representation
func ParseBool(s string) bool {
	// Convert to lowercase and trim quotes
	s = strings.ToLower(strings.Trim(s, "\""))
	
	// Check if it's a boolean value
	return s == "true"
}

// IsNumber checks if a value is a number
func IsNumber(val interface{}) bool {
	_, ok := val.(float64)
	return ok
}

// IsString checks if a value is a string
func IsString(val interface{}) bool {
	_, ok := val.(string)
	return ok
}

// IsBool checks if a value is a boolean
func IsBool(val interface{}) bool {
	_, ok := val.(bool)
	return ok
}

// IsArray checks if a value is an array
func IsArray(val interface{}) bool {
	_, ok := val.([]interface{})
	return ok
}

// ToString converts a value to its string representation
func ToString(val interface{}) string {
	switch v := val.(type) {
	case nil:
		return "null"
	case float64:
		// Convert to integer if it has no fractional part
		if v == float64(int(v)) {
			return fmt.Sprintf("%d", int(v))
		}
		return fmt.Sprintf("%g", v)
	case string:
		return fmt.Sprintf("\"%s\"", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case []interface{}:
		elements := make([]string, len(v))
		for i, elem := range v {
			elements[i] = ToString(elem)
		}
		return "[" + strings.Join(elements, ", ") + "]"
	default:
		return fmt.Sprintf("%v", v)
	}
}

// RuntimeError represents a runtime error during interpretation
type RuntimeError struct {
	Message string
}

// Error returns the error message
func (e *RuntimeError) Error() string {
	return e.Message
}
