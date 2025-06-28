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
		// 1. Literal Tests

		// 1.1 String Literals
		{"single quoted string", "'Hello World'", true},
		{"double quoted string", `"Hello World"`, true},
		{"double quotes within single quotes", `'She said "Hello"'`, true},
		{"single quote within double quotes", `"It's a beautiful day"`, true},
		{"escaped single quote", `'It\'s escaped'`, true},
		{"escaped double quotes", `"She said \"Hello\""`, true},
		{"empty string single quotes", "''", true},
		{"empty string double quotes", `""`, true},

		// 1.2 Integer Literals
		{"zero", "0", true},
		{"positive integer", "123", true},
		{"large integer", "999999", true},

		// 1.3 Float Literals
		{"basic decimal", "3.14", true},
		{"decimal starting with zero", "0.5", true},
		{"large decimal", "123.456", true},
		{"scientific notation positive exponent", "1.23e4", true},
		{"scientific notation negative exponent", "1.23e-4", true},
		{"scientific notation uppercase E positive", "1.23E+4", true},
		{"scientific notation integer part only", "2e10", true},

		// 1.4 Boolean Literals
		{"boolean true", "true", true},
		{"boolean false", "false", true},

		// 2. Column Reference Tests
		{"basic identifier column", "[name]", true},
		{"column with underscore", "[user_id]", true},
		{"camelCase column", "[firstName]", true},
		{"column starting with underscore", "[_private]", true},
		{"column with numbers", "[column123]", true},
		{"single character column", "[a]", true},
		{"long column name", "[averylongcolumnnamethatshouldstillwork]", true},

		// 3. Function Call Tests
		{"function no arguments", "NOW()", true},
		{"function no arguments COUNT", "COUNT()", true},
		{"function with column argument", "SUM([price])", true},
		{"function with literal argument", "ABS(-5)", true},
		{"function with string argument", "UPPER('hello')", true},
		{"function with two arguments", "MAX([score1], [score2])", true},
		{"function with three arguments", "CONCAT('Hello', ' ', 'World')", true},
		{"function with three columns", "ADD([a], [b], [c])", true},
		{"nested functions", "SUM(MAX([a], [b]))", true},
		{"multiple nested functions", "CONCAT(UPPER([first]), LOWER([last]))", true},

		// 4. Arithmetic Operation Tests
		{"addition", "1 + 2", true},
		{"subtraction", "5 - 3", true},
		{"multiplication", "4 * 6", true},
		{"division", "8 / 2", true},
		{"exponentiation", "2 ^ 3", true},
		{"column multiplication", "[price] * [quantity]", true},
		{"column and literal addition", "[total] + 100", true},
		{"column exponentiation", "[base] ^ 2", true},

		// 5. Precedence Tests
		{"multiplication precedence", "2 + 3 * 4", true},
		{"parentheses precedence", "(2 + 3) * 4", true},
		{"exponentiation first then multiply", "2 ^ 3 * 4", true},
		{"multiply then exponentiation", "2 * 3 ^ 4", true},
		{"right associative exponentiation", "2 ^ 3 ^ 4", true},

		// 6. Unary Minus Tests
		{"negative literal", "-5", true},
		{"negative column reference", "-[value]", true},
		{"negative parenthesized expression", "-(2 + 3)", true},

		// 7. Comparison Operation Tests
		{"column equality", "[age] == 25", true},
		{"column inequality", "[status] != 'inactive'", true},
		{"string equality", "'hello' == 'hello'", true},

		// 8. Logical Operation Tests
		{"basic AND", "true && false", true},
		{"column AND", "[active] && [verified]", true},
		{"comparison result AND", "[age] == 25 && [status] == 'active'", true},
		{"basic OR", "true || false", true},
		{"column OR", "[premium] || [vip]", true},
		{"comparison result OR", "[type] == 'A' || [type] == 'B'", true},
		{"AND precedence test", "[a] && [b] || [c]", true},
		{"OR AND precedence", "[a] || [b] && [c]", true},
		{"parentheses change precedence", "([a] || [b]) && [c]", true},

		// 9. Complex Expression Tests
		{"complex arithmetic comparison", "[price] * [quantity] > 1000 && [status] == 'active'", true},
		{"complex with parentheses", "([subtotal] + [tax]) * [discount] != 0", true},
		{"function arithmetic comparison", "SUM([values]) / COUNT([values]) > [threshold]", true},
		{"string function comparison", "CONCAT([first], ' ', [last]) == 'John Doe'", true},
		{"multiple function arithmetic", "MAX([a], [b]) + MIN([c], [d])", true},
		{"deeply nested arithmetic", "((([a] + [b]) * [c]) / [d]) ^ 2", true},
		{"complex logical expression", "([x] && [y]) || ([z] && [w])", true},
		{"complex function arithmetic", "SUM([a]) + SUM([b]) * COUNT([c])", true},

		// 10. Parentheses Tests
		{"simple parentheses", "(1 + 2)", true},
		{"nested parentheses", "((1 + 2) * 3)", true},
		{"multiple parenthesis groups", "(([a] + [b]) * ([c] - [d]))", true},
		{"single value parentheses", "([value])", true},
		{"function parentheses", "(SUM([values]))", true},

		// 11. Error Cases (Should cause parse errors)
		{"incomplete expression", "1 +", false},
		{"unclosed bracket", "[unclosed", false},
		{"lowercase function name", "function()", false},
		{"incomplete comparison", "[col] ==", false},
		{"invalid number", "123abc", false},
		{"unclosed string", "'unclosed string", false},
		{"spaces in column name", "[col with spaces]", false},
		{"consecutive operators", "1 + + 2", false},

		// Additional test cases from the original file
		{"empty string", "", false},
		{"unmatched parentheses", "(1 + 2", false},
		{"invalid operator", "5 # 3", false},
		{"double operators", "5 ++ 3", false},
		{"empty parentheses in expression", "5 + ()", false},
		{"invalid function syntax", "FUNC(,)", false},
		{"trailing operator", "5 +", false},
		{"leading operator", "* 5", false},
		{"empty brackets", "[]", false},

		// Additional complex cases
		{"function with expression argument", "MAX(1 + 2, 3 * 4)", true},
		{"comparison with greater than", "[price] > 100", true},
		{"comparison with less than", "[age] < 65", true},
		{"comparison with greater or equal", "[score] >= 80", true},
		{"comparison with less or equal", "[discount] <= 0.5", true},
		{"mixed operators", "([a] + [b]) * [c] == [d] - [e]", true},
		{"function in comparison", "SUM([values]) > 1000", true},
		{"negative float", "-3.14", true},
		{"negative in expression", "[total] + -100", true},
		{"complex nested functions", "ROUND(AVG(SUM([values]), COUNT([items])))", true},
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

// Benchmark deeply nested expressions for performance testing
func BenchmarkValidator_ValidateDeeplyNested(b *testing.B) {
	validator := NewValidator()
	// Create a deeply nested expression
	deepExpr := "((((((((([a] + [b]) * [c]) - [d]) / [e]) ^ [f]) + [g]) * [h]) - [i]) / [j])"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		validator.Validate(deepExpr)
	}
}

// Test specific operator precedence and associativity
func TestValidator_OperatorPrecedence(t *testing.T) {
	validator := NewValidator()

	precedenceTests := []struct {
		name       string
		expression string
		expected   bool
		comment    string
	}{
		// Arithmetic precedence
		{"multiply before add", "2 + 3 * 4", true, "Should parse as 2 + (3 * 4)"},
		{"power before multiply", "2 * 3 ^ 4", true, "Should parse as 2 * (3 ^ 4)"},
		{"power right associative", "2 ^ 3 ^ 4", true, "Should parse as 2 ^ (3 ^ 4)"},

		// Logical precedence
		{"AND before OR", "[a] || [b] && [c]", true, "Should parse as [a] || ([b] && [c])"},
		{"comparison before logical", "[a] > 5 && [b] < 10", true, "Should parse as ([a] > 5) && ([b] < 10)"},

		// Mixed precedence
		{"arithmetic before comparison", "[a] + [b] > [c] * [d]", true, "Should parse as ([a] + [b]) > ([c] * [d])"},
		{"comparison before logical", "[a] > 5 || [b] < 3 && [c] == 7", true, "Complex precedence"},
	}

	for _, tt := range precedenceTests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.Validate(tt.expression)
			if result != tt.expected {
				t.Errorf("Validate(%q) = %v, want %v // %s", tt.expression, result, tt.expected, tt.comment)
			}
		})
	}
}

// Test edge cases for string escaping
func TestValidator_StringEscaping(t *testing.T) {
	validator := NewValidator()

	escapeTests := []struct {
		name       string
		expression string
		expected   bool
	}{
		{"escaped backslash", `'test\\path'`, true},
		{"escaped newline", `'line1\nline2'`, true},
		{"escaped tab", `'col1\tcol2'`, true},
		{"mixed escapes", `'Quote: \', Tab: \t, Newline: \n'`, true},
		{"double quote escaped backslash", `"test\\path"`, true},
		{"unicode escape", `'\u0041'`, true}, // May depend on grammar support
	}

	for _, tt := range escapeTests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.Validate(tt.expression)
			if result != tt.expected {
				t.Errorf("Validate(%q) = %v, want %v", tt.expression, result, tt.expected)
			}
		})
	}
}

// Test function name constraints (uppercase only)
func TestValidator_FunctionNameConstraints(t *testing.T) {
	validator := NewValidator()

	functionTests := []struct {
		name       string
		expression string
		expected   bool
	}{
		{"uppercase function", "SUM([a])", true},
		{"all caps function", "COUNT([b])", true},
		{"lowercase function", "sum([a])", false},
		{"mixed case function", "Sum([a])", false},
		{"function with numbers", "SUM2([a])", false},
		{"function with underscore", "SUM_VALUES([a])", false},
	}

	for _, tt := range functionTests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.Validate(tt.expression)
			if result != tt.expected {
				t.Errorf("Validate(%q) = %v, want %v", tt.expression, result, tt.expected)
			}
		})
	}
}

// Test identifier constraints
func TestValidator_IdentifierConstraints(t *testing.T) {
	validator := NewValidator()

	identifierTests := []struct {
		name       string
		expression string
		expected   bool
	}{
		{"starts with letter", "[column]", true},
		{"starts with underscore", "[_private]", true},
		{"contains numbers", "[col123]", true},
		{"contains underscore", "[first_name]", true},
		{"starts with number", "[123col]", false},
		{"contains space", "[col name]", false},
		{"contains special char", "[col-name]", false},
		{"contains dot", "[table.column]", false},
	}

	for _, tt := range identifierTests {
		t.Run(tt.name, func(t *testing.T) {
			result := validator.Validate(tt.expression)
			if result != tt.expected {
				t.Errorf("Validate(%q) = %v, want %v", tt.expression, result, tt.expected)
			}
		})
	}
}
