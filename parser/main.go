package main

import (
	"fmt"

	"antlr-editor/parser/core"
)

func main() {
	fmt.Println("ANTLR Editor Parser - Expression Validator")

	// Example usage
	validator := core.NewValidator()

	examples := []string{
		"1 + 2",
		"[price] * [quantity]",
		"SUM(x, y, z)",
		"(a == 5) && (b != 0)",
		"5 ++ 3", // Invalid: double operators
	}

	for _, expr := range examples {
		isValid := validator.Validate(expr)
		status := "✓"
		if !isValid {
			status = "✗"
		}
		fmt.Printf("%s %s\n", status, expr)
	}
}
