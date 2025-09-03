//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"

	"antlr-editor/analyzer/core/app"
	"antlr-editor/analyzer/core/app/formatter"
)

// Global instances for WASM usage
var analyzer = app.NewApp()


// parseTree function exposed to JavaScript
func parseTree(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return js.ValueOf(map[string]any{
			"tree":   nil,
			"errors": []any{
				map[string]any{
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
	result := analyzer.ParseTree(expression)

	return js.ValueOf(result.AsMap())
}

// validate function exposed to JavaScript
func validate(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return js.ValueOf(false)
	}

	expression := args[0].String()

	return js.ValueOf(analyzer.Validate(expression))
}


// lint function exposed to JavaScript
func lint(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return js.ValueOf([]any{
			map[string]any{
				"message": "Invalid arguments",
				"line":    -1,
				"column":  -1,
				"start":   -1,
				"end":     -1,
			},
		})
	}

	expression := args[0].String()
	errors := analyzer.Lint(expression)

	// Convert errors to JavaScript-compatible format
	jsErrors := make([]any, len(errors))
	for i, err := range errors {
		jsErrors[i] = map[string]any{
			"message": err.Message,
			"line":    err.Line,
			"column":  err.Column,
			"start":   err.Start,
			"end":     err.End,
		}
	}

	return js.ValueOf(jsErrors)
}

// tokenize function exposed to JavaScript
func tokenize(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return js.ValueOf(map[string]any{
			"tokens": []any{},
			"errors": []any{
				map[string]any{
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
	result := analyzer.Tokenize(expression)

	return js.ValueOf(result.AsMap())
}


// format function exposed to JavaScript
func format(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return js.ValueOf("")
	}

	expression := args[0].String()
	formatted := analyzer.Format(expression)

	return js.ValueOf(formatted)
}

// formatWithOptions function exposed to JavaScript
func formatWithOptions(this js.Value, args []js.Value) any {
	if len(args) <= 1 {
		return format(this, args)
	}

	expression := args[0].String()
	optionsJS := args[1]

	// Extract options from JavaScript object
	options := formatter.DefaultFormatOptions()

	if !optionsJS.IsNull() && !optionsJS.IsUndefined() {
		if indentSize := optionsJS.Get("indentSize"); !indentSize.IsUndefined() {
			options = options.WithIndentSize(indentSize.Int())
		}
		if maxLineLength := optionsJS.Get("maxLineLength"); !maxLineLength.IsUndefined() {
			options = options.WithMaxLineLength(maxLineLength.Int())
		}
		if spaceAroundOps := optionsJS.Get("spaceAroundOps"); !spaceAroundOps.IsUndefined() {
			options = options.WithSpaceAroundOps(spaceAroundOps.Bool())
		}
		if breakLongExpressions := optionsJS.Get("breakLongExpressions"); !breakLongExpressions.IsUndefined() {
			options = options.WithBreakLongExpressions(breakLongExpressions.Bool())
		}
	}

	formatted := analyzer.FormatWithOptions(expression, options)
	return js.ValueOf(formatted)
}


// main function registers WASM functions and keeps the program running
func main() {
	// Register functions
	js.Global().Set("parseTree", js.FuncOf(parseTree))
	js.Global().Set("lint", js.FuncOf(lint))
	js.Global().Set("tokenize", js.FuncOf(tokenize))
	js.Global().Set("validate", js.FuncOf(validate))
	js.Global().Set("format", js.FuncOf(format))
	js.Global().Set("formatWithOptions", js.FuncOf(formatWithOptions))
	

	// Keep the Go program running
	select {}
}
