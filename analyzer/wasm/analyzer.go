//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"

	"antlr-editor/parser/core/app"
)

// Global instances for WASM usage
var analyzer = app.NewApp()

// validate function exposed to JavaScript
func validate(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return js.ValueOf(false)
	}

	expression := args[0].String()

	return js.ValueOf(analyzer.Validate(expression))
}

// analyze function exposed to JavaScript
func analyze(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return js.ValueOf(map[string]any{
			"tokens": []any{},
			"errors": []any {
				map[string]any {
					"message": "Invalid arguments",
					"line":    -1,
					"column":  -1,
					"start":   -1,
					"end":     -1,
				},
			},
		})
	}

	expression := args[0].String()
	result := analyzer.Analyze(expression)

	return js.ValueOf(result.AsMap())
}

// main function registers WASM functions and keeps the program running
func main() {
	// Register functions
	js.Global().Set("validateExpression", js.FuncOf(validate))
	js.Global().Set("analyzeExpression", js.FuncOf(analyze))

	// Keep the Go program running
	select {}
}
