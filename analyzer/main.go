package main

import (
	"fmt"

	"antlr-editor/analyzer/core/app"
)

func main() {
	fmt.Println("ANTLR Editor Parser - Expression Validator")

	// Example usage
	application := app.NewApp()

	examples := []string{
		"1 + 2",
		"[price] * [quantity]",
		"SUM(x, y, z)",
		"(a == 5) && (b != 0)",
		"5 ++ 3",    // Invalid: double operators
		"123abc",    // Invalid: error characters
		"hello@#$%", // Invalid: multiple error characters
	}

	for _, expr := range examples {
		isValid := application.Validate(expr)
		status := "✓"
		if !isValid {
			status = "✗"
		}
		fmt.Printf("%s %s\n", status, expr)

		// Show details for invalid expressions containing error characters
		if expr == "123abc" || expr == "hello@#$%" {
			result := application.Analyze(expr)
			fmt.Printf("  Tokens: %d\n", len(result.Tokens))
			for _, tok := range result.Tokens {
				fmt.Printf("    '%s' (type: %s, pos: %d-%d)\n",
					tok.Text, tok.Type, tok.Start, tok.End)
			}
		}
	}
}
