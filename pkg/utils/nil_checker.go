package utils

import "reflect"

// IsNil checks if the given value is nil.
//
// It handles both typed and untyped nil cases and uses reflection to
// differentiate between nil values and non-nil values for reference types.
//
// Supported types:
// - Channels
// - Functions
// - Interfaces
// - Maps
// - Pointers
// - Slices
//
// For non-reference types (e.g., int, float, struct, etc.), this function always
// returns false because such types cannot be nil in Go.
//
// Parameters:
// - value: The value to check. Can be of any type.
//
// Returns:
// - true if the value is nil (either typed or untyped nil).
// - false if the value is not nil or belongs to a type that cannot be nil.
func IsNil(value any) bool {
	// If the value is completely nil (no dynamic type, no value), return true.
	if value == nil {
		return true
	}

	// Obtain the reflection-based representation of the value.
	v := reflect.ValueOf(value)

	// Check if the reflected value is valid. Invalid values cannot be nil.
	if !v.IsValid() {
		return false
	}

	// Use reflection to check if the value is nil for reference types only.
	// Non-reference types (e.g., int, float, struct) cannot be nil.
	switch v.Kind() {
	case reflect.Chan, // Channels
		reflect.Func,      // Functions
		reflect.Interface, // Interfaces
		reflect.Map,       // Maps
		reflect.Ptr,       // Pointers
		reflect.Slice:     // Slices
		// For reference types, use IsNil() to check nil value.
		return v.IsNil()
	default:
		// Non-reference types cannot be nil, return false.
		return false
	}
}
