package main

import (
	"antlr-editor/parser/core"
)

// Global validator instance for WASM usage
var validator = core.NewValidator()

// ValidateWASM is a WASM-compatible wrapper for the Validate function
// This will be used when compiling to WebAssembly
//
//go:export validate
func ValidateWASM(expression string) int {
	if validator.Validate(expression) {
		return 1
	}
	return 0
}

// main function required for WASM builds
func main() {
	// This is a no-op main function for WASM builds
	// The actual functionality is exposed through ValidateWASM
}
