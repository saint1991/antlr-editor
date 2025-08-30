package main

/*
#include <stdint.h>
#include <stdlib.h>

#include "struct/analyzer.h"
#include "struct/error.h"
#include "struct/token.h"
*/
import "C"
import (
	"unsafe"

	"antlr-editor/analyzer/core/app"
	"antlr-editor/analyzer/core/models"
)

// TokenType to C enum mapping
var toCTokenType = map[models.TokenType]C.enum_TokenType{
	models.TokenString:          C.TOKEN_TYPE_STRING,
	models.TokenInteger:         C.TOKEN_TYPE_INTEGER,
	models.TokenFloat:           C.TOKEN_TYPE_FLOAT,
	models.TokenBoolean:         C.TOKEN_TYPE_BOOLEAN,
	models.TokenColumnReference: C.TOKEN_TYPE_COLUMN_REFERENCE,
	models.TokenFunction:        C.TOKEN_TYPE_FUNCTION,
	models.TokenOperator:        C.TOKEN_TYPE_OPERATOR,
	models.TokenComma:           C.TOKEN_TYPE_COMMA,
	models.TokenLeftParen:       C.TOKEN_TYPE_LEFT_PAREN,
	models.TokenRightParen:      C.TOKEN_TYPE_RIGHT_PAREN,
	models.TokenLeftBracket:     C.TOKEN_TYPE_LEFT_BRACKET,
	models.TokenRightBracket:    C.TOKEN_TYPE_RIGHT_BRACKET,
	models.TokenWhitespace:      C.TOKEN_TYPE_WHITESPACE,
	models.TokenError:           C.TOKEN_TYPE_ERROR,
	models.TokenEOF:             C.TOKEN_TYPE_EOF,
}

// Free allocated C strings in CTokenInfo
func freeCTokenInfo(token *C.CTokenInfo) {
	if token.text != nil {
		C.free(unsafe.Pointer(token.text))
	}
}

// Free allocated C strings in CErrorInfo
func freeCErrorInfo(err *C.CErrorInfo) {
	if err.message != nil {
		C.free(unsafe.Pointer(err.message))
	}
}

// Convert Go TokenInfo to C struct
func ToCTokenInfo(token models.TokenInfo) C.CTokenInfo {
	return C.CTokenInfo{
		token_type: toCTokenType[token.Type],
		text:       C.CString(token.Text),
		start:      C.int32_t(token.Start),
		end:        C.int32_t(token.End),
		line:       C.int32_t(token.Line),
		column:     C.int32_t(token.Column),
	}
}

// Convert Go ErrorInfo to C struct
func ToCErrorInfo(err models.ErrorInfo) C.CErrorInfo {
	return C.CErrorInfo{
		message: C.CString(err.Message),
		line:    C.int32_t(err.Line),
		column:  C.int32_t(err.Column),
		start:   C.int32_t(err.Start),
		end:     C.int32_t(err.End),
	}
}

// Global instances for FFI usage
var analyzer = app.NewApp()

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

	if analyzer.Validate(expressionStr) {
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

	if analyzer.Validate(expressionStr) {
		return 1
	}
	return 0
}

func tokenize(expression string) *C.CTokenizeResult {
	// Get analysis result
	result := analyzer.Analyze(expression)

	// Allocate C struct
	cResult := (*C.CTokenizeResult)(C.malloc(C.sizeof_CTokenizeResult))
	if cResult == nil {
		return nil
	}

	// Convert tokens
	if len(result.Tokens) > 0 {
		cResult.token_count = C.int32_t(len(result.Tokens))
		cResult.tokens = (*C.CTokenInfo)(C.malloc(C.size_t(len(result.Tokens)) * C.sizeof_CTokenInfo))

		// Copy each token
		tokens := (*[1 << 30]C.CTokenInfo)(unsafe.Pointer(cResult.tokens))[:len(result.Tokens):len(result.Tokens)]
		for i, token := range result.Tokens {
			tokens[i] = ToCTokenInfo(token)
		}
	} else {
		cResult.token_count = 0
		cResult.tokens = nil
	}

	// Convert errors
	if len(result.Errors) > 0 {
		cResult.error_count = C.int32_t(len(result.Errors))
		cResult.errors = (*C.CErrorInfo)(C.malloc(C.size_t(len(result.Errors)) * C.sizeof_CErrorInfo))

		// Copy each error
		errors := (*[1 << 30]C.CErrorInfo)(unsafe.Pointer(cResult.errors))[:len(result.Errors):len(result.Errors)]
		for i, err := range result.Errors {
			errors[i] = ToCErrorInfo(err)
		}
	} else {
		cResult.error_count = 0
		cResult.errors = nil
	}

	return cResult
}

// TokenizeFFI tokenizes expression and returns TokenizeResult struct
// The caller is responsible for freeing the returned struct using FreeTokenizeResult
//
//export TokenizeFFI
func TokenizeFFI(expression *C.char, length C.int) *C.CTokenizeResult {
	if expression == nil {
		return nil
	}

	// Convert C string to Go string
	expressionStr := C.GoStringN(expression, length)

	return tokenize(expressionStr)
}

// TokenizeFFIString tokenizes a null-terminated C string expression
// and returns TokenizeResult struct
// The caller is responsible for freeing the returned struct using FreeTokenizeResult
//
//export TokenizeFFIString
func TokenizeFFIString(expression *C.char) *C.CTokenizeResult {
	if expression == nil {
		return nil
	}

	// Convert C string to Go string
	expressionStr := C.GoString(expression)

	return tokenize(expressionStr)
}

// FreeTokenizeResult frees the memory allocated by TokenizeFFI
//
//export FreeTokenizeResult
func FreeTokenizeResult(result *C.CTokenizeResult) {
	if result == nil {
		return
	}

	// Free tokens
	if result.tokens != nil && result.token_count > 0 {
		tokens := (*[1 << 30]C.CTokenInfo)(unsafe.Pointer(result.tokens))[:result.token_count:result.token_count]
		for i := range tokens {
			freeCTokenInfo(&tokens[i])
		}
		C.free(unsafe.Pointer(result.tokens))
	}

	// Free errors
	if result.errors != nil && result.error_count > 0 {
		errors := (*[1 << 30]C.CErrorInfo)(unsafe.Pointer(result.errors))[:result.error_count:result.error_count]
		for i := range errors {
			freeCErrorInfo(&errors[i])
		}
		C.free(unsafe.Pointer(result.errors))
	}

	// Free the result struct itself
	C.free(unsafe.Pointer(result))
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

// FormatFFI formats expression and returns formatted string
// The caller is responsible for freeing the returned string using FreeString
//
//export FormatFFI
func FormatFFI(expression *C.char, length C.int) *C.char {
	if expression == nil {
		return nil
	}

	// Convert C string to Go string
	expressionStr := C.GoStringN(expression, length)

	// Format the expression
	formatted := analyzer.Format(expressionStr)

	// Return C string (caller must free)
	return C.CString(formatted)
}

// FormatFFIString formats a null-terminated C string expression
// and returns formatted string
// The caller is responsible for freeing the returned string using FreeString
//
//export FormatFFIString
func FormatFFIString(expression *C.char) *C.char {
	if expression == nil {
		return nil
	}

	// Convert C string to Go string
	expressionStr := C.GoString(expression)

	// Format the expression
	formatted := analyzer.Format(expressionStr)

	// Return C string (caller must free)
	return C.CString(formatted)
}

// main function required for FFI builds
func main() {
	// This is a no-op main function for FFI builds
	// The actual functionality is exposed through ValidateFFI functions
}
