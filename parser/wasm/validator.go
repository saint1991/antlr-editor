package main

import (
	"encoding/json"

	"antlr-editor/parser/core"
	"antlr-editor/parser/core/models"
)

// Global instances for WASM usage
var validator = core.NewValidator()
var analyzer = core.NewAnalyzer()

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

// AnalyzeWASM is a WASM-compatible wrapper for the Analyze function
// Returns JSON string with token information for syntax highlighting
//
//go:export analyze
func AnalyzeWASM(expression string) string {
	result := analyzer.Analyze(expression)

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
					End:     len(expression),
				},
			},
		}
		errorBytes, _ := json.Marshal(errorResult)
		return string(errorBytes)
	}

	return string(jsonBytes)
}

// main function required for WASM builds
func main() {
	// This is a no-op main function for WASM builds
	// The actual functionality is exposed through ValidateWASM
}
