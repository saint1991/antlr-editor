package main

import (
	"C"

	"antlr-editor/parser/core"
)

// Global validator instance for FFI usage
var validator = core.NewValidator()

// ValidateFFI is an FFI-compatible wrapper for the Validate function
// This can be called from Python using ctypes or other FFI systems
//
//export ValidateFFI
func ValidateFFI(expression *C.char, length C.int) C.int {
	if expression == nil || length <= 0 {
		return 0
	}

	// Convert C string to Go string
	expressionStr := C.GoStringN(expression, length)

	if validator.Validate(expressionStr) {
		return 1
	}
	return 0
}

// ValidateFFIString is a simpler FFI wrapper that takes a null-terminated C string
//
//export ValidateFFIString
func ValidateFFIString(expression *C.char) C.int {
	if expression == nil {
		return 0
	}

	// Convert C string to Go string
	expressionStr := C.GoString(expression)

	if validator.Validate(expressionStr) {
		return 1
	}
	return 0
}

// main function required for FFI builds
func main() {
	// This is a no-op main function for FFI builds
	// The actual functionality is exposed through ValidateFFI functions
}
