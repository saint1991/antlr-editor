package app

import (
	"encoding/json"
	"testing"

	"antlr-editor/parser/core/models"
)

func TestAnalyzer_Validate(t *testing.T) {
	analyzer := newAnalyzer()

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
			result := analyzer.Validate(tt.expression)
			if result != tt.expected {
				t.Errorf("Validate(%q) = %v, want %v", tt.expression, result, tt.expected)
			}
		})
	}
}

func TestAnalyzer_SimpleExpression(t *testing.T) {
	analyzer := newAnalyzer()
	result := analyzer.Analyze("SUM([price] * [quantity]) > 1000")

	if len(result.Errors) != 0 {
		t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
	}

	// Check that we have tokens
	if len(result.Tokens) == 0 {
		t.Errorf("Expected tokens, got none")
	}

	// Verify some key token types
	foundFunction := false
	foundColumnReference := false
	foundOperator := false
	foundParen := false

	for _, token := range result.Tokens {
		switch token.Type {
		case models.TokenFunction:
			foundFunction = true
			if token.Text != "SUM" {
				t.Errorf("Expected function token 'SUM', got '%s'", token.Text)
			}
		case models.TokenColumnReference:
			foundColumnReference = true
		case models.TokenOperator:
			foundOperator = true
		case models.TokenLeftParen, models.TokenRightParen:
			foundParen = true
		}
	}

	if !foundFunction {
		t.Errorf("Expected to find function token")
	}
	if !foundColumnReference {
		t.Errorf("Expected to find column reference token")
	}
	if !foundOperator {
		t.Errorf("Expected to find operator token")
	}
	if !foundParen {
		t.Errorf("Expected to find parenthesis token")
	}
}

func TestAnalyzer_ComplexExpression(t *testing.T) {
	analyzer := newAnalyzer()
	result := analyzer.Analyze("SUM([price] * [quantity]) > 1000 && [status] == 'active'")

	if len(result.Errors) != 0 {
		t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
	}

	// Check for string literal
	foundString := false
	for _, token := range result.Tokens {
		if token.Type == models.TokenString && token.Text == "'active'" {
			foundString = true
			break
		}
	}

	if !foundString {
		t.Errorf("Expected to find string literal 'active'")
	}
}

func TestAnalyzer_LiteralTypes(t *testing.T) {
	testCases := []struct {
		expression   string
		expectedType models.TokenType
		expectedText string
	}{
		{"123", models.TokenInteger, "123"},
		{"123.45", models.TokenFloat, "123.45"},
		{"true", models.TokenBoolean, "true"},
		{"false", models.TokenBoolean, "false"},
		{"'hello'", models.TokenString, "'hello'"},
		{`"world"`, models.TokenString, `"world"`},
	}

	analyzer := newAnalyzer()

	for _, tc := range testCases {
		t.Run(tc.expression, func(t *testing.T) {
			result := analyzer.Analyze(tc.expression)

			if len(result.Errors) != 0 {
				t.Errorf("Expected valid expression for %s, got %d errors", tc.expression, len(result.Errors))
			}

			found := false
			for _, token := range result.Tokens {
				if token.Type == tc.expectedType && token.Text == tc.expectedText {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Expected to find token type %s with text '%s' in expression '%s'", tc.expectedType, tc.expectedText, tc.expression)
			}
		})
	}
}

func TestAnalyzer_ErrorDetection(t *testing.T) {
	analyzer := newAnalyzer()
	result := analyzer.Analyze("SUM([price] * ") // Incomplete expression

	if len(result.Errors) == 0 {
		t.Errorf("Expected invalid expression, got valid")
	}

	if len(result.Errors) == 0 {
		t.Errorf("Expected errors, got none")
	}

	// Should still have tokens
	if len(result.Tokens) == 0 {
		t.Errorf("Expected tokens even with errors")
	}
}

func TestAnalyzer_EmptyExpression(t *testing.T) {
	analyzer := newAnalyzer()
	result := analyzer.Analyze("")

	// Empty expression should be invalid
	if len(result.Errors) == 0 {
		t.Errorf("Expected invalid empty expression, got no errors")
	}

	if len(result.Tokens) != 0 {
		t.Errorf("Expected no tokens for empty expression, got %d", len(result.Tokens))
	}
}

func TestAnalyzer_WhitespaceHandling(t *testing.T) {
	analyzer := newAnalyzer()
	result := analyzer.Analyze("SUM( [price] )")

	if len(result.Errors) != 0 {
		t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
	}

	// Check for whitespace tokens
	foundWhitespace := false
	for _, token := range result.Tokens {
		if token.Type == models.TokenWhitespace {
			foundWhitespace = true
			break
		}
	}

	if !foundWhitespace {
		t.Errorf("Expected to find whitespace tokens")
	}
}

func TestAnalyzer_PositionAccuracy(t *testing.T) {
	analyzer := newAnalyzer()
	expression := "SUM([price])"
	result := analyzer.Analyze(expression)

	if len(result.Errors) != 0 {
		t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
	}

	// Check position accuracy for some tokens
	for _, token := range result.Tokens {
		if token.Type != models.TokenEOF && token.End <= token.Start {
			t.Errorf("Invalid token position: start=%d, end=%d for token type %s", token.Start, token.End, token.Type)
		}

		if token.Start < 0 || token.End > len(expression) {
			t.Errorf("Token position out of bounds: start=%d, end=%d, expression length=%d", token.Start, token.End, len(expression))
		}

		// Verify token text matches the actual text at that position
		if token.Type != models.TokenWhitespace && token.Type != models.TokenEOF {
			actualText := expression[token.Start:token.End]
			if actualText != token.Text {
				t.Errorf("Token text mismatch: expected '%s', got '%s' at position %d-%d", actualText, token.Text, token.Start, token.End)
			}
		}
	}
}

func TestAnalyzer_MultilineExpression(t *testing.T) {
	analyzer := newAnalyzer()
	expression := "SUM([price])\n> 1000"
	result := analyzer.Analyze(expression)

	if len(result.Errors) != 0 {
		t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
	}

	// Check line numbers
	foundMultipleLines := false
	for _, token := range result.Tokens {
		if token.Line > 1 {
			foundMultipleLines = true
			break
		}
	}

	if !foundMultipleLines {
		t.Errorf("Expected to find tokens on multiple lines")
	}
}

func TestAnalyzer_JSONSerialization(t *testing.T) {
	analyzer := newAnalyzer()
	result := analyzer.Analyze("SUM([price]) > 1000")

	// Test that the result can be serialized to JSON
	jsonBytes, err := json.Marshal(result)
	if err != nil {
		t.Errorf("Failed to serialize result to JSON: %v", err)
	}

	// Test that it can be deserialized back
	var deserializedResult AnalysisResult
	err = json.Unmarshal(jsonBytes, &deserializedResult)
	if err != nil {
		t.Errorf("Failed to deserialize result from JSON: %v", err)
	}

	// Basic validation that deserialized data matches
	if len(deserializedResult.Errors) != len(result.Errors) {
		t.Errorf("Error count mismatch after JSON round-trip")
	}

	if len(deserializedResult.Tokens) != len(result.Tokens) {
		t.Errorf("Token count mismatch after JSON round-trip")
	}
}

func TestAnalyzer_OperatorTypes(t *testing.T) {
	operators := []string{"+", "-", "*", "/", "^", "<", "<=", ">", ">=", "==", "!=", "&&", "||"}
	analyzer := newAnalyzer()

	for _, op := range operators {
		t.Run(op, func(t *testing.T) {
			expression := "1 " + op + " 2"
			result := analyzer.Analyze(expression)

			foundOperator := false
			for _, token := range result.Tokens {
				if token.Type == models.TokenOperator && token.Text == op {
					foundOperator = true
					break
				}
			}

			if !foundOperator {
				t.Errorf("Expected to find operator '%s' in expression '%s'", op, expression)
			}
		})
	}
}

func TestAnalyzer_DelimiterTypes(t *testing.T) {
	delimiters := []string{"(", ")", "[", "]", ","}
	analyzer := newAnalyzer()

	expression := "SUM([price], [quantity])"
	result := analyzer.Analyze(expression)

	if len(result.Errors) != 0 {
		t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
	}

	for _, delim := range delimiters {
		foundDelimiter := false
		for _, token := range result.Tokens {
			// Check for specific delimiter token types
			var expectedType models.TokenType
			switch delim {
			case "(":
				expectedType = models.TokenLeftParen
			case ")":
				expectedType = models.TokenRightParen
			case "[":
				expectedType = models.TokenLeftBracket
			case "]":
				expectedType = models.TokenRightBracket
			case ",":
				expectedType = models.TokenComma
			}

			if token.Type == expectedType && token.Text == delim {
				foundDelimiter = true
				break
			}
		}

		if !foundDelimiter {
			t.Errorf("Expected to find delimiter '%s' in expression '%s'", delim, expression)
		}
	}
}

func TestAnalyzer_ErrorRecovery(t *testing.T) {
	analyzer := newAnalyzer()

	testCases := []struct {
		name         string
		expression   string
		expectErrors bool
	}{
		{"Unclosed parenthesis", "SUM([price]", true},
		{"Invalid operator", "1 ++ 2", true},
		{"Missing operand", "1 +", true},
		{"Invalid character", "1 @ 2", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := analyzer.Analyze(tc.expression)

			if tc.expectErrors {
				if len(result.Errors) == 0 {
					t.Errorf("Expected errors for expression '%s', got none", tc.expression)
				}
			}

			// Should still produce tokens even with errors
			if len(result.Tokens) == 0 {
				t.Errorf("Expected some tokens even with errors in expression '%s'", tc.expression)
			}
		})
	}
}

func TestAnalyzer_ErrorPosition(t *testing.T) {
	analyzer := newAnalyzer()
	result := analyzer.Analyze("1 ++ 2") // Invalid double plus

	if len(result.Errors) == 0 {
		t.Errorf("Expected errors for invalid expression")
	}

	for _, err := range result.Errors {
		if err.Line < 1 {
			t.Errorf("Invalid error line: %d", err.Line)
		}
		if err.Column < 0 {
			t.Errorf("Invalid error column: %d", err.Column)
		}
		if err.Start < 0 || err.End < err.Start {
			t.Errorf("Invalid error position: start=%d, end=%d", err.Start, err.End)
		}
	}
}

// Benchmark test for performance
func BenchmarkAnalyzer_SimpleExpression(b *testing.B) {
	analyzer := newAnalyzer()
	expression := "SUM([price] * [quantity]) > 1000 && [status] == 'active'"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.Analyze(expression)
	}
}

func BenchmarkAnalyzer_ComplexExpression(b *testing.B) {
	analyzer := newAnalyzer()
	expression := "((SUM([revenue]) - SUM([cost])) / SUM([revenue])) * 100 > 15.5 && [region] == 'North' || ([year] >= 2020 && [month] IN ('Jan', 'Feb', 'Mar'))"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.Analyze(expression)
	}
}

// Test the integrated validation functionality
func TestAnalyzer_ValidationIntegration(t *testing.T) {
	analyzer := newAnalyzer()

	tests := []struct {
		name       string
		expression string
		wantValid  bool
	}{
		{
			name:       "valid simple expression",
			expression: "1 + 2",
			wantValid:  true,
		},
		{
			name:       "valid function call",
			expression: "SUM([price])",
			wantValid:  true,
		},
		{
			name:       "valid complex expression",
			expression: "SUM([price] * [quantity]) > 1000",
			wantValid:  true,
		},
		{
			name:       "invalid syntax",
			expression: "1 + + 2",
			wantValid:  false,
		},
		{
			name:       "unclosed parenthesis",
			expression: "(1 + 2",
			wantValid:  false,
		},
		{
			name:       "empty expression",
			expression: "",
			wantValid:  false,
		},
		{
			name:       "incomplete expression",
			expression: "1 +",
			wantValid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test via Analyze method and IsValid()
			result := analyzer.Analyze(tt.expression)
			if result.IsValid() != tt.wantValid {
				t.Errorf("Analyze().IsValid() = %v, want %v", result.IsValid(), tt.wantValid)
			}

			// Test via direct Validate method
			isValid := analyzer.Validate(tt.expression)
			if isValid != tt.wantValid {
				t.Errorf("Validate() = %v, want %v", isValid, tt.wantValid)
			}

			// Both methods should return the same result
			if result.IsValid() != isValid {
				t.Errorf("Analyze().IsValid() = %v, but Validate() = %v", result.IsValid(), isValid)
			}
		})
	}
}

func TestAnalysisResult_IsValid(t *testing.T) {
	tests := []struct {
		name   string
		errors []models.ErrorInfo
		want   bool
	}{
		{
			name:   "no errors",
			errors: []models.ErrorInfo{},
			want:   true,
		},
		{
			name: "with errors",
			errors: []models.ErrorInfo{
				{Message: "syntax error", Line: 1, Column: 5},
			},
			want: false,
		},
		{
			name: "multiple errors",
			errors: []models.ErrorInfo{
				{Message: "syntax error", Line: 1, Column: 5},
				{Message: "semantic error", Line: 1, Column: 10},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := &AnalysisResult{
				Tokens: []models.TokenInfo{},
				Errors: tt.errors,
			}
			if got := result.IsValid(); got != tt.want {
				t.Errorf("AnalysisResult.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
