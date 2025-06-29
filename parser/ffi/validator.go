package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"encoding/json"
	"unsafe"

	"antlr-editor/parser/core"
	"antlr-editor/parser/core/models"
)

// Global instances for FFI usage
var validator = core.NewValidator()
var analyzer = core.NewAnalyzer()

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

// AnalyzeFFI is an FFI-compatible wrapper for the Analyze function
// Returns JSON string with token information for syntax highlighting
// The caller is responsible for freeing the returned string using FreeString
//
//export AnalyzeFFI
func AnalyzeFFI(expression *C.char) *C.char {
	if expression == nil {
		return nil
	}

	// Convert C string to Go string
	expressionStr := C.GoString(expression)

	result := analyzer.Analyze(expressionStr)

	jsonBytes, err := json.Marshal(result)
	if err != nil {
		// Return error result in JSON format
		errorResult := core.AnalysisResult{
			Tokens: []models.TokenInfo{},
			Errors: []models.ErrorInfo{
				{
					Message: "Failed to serialize analysis result",
					Line:    1,
					Column:  0,
					Start:   0,
					End:     len(expressionStr),
				},
			},
		}
		errorBytes, _ := json.Marshal(errorResult)
		return C.CString(string(errorBytes))
	}

	return C.CString(string(jsonBytes))
}

// FreeString frees a string allocated by the Go runtime
// This must be called to free strings returned by AnalyzeFFI
//
//export FreeString
func FreeString(s *C.char) {
	if s != nil {
		// Free the C string allocated by C.CString
		C.free(unsafe.Pointer(s))
	}
}

// main function required for FFI builds
func main() {
	// This is a no-op main function for FFI builds
	// The actual functionality is exposed through ValidateFFI functions
}
