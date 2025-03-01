package utils

import "reflect"

// IsNil determines whether the provided value is nil.
//
// This function is capable of handling both untyped nil and typed nil cases.
// It uses the `reflect` package to differentiate between nil and non-nil
// reference types.
//
// Supported types for nil checks include:
//   - Channels (`chan`)
//   - Functions (`func`)
//   - Interfaces (`interface{}`)
//   - Maps (`map`)
//   - Pointers (`*Type`)
//   - Slices (`[]Type`)
//
// For non-reference types (e.g., int, float, struct), this function always returns `false`
// because such types cannot be nil in Go.
//
// Parameters:
//   - value: The value to be checked. It can be of any type.
//
// Returns:
//   - `true`: If the value is nil (typed or untyped).
//   - `false`: If the value is not nil or belongs to a type that cannot be nil.
func IsNil(value any) bool {
	// If the value is untyped nil, return true.
	if value == nil {
		return true
	}

	// Use reflection to analyze the dynamic type and value.
	v := reflect.ValueOf(value)

	// If the value is not valid in reflection terms, it cannot be nil.
	if !v.IsValid() {
		return false
	}

	// Check if the value is nil for reference types.
	// Non-reference types cannot be nil at all.
	switch v.Kind() {
	case reflect.Chan, // Channels
		reflect.Func,      // Functions
		reflect.Interface, // Interfaces
		reflect.Map,       // Maps
		reflect.Ptr,       // Pointers
		reflect.Slice:     // Slices
		// Use IsNil to check if the reflected value is nil.
		return v.IsNil()
	default:
		// For non-reference types, return false as they cannot be nil.
		return false
	}
}
