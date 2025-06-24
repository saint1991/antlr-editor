package core

import (
	"testing"
)

func TestValidator_Validate(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name       string
		expression string
		expected   bool
	}{
		// Empty/invalid cases
		{"empty string", "", false},

		// Simple literals
		{"integer literal", "42", true},
		{"float literal", "3.14", true},
		{"string literal", `"hello"`, true},
		{"boolean true", "true", true},
		{"boolean false", "false", true},

		// Arithmetic expressions
		{"simple addition", "1 + 2", true},
		{"simple subtraction", "5 - 3", true},
		{"simple multiplication", "4 * 6", true},
		{"simple division", "8 / 2", true},
		{"power operation", "2 ^ 3", true},

		// Complex arithmetic
		{"multiple operations", "1 + 2 * 3", true},
		{"parentheses", "(1 + 2) * 3", true},
		{"nested parentheses", "((1 + 2) * 3) / 2", true},

		// Comparison operations
		{"equality", "5 == 5", true},
		{"inequality", "5 != 3", true},
		{"mixed comparison", "x == 5", true},

		// Logical operations
		{"logical and", "true && false", true},
		{"logical or", "true || false", true},
		{"complex logical", "(x == 5) && (y != 3)", true},

		// Column references
		{"simple column", "[column_name]", true},
		{"column with underscore", "[user_id]", true},
		{"column with numbers", "[col123]", true},

		// Function calls
		{"function no args", "SUM()", true},
		{"function with args", "MAX(a, b)", true},
		{"function with multiple args", "AVERAGE(x, y, z)", true},
		{"nested function calls", "MAX(MIN(a, b), c)", true},

		// Complex expressions
		{"arithmetic with columns", "[price] * [quantity]", true},
		{"function with arithmetic", "SUM([price] * [quantity])", true},
		{"mixed operations", "([price] * [quantity]) + TAX([price])", true},

		// Edge cases that should be valid
		{"negative number", "-5", true},
		{"decimal number", "0.5", true},
		{"string with spaces", `"hello world"`, true},

		// Invalid syntax cases
		{"unmatched parentheses", "(1 + 2", false},
		{"unmatched brackets", "[column", false},
		{"invalid operator", "5 # 3", false},
		{"double operators", "5 ++ 3", false},
		{"empty parentheses in expression", "5 + ()", false},
		{"invalid function syntax", "FUNC(,)", false},
		{"trailing operator", "5 +", false},
		{"leading operator", "* 5", false},
		{"empty brackets", "[]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.Validate(tt.expression)
			if result != tt.expected {
				t.Errorf("Validate(%q) = %v, want %v", tt.expression, result, tt.expected)
			}
		})
	}
}

func TestNewValidator(t *testing.T) {
	validator := NewValidator()
	if validator == nil {
		t.Error("NewValidator() returned nil")
	}
}

// Benchmark tests
func BenchmarkValidator_Validate(b *testing.B) {
	validator := NewValidator()
	expressions := []string{
		"1 + 2",
		"[price] * [quantity]",
		"SUM([price] * [quantity])",
		"(([price] * [quantity]) + TAX([price])) / COUNT([items])",
	}

	for _, expr := range expressions {
		b.Run(expr, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				validator.Validate(expr)
			}
		})
	}
}

func BenchmarkValidator_ValidateComplex(b *testing.B) {
	validator := NewValidator()
	complexExpr := "((([price] * [quantity]) + TAX([price])) / COUNT([items])) && (STATUS([order]) == \"completed\")"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator.Validate(complexExpr)
	}
}
