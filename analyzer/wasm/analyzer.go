//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"

	"antlr-editor/analyzer/core/app"
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
	result := analyzer.Analyze(expression)

	return js.ValueOf(result.AsMap())
}

// format function exposed to JavaScript
func format(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return js.ValueOf("")
	}

	expression := args[0].String()
	formatted := analyzer.Format(expression)

	return js.ValueOf(formatted)
}

// formatWithOptions function exposed to JavaScript
func formatWithOptions(this js.Value, args []js.Value) any {
	if len(args) != 2 {
		return js.ValueOf("")
	}

	expression := args[0].String()
	optionsJS := args[1]

	// Extract options from JavaScript object
	options := app.DefaultFormatOptions()

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
		if uppercaseFunctions := optionsJS.Get("uppercaseFunctions"); !uppercaseFunctions.IsUndefined() {
			options = options.WithUppercaseFunctions(uppercaseFunctions.Bool())
		}
		if removeUnnecessaryParens := optionsJS.Get("removeUnnecessaryParens"); !removeUnnecessaryParens.IsUndefined() {
			options = options.WithRemoveUnnecessaryParens(removeUnnecessaryParens.Bool())
		}
		if breakLongExpressions := optionsJS.Get("breakLongExpressions"); !breakLongExpressions.IsUndefined() {
			options = options.WithBreakLongExpressions(breakLongExpressions.Bool())
		}
		if alignOperators := optionsJS.Get("alignOperators"); !alignOperators.IsUndefined() {
			options = options.WithAlignOperators(alignOperators.Bool())
		}
	}

	formatted := analyzer.FormatWithOptions(expression, options)
	return js.ValueOf(formatted)
}

// main function registers WASM functions and keeps the program running
func main() {
	// Register functions
	js.Global().Set("validateExpression", js.FuncOf(validate))
	js.Global().Set("analyzeExpression", js.FuncOf(analyze))
	js.Global().Set("formatExpression", js.FuncOf(format))
	js.Global().Set("formatExpressionWithOptions", js.FuncOf(formatWithOptions))

	// Keep the Go program running
	select {}
}
