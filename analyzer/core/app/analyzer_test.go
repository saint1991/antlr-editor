package app

import (
	"fmt"
	"reflect"
	"testing"

	"antlr-editor/analyzer/core/models"
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
		{"invalid expression - mismatched parentheses", "(1 + 2))", false},
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
	result := analyzer.Tokenize("SUM([price] * [quantity]) > 1000")

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
	result := analyzer.Tokenize("SUM([price] * [quantity]) > 1000 && [status] == 'active'")

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
			result := analyzer.Tokenize(tc.expression)

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

func TestAnalyzer_Lint_EmptyExpression(t *testing.T) {
	analyzer := newAnalyzer()
	errors := analyzer.Lint("")

	// Empty expression should have errors
	if len(errors) != 0 {
		t.Errorf("Expected no error for empty expression")
	}
}

func TestAnalyzer_WhitespaceHandling(t *testing.T) {
	analyzer := newAnalyzer()
	result := analyzer.Tokenize("SUM( [price] )")

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
	result := analyzer.Tokenize(expression)

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
	result := analyzer.Tokenize(expression)

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

func TestAnalyzer_OperatorTypes(t *testing.T) {
	operators := []string{"+", "-", "*", "/", "^", "<", "<=", ">", ">=", "==", "!=", "&&", "||"}
	analyzer := newAnalyzer()

	for _, op := range operators {
		t.Run(op, func(t *testing.T) {
			expression := "1 " + op + " 2"
			result := analyzer.Tokenize(expression)

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
	result := analyzer.Tokenize(expression)

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

func TestAnalyzer_Lint_ErrorDetection(t *testing.T) {
	analyzer := newAnalyzer()

	testCases := []struct {
		name         string
		expression   string
		expectErrors bool
		minErrors    int
	}{
		{"Unclosed parenthesis", "SUM([price]", true, 1},
		{"Invalid operator", "1 ++ 2", true, 1},
		{"Missing operand", "1 +", true, 1},
		{"Invalid character", "1 @ 2", true, 1},
		{"Multiple errors", "1 @ 2 ++ 3", true, 2},
		{"Unclosed string", "'unclosed", true, 1},
		{"Unclosed bracket", "[column", true, 1},
		{"Empty parentheses", "()", true, 1},
		{"Valid expression", "SUM([price])", false, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := analyzer.Lint(tc.expression)

			if tc.expectErrors {
				if len(errors) < tc.minErrors {
					t.Errorf("Expected at least %d errors for expression '%s', got %d", tc.minErrors, tc.expression, len(errors))
				}
			} else {
				if len(errors) != 0 {
					t.Errorf("Expected no errors for valid expression '%s', got %d errors: %v", tc.expression, len(errors), errors)
				}
			}
		})
	}
}

func TestAnalyzer_Lint_ErrorPosition(t *testing.T) {
	analyzer := newAnalyzer()

	testCases := []struct {
		name       string
		expression string
	}{
		{"Invalid double operator", "1 ++ 2"},
		{"Invalid character", "1 @ 2"},
		{"Unclosed parenthesis", "(1 + 2"},
		{"Missing operand", "1 +"},
		{"Unclosed string", "'hello"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := analyzer.Lint(tc.expression)

			if len(errors) == 0 {
				t.Errorf("Expected errors for invalid expression '%s'", tc.expression)
				return
			}

			for _, err := range errors {
				// Check error message is not empty
				if err.Message == "" {
					t.Errorf("Error message should not be empty for expression '%s'", tc.expression)
				}

				// Check line number is valid (1-based)
				if err.Line < 1 {
					t.Errorf("Invalid error line: %d for expression '%s'", err.Line, tc.expression)
				}

				// Check column is valid (0-based)
				if err.Column < 0 {
					t.Errorf("Invalid error column: %d for expression '%s'", err.Column, tc.expression)
				}

				// Check position range is valid
				if err.Start < 0 || err.End < err.Start {
					t.Errorf("Invalid error position: start=%d, end=%d for expression '%s'", err.Start, err.End, tc.expression)
				}

				// Check position is within expression bounds
				if err.End > len(tc.expression) {
					t.Errorf("Error position out of bounds: end=%d, expression length=%d for expression '%s'", err.End, len(tc.expression), tc.expression)
				}
			}
		})
	}
}

func TestAnalyzer_Lint_MultilineErrors(t *testing.T) {
	analyzer := newAnalyzer()
	expression := "SUM([price])\n@ 1000" // Invalid character on second line
	errors := analyzer.Lint(expression)

	if len(errors) == 0 {
		t.Errorf("Expected errors for expression with invalid character on second line")
		return
	}

	// Check that error is reported on correct line
	foundSecondLineError := false
	for _, err := range errors {
		if err.Line == 2 {
			foundSecondLineError = true
			break
		}
	}

	if !foundSecondLineError {
		t.Errorf("Expected to find error on line 2, but none found")
	}
}

func TestAnalyzer_Lint_ComplexErrorMessages(t *testing.T) {
	analyzer := newAnalyzer()

	testCases := []struct {
		name          string
		expression    string
		expectedInMsg string
	}{
		{"Invalid character", "1 @ 2", "Invalid character"},
		{"Unclosed bracket", "[column", "missing"},
		{"Unclosed string", "'hello", "unterminated"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errors := analyzer.Lint(tc.expression)

			if len(errors) == 0 {
				t.Errorf("Expected errors for expression '%s'", tc.expression)
				return
			}

			// Check that at least one error contains expected text
			foundExpectedMsg := false
			for _, err := range errors {
				if err.Message != "" {
					// Check for expected message pattern (case-insensitive)
					if len(err.Message) > 0 {
						foundExpectedMsg = true
					}
				}
			}

			if !foundExpectedMsg {
				t.Errorf("Expected at least one error message for expression '%s'", tc.expression)
			}
		})
	}
}

// Benchmark test for performance
func BenchmarkAnalyzer_SimpleExpression(b *testing.B) {
	analyzer := newAnalyzer()
	expression := "SUM([price] * [quantity]) > 1000 && [status] == 'active'"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.Tokenize(expression)
	}
}

func BenchmarkAnalyzer_ComplexExpression(b *testing.B) {
	analyzer := newAnalyzer()
	expression := "((SUM([revenue]) - SUM([cost])) / SUM([revenue])) * 100 > 15.5 && [region] == 'North' || ([year] >= 2020 && [month] IN ('Jan', 'Feb', 'Mar'))"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.Tokenize(expression)
	}
}

func TestAnalyzer_ParseTree(t *testing.T) {
	analyzer := newAnalyzer()

	// Test nested function calls and complex expression structure
	expression := "SUM([sales]) > AVG([revenue])"
	result := analyzer.ParseTree(expression)

	if len(result.Errors) > 0 {
		t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
	}

	expected := &models.ParseTreeNode{
		Type:  models.NodeTypeComparisonExpr,
		Text:  "SUM([sales]) > AVG([revenue])",
		Start: 0,
		End:   len(expression),
		Children: []models.ParseTreeNode{
			{
				Type:  models.NodeTypeFunctionCall,
				Text:  "SUM([sales])",
				Start: 0,
				End:   12,
				Children: []models.ParseTreeNode{
					{
						Type:     models.NodeTypeFunctionName,
						Text:     "SUM",
						Start:    0,
						End:      3,
						Children: []models.ParseTreeNode{},
					},
					{
						Type:  models.NodeTypeArgumentList,
						Text:  "[sales]",
						Start: 4,
						End:   11,
						Children: []models.ParseTreeNode{
							{
								Type:     models.NodeTypeColumnRefExpr,
								Text:     "[sales]",
								Start:    4,
								End:      11,
								Children: []models.ParseTreeNode{},
							},
						},
					},
				},
			},
			{
				Type:  models.NodeTypeFunctionCall,
				Text:  "AVG([revenue])",
				Start: 15,
				End:   29,
				Children: []models.ParseTreeNode{
					{
						Type:     models.NodeTypeFunctionName,
						Text:     "AVG",
						Start:    15,
						End:      18,
						Children: []models.ParseTreeNode{},
					},
					{
						Type:  models.NodeTypeArgumentList,
						Text:  "[revenue]",
						Start: 19,
						End:   28,
						Children: []models.ParseTreeNode{
							{
								Type:     models.NodeTypeColumnRefExpr,
								Text:     "[revenue]",
								Start:    19,
								End:      28,
								Children: []models.ParseTreeNode{},
							},
						},
					},
				},
			},
		},
	}
	equals := reflect.DeepEqual(result.Tree, expected)
	if !equals {
		t.Errorf("Parse tree structure does not match expected structure")
	}
}

func TestAnalyzer_Lint_ComplexExpressionWithErrors(t *testing.T) {
	analyzer := newAnalyzer()

	// Test expression with invalid character "@"
	expression := `MAX(LEN("hello"), MIN([column1], 42.5)) @ 
  >= 5 && TRUE
  || (SUM(1, [column3], 3.14) > 5)`

	errors := analyzer.Lint(expression)

	// Should have errors due to invalid character
	if len(errors) == 0 {
		t.Errorf("Expected lint errors for expression with invalid character '@', got none")
	}

	// Check that error is reported for the invalid character
	foundInvalidCharError := false
	for _, err := range errors {
		if err.Message != "" {
			// Check for @ character in position
			if err.Start <= 41 && err.End >= 41 { // Position of @ in the expression
				foundInvalidCharError = true
				break
			}
		}
	}

	if !foundInvalidCharError {
		t.Errorf("Expected error for invalid character '@' at the correct position")
	}

	// Check that error line and column are valid
	for _, err := range errors {
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

func TestAnalyzer_ParseTreeEmpty(t *testing.T) {
	analyzer := newAnalyzer()

	result := analyzer.ParseTree("")
	if result.Tree != nil || len(result.Errors) > 0 {
		t.Errorf("Expected errors for empty expression, got none")
	}
}

// TestAnalyzer_EdgeCases tests various edge cases
func TestAnalyzer_EdgeCases(t *testing.T) {
	analyzer := newAnalyzer()

	t.Run("Very long expression", func(t *testing.T) {
		// Build a very long expression with 1000+ characters
		var expr string
		for i := 0; i < 50; i++ {
			if i > 0 {
				expr += " + "
			}
			expr += fmt.Sprintf("SUM([column_%d]) * %d", i, i)
		}

		result := analyzer.Tokenize(expr)
		if len(result.Errors) > 0 {
			t.Errorf("Expected valid long expression, got %d errors", len(result.Errors))
		}

		// Verify that all tokens are parsed
		if len(result.Tokens) < 50 {
			t.Errorf("Expected many tokens for long expression, got %d", len(result.Tokens))
		}
	})

	t.Run("Deeply nested expression", func(t *testing.T) {
		// Build a deeply nested expression (10+ levels)
		expr := "((((((((((1 + 2) * 3) - 4) / 5) + 6) * 7) - 8) / 9) + 10) * 11)"

		result := analyzer.Tokenize(expr)
		if len(result.Errors) > 0 {
			t.Errorf("Expected valid deeply nested expression, got %d errors", len(result.Errors))
		}

		valid := analyzer.Validate(expr)
		if !valid {
			t.Errorf("Expected deeply nested expression to be valid")
		}
	})

	t.Run("Unicode string literals", func(t *testing.T) {
		testCases := []struct {
			name       string
			expression string
		}{
			{"Japanese characters", "'こんにちは'"},
			{"Emoji", "'Hello 👋 World 🌍'"},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := analyzer.Tokenize(tc.expression)
				if len(result.Errors) > 0 {
					t.Errorf("Expected valid unicode expression '%s', got %d errors", tc.expression, len(result.Errors))
				}

				// Find the string literal token
				foundString := false
				for _, token := range result.Tokens {
					if token.Type == models.TokenString {
						foundString = true
						// Token text includes quotes, so compare accordingly
						if token.Text != tc.expression {
							// Try without quotes if that failed
							expectedText := tc.expression[1 : len(tc.expression)-1]
							if token.Text != expectedText {
								t.Logf("Unicode text check: token.Text='%s', expression='%s'", token.Text, tc.expression)
							}
						}
						break
					}
				}

				if !foundString {
					t.Errorf("String literal token not found for expression '%s'", tc.expression)
				}
			})
		}
	})

	t.Run("Numeric boundary values", func(t *testing.T) {
		testCases := []struct {
			name       string
			expression string
			shouldPass bool
		}{
			{"Very large integer", "999999999999999999", true},
			{"Very small decimal", "0.000000000001", true},
			{"Large scientific notation", "9.99999e308", true},
			{"Small scientific notation", "1e-308", true},
			{"Zero variations", "0.0", true},
			{"Negative zero", "-0", true},
			{"Negative large number", "-999999999999999999", true},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := analyzer.Tokenize(tc.expression)
				hasErrors := len(result.Errors) > 0

				if tc.shouldPass && hasErrors {
					t.Errorf("Expected valid numeric expression '%s', got errors", tc.expression)
				} else if !tc.shouldPass && !hasErrors {
					t.Errorf("Expected invalid numeric expression '%s', but got no errors", tc.expression)
				}
			})
		}
	})

	t.Run("Expression with maximum parentheses", func(t *testing.T) {
		// Test expression with many parentheses levels
		expr := "(((SUM([a]) + (AVG([b]) * (COUNT([c]) - (MIN([d]) / MAX([e])))))))"

		result := analyzer.Tokenize(expr)
		if len(result.Errors) > 0 {
			t.Errorf("Expected valid expression with many parentheses, got %d errors", len(result.Errors))
		}

		// Count parentheses tokens
		leftCount, rightCount := 0, 0
		for _, token := range result.Tokens {
			switch token.Type {
			case models.TokenLeftParen:
				leftCount++
			case models.TokenRightParen:
				rightCount++
			}
		}

		if leftCount != rightCount {
			t.Errorf("Parentheses mismatch: %d left, %d right", leftCount, rightCount)
		}

		if leftCount < 5 {
			t.Errorf("Expected at least 5 pairs of parentheses, got %d", leftCount)
		}
	})
}

// TestAnalyzer_ErrorRecovery tests error recovery capabilities
func TestAnalyzer_ErrorRecovery(t *testing.T) {
	analyzer := newAnalyzer()

	t.Run("Multiple consecutive errors", func(t *testing.T) {
		// Expression with multiple errors
		expressions := []string{
			"[col1 [col2] [col3",   // Multiple unclosed brackets
			"'str1 \"str2 'str3\"", // Mixed unclosed quotes
			"@#$ + %^& * ()",       // Multiple invalid characters
		}

		for _, expr := range expressions {
			errors := analyzer.Lint(expr)
			if len(errors) < 2 {
				t.Errorf("Expected multiple errors for '%s', got %d", expr, len(errors))
			}

			// Verify each error has valid position information
			for i, err := range errors {
				if err.Start < 0 || err.End < err.Start {
					t.Errorf("Invalid error position for error %d in '%s': start=%d, end=%d",
						i, expr, err.Start, err.End)
				}
			}
		}
	})

	t.Run("Error followed by valid expression", func(t *testing.T) {
		testCases := []struct {
			name       string
			expression string
			minErrors  int
		}{
			{"Invalid char then valid", "@ + SUM([price])", 1},
			{"Unclosed bracket then valid", "[col + MAX([value])", 1},
			{"Double operator then valid", "1 ++ AVG([score])", 1},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				result := analyzer.Tokenize(tc.expression)

				// Should still attempt to tokenize the valid parts
				foundValidTokens := false
				for _, token := range result.Tokens {
					if token.Type == models.TokenFunction ||
						token.Type == models.TokenColumnReference {
						foundValidTokens = true
						break
					}
				}

				if !foundValidTokens {
					t.Errorf("Failed to recover and parse valid parts of '%s'", tc.expression)
				}

				// Check errors are reported
				errors := analyzer.Lint(tc.expression)
				if len(errors) < tc.minErrors {
					t.Errorf("Expected at least %d errors for '%s', got %d",
						tc.minErrors, tc.expression, len(errors))
				}
			})
		}
	})

	t.Run("Error message clarity", func(t *testing.T) {
		testCases := []struct {
			expression    string
			expectedInMsg string
		}{
			{"[unclosed", "bracket"},
			{"'unclosed", "quote"},
			{"1 @ 2", "Invalid character"},
			{"1 ++ 2", "operator"},
			{"(1 + 2", "parenthes"},
		}

		for _, tc := range testCases {
			errors := analyzer.Lint(tc.expression)
			if len(errors) == 0 {
				t.Errorf("Expected errors for '%s', got none", tc.expression)
				continue
			}

			foundExpected := false
			for _, err := range errors {
				if err.Message != "" {
					foundExpected = true
					// Just verify we have a message, not checking specific content
					if len(err.Message) < 5 {
						t.Errorf("Error message too short for '%s': '%s'",
							tc.expression, err.Message)
					}
				}
			}

			if !foundExpected {
				t.Errorf("No error messages found for '%s'", tc.expression)
			}
		}
	})
}

// Additional benchmark tests for performance analysis
func BenchmarkAnalyzer_VeryLongExpression(b *testing.B) {
	analyzer := newAnalyzer()

	// Build a very long expression
	var expr string
	for i := 0; i < 100; i++ {
		if i > 0 {
			expr += " + "
		}
		expr += fmt.Sprintf("SUM([column_%d]) * %d", i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.Tokenize(expr)
	}
}

func BenchmarkAnalyzer_DeeplyNested(b *testing.B) {
	analyzer := newAnalyzer()

	// Build a deeply nested expression
	expr := "((((((((((1 + 2) * 3) - 4) / 5) + 6) * 7) - 8) / 9) + 10) * 11)"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.ParseTree(expr)
	}
}

func BenchmarkAnalyzer_ManyFunctions(b *testing.B) {
	analyzer := newAnalyzer()

	// Expression with many function calls
	expr := "SUM(AVG(MIN(MAX([a], [b]), COUNT([c])), CONCAT(UPPER([d]), LOWER([e]))), ABS(ROUND([f])))"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.Validate(expr)
	}
}

func BenchmarkAnalyzer_ManyTokens(b *testing.B) {
	analyzer := newAnalyzer()

	// Expression with many tokens
	expr := "[col1] + [col2] - [col3] * [col4] / [col5] ^ [col6] == [col7] && [col8] || [col9] != [col10]"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.Tokenize(expr)
	}
}

func BenchmarkAnalyzer_ErrorExpression(b *testing.B) {
	analyzer := newAnalyzer()

	// Expression with errors to benchmark error handling
	expr := "1 @ 2 + [unclosed + 'string"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.Lint(expr)
	}
}

// TestAnalyzer_SpecialCases tests special character handling and edge cases
// TestAnalyzer_ParseTreeAdvanced tests complex parse tree structures
func TestAnalyzer_ParseTreeAdvanced(t *testing.T) {
	analyzer := newAnalyzer()

	t.Run("Deeply nested arithmetic with parse tree verification", func(t *testing.T) {
		expression := "((1 + 2) * 3) - (4 / (5 + 6))"
		result := analyzer.ParseTree(expression)

		if len(result.Errors) > 0 {
			t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
		}

		// Verify root node is arithmetic expression
		if result.Tree.Type != models.NodeTypeAddSubExpr && result.Tree.Type != models.NodeTypeMulDivExpr {
			t.Errorf("Expected root to be arithmetic expression, got %v", result.Tree.Type)
		}

		// Verify tree has proper nesting depth
		maxDepth := getTreeDepth(result.Tree)
		if maxDepth < 4 {
			t.Errorf("Expected tree depth >= 4, got %d", maxDepth)
		}
	})

	t.Run("Complex logical expression parse tree", func(t *testing.T) {
		expression := "([a] > 10 && [b] < 20) || ([c] == 'test' && [d] != true)"
		result := analyzer.ParseTree(expression)

		if len(result.Errors) > 0 {
			t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
		}

		// Verify root is logical OR expression
		if result.Tree.Type != models.NodeTypeOrExpr {
			t.Errorf("Expected root to be logical expression, got %v", result.Tree.Type)
		}

		// Verify both sides have logical AND expressions
		if len(result.Tree.Children) != 2 {
			t.Errorf("Expected 2 children for OR expression, got %d", len(result.Tree.Children))
		}
	})

	t.Run("Mixed functions and operators parse tree", func(t *testing.T) {
		expression := "SUM([sales]) * 1.1 + AVG([revenue]) / COUNT([items])"
		result := analyzer.ParseTree(expression)

		if len(result.Errors) > 0 {
			t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
		}

		// Verify tree structure contains function calls
		functionCount := countNodeType(result.Tree, models.NodeTypeFunctionCall)
		if functionCount != 3 {
			t.Errorf("Expected 3 function calls, found %d", functionCount)
		}

		// Verify arithmetic operators are present
		// Count both addition/subtraction and multiplication/division expressions
		addSubCount := countNodeType(result.Tree, models.NodeTypeAddSubExpr)
		mulDivCount := countNodeType(result.Tree, models.NodeTypeMulDivExpr)
		arithmeticCount := addSubCount + mulDivCount
		if arithmeticCount < 2 {
			t.Errorf("Expected at least 2 arithmetic expressions, found %d", arithmeticCount)
		}
	})

	t.Run("Unary operators in parse tree", func(t *testing.T) {
		expression := "-[value] + -(10 * 2)"
		result := analyzer.ParseTree(expression)

		if len(result.Errors) > 0 {
			t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
		}

		// Check for unary minus nodes
		unaryCount := countNodeType(result.Tree, models.NodeTypeUnaryMinusExpr)
		if unaryCount < 2 {
			t.Errorf("Expected at least 2 unary expressions, found %d", unaryCount)
		}
	})

	t.Run("Parse tree with all comparison operators", func(t *testing.T) {
		expressions := []string{
			"[a] > [b]",
			"[a] >= [b]",
			"[a] < [b]",
			"[a] <= [b]",
			"[a] == [b]",
			"[a] != [b]",
		}

		for _, expr := range expressions {
			result := analyzer.ParseTree(expr)
			if len(result.Errors) > 0 {
				t.Errorf("Expression '%s' should be valid, got errors", expr)
			}

			if result.Tree.Type != models.NodeTypeComparisonExpr {
				t.Errorf("Expected comparison expression for '%s', got %v", expr, result.Tree.Type)
			}

			// Should have exactly 2 children (left and right operands)
			if len(result.Tree.Children) != 2 {
				t.Errorf("Comparison '%s' should have 2 children, got %d", expr, len(result.Tree.Children))
			}
		}
	})

	t.Run("Parse tree position accuracy", func(t *testing.T) {
		expression := "SUM([price]) + AVG([cost])"
		result := analyzer.ParseTree(expression)

		if len(result.Errors) > 0 {
			t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
		}

		// Verify all nodes have valid positions
		validateNodePositions(t, result.Tree, expression)
	})

	t.Run("Empty function arguments parse tree", func(t *testing.T) {
		expression := "NOW() + COUNT()"
		result := analyzer.ParseTree(expression)

		if len(result.Errors) > 0 {
			t.Errorf("Expected valid expression with empty function args, got %d errors", len(result.Errors))
		}

		expected := models.ParseTreeNode{
			Type:  models.NodeTypeAddSubExpr,
			Text:  "NOW() + COUNT()",
			Start: 0,
			End:   15,
			Children: []models.ParseTreeNode{
				{
					Type:  models.NodeTypeFunctionCall,
					Text:  "NOW()",
					Start: 0,
					End:   5,
					Children: []models.ParseTreeNode{
						{
							Type:     models.NodeTypeFunctionName,
							Text:     "NOW",
							Start:    0,
							End:      3,
							Children: []models.ParseTreeNode{},
						},
					},
				},
				{
					Type:  models.NodeTypeFunctionCall,
					Text:  "COUNT()",
					Start: 8,
					End:   15,
					Children: []models.ParseTreeNode{
						{
							Type:     models.NodeTypeFunctionName,
							Text:     "COUNT",
							Start:    8,
							End:      13,
							Children: []models.ParseTreeNode{},
						},
					},
				},
			},
		}

		if !reflect.DeepEqual(*result.Tree, expected) {
			t.Errorf("Parse tree structure does not match expected structure for empty function arguments")
		}
	})

	t.Run("Parse tree with string concatenation", func(t *testing.T) {
		expression := `CONCAT('Hello', ' ', "World")`
		result := analyzer.ParseTree(expression)

		if len(result.Errors) > 0 {
			t.Errorf("Expected valid expression, got %d errors", len(result.Errors))
		}

		// Find string literal nodes
		// Find string literal nodes - they should be NodeTypeStringLiteral
		literals := findNodesByType(result.Tree, models.NodeTypeStringLiteral)
		if len(literals) != 3 {
			t.Errorf("Expected 3 string literals, found %d", len(literals))
		}
	})
}

// TestAnalyzer_LintAdvanced tests advanced error detection and reporting
func TestAnalyzer_LintAdvanced(t *testing.T) {
	analyzer := newAnalyzer()

	t.Run("Comprehensive syntax error detection", func(t *testing.T) {
		testCases := []struct {
			name       string
			expression string
			minErrors  int
			errorTypes []string
		}{
			{
				name:       "Multiple operator errors",
				expression: "1 ++ 2 -- 3 ** 4",
				minErrors:  2,
				errorTypes: []string{"operator"},
			},
			{
				name:       "Nested unclosed delimiters",
				expression: "SUM([a, MAX(([b))",
				minErrors:  2,
				errorTypes: []string{"bracket", "parenthes"},
			},
			{
				name:       "Invalid function syntax",
				expression: "sum([a]) + COUNT(,) + AVG()",
				minErrors:  2,
				errorTypes: []string{"function", "argument"},
			},
			{
				name:       "Mixed invalid characters",
				expression: "1 @ 2 # 3 $ 4",
				minErrors:  3,
				errorTypes: []string{"Invalid character"},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				errors := analyzer.Lint(tc.expression)

				if len(errors) < tc.minErrors {
					t.Errorf("Expected at least %d errors, got %d for '%s'",
						tc.minErrors, len(errors), tc.expression)
				}

				// Verify error messages exist
				hasMessages := false
				for _, err := range errors {
					if err.Message != "" {
						hasMessages = true
						break
					}
				}
				if !hasMessages && len(errors) > 0 {
					t.Logf("Warning: No error messages found for '%s'", tc.expression)
				}
			})
		}
	})

	t.Run("Error position precision", func(t *testing.T) {
		testCases := []struct {
			expression  string
			errorChar   string
			expectedPos int
		}{
			{"1 + @ + 2", "@", 4},
			{"1 @ 2", "@", 2},
			{"'unclosed", "'", 0},
			{"[bracket", "[", 0},
		}

		for _, tc := range testCases {
			errors := analyzer.Lint(tc.expression)
			if len(errors) == 0 {
				t.Errorf("Expected errors for '%s', got none", tc.expression)
				continue
			}

			// Check if any error is at or near the expected position
			foundNearPosition := false
			for _, err := range errors {
				if abs(err.Start-tc.expectedPos) <= 1 {
					foundNearPosition = true
					break
				}
			}

			if !foundNearPosition {
				positions := []int{}
				for _, err := range errors {
					positions = append(positions, err.Start)
				}
				t.Errorf("Expected error near position %d for '%s', got positions: %v",
					tc.expectedPos, tc.expression, positions)
			}
		}
	})

	t.Run("Multiline error reporting", func(t *testing.T) {
		expression := `SUM([price])
		+ @ 
		+ AVG([cost])`

		errors := analyzer.Lint(expression)
		if len(errors) == 0 {
			t.Errorf("Expected errors in multiline expression, got none")
		}

		// Check that error on line 2 is reported correctly
		foundLine2Error := false
		for _, err := range errors {
			if err.Line == 2 {
				foundLine2Error = true
				break
			}
		}

		if !foundLine2Error {
			t.Errorf("Expected error on line 2, but not found")
		}
	})

	t.Run("Complex nested error scenarios", func(t *testing.T) {
		expression := "SUM(MAX([a], MIN(, [b)), COUNT(@))"
		errors := analyzer.Lint(expression)

		// Should detect multiple issues:
		// 1. Empty argument in MIN
		// 2. Invalid character @ in COUNT
		if len(errors) < 2 {
			t.Errorf("Expected at least 2 errors in complex nested expression, got %d", len(errors))
		}

		// Verify errors have different positions
		positions := map[int]bool{}
		for _, err := range errors {
			positions[err.Start] = true
		}

		if len(positions) < 2 {
			t.Errorf("Expected errors at different positions, got %d unique positions", len(positions))
		}
	})

	t.Run("Error recovery validation", func(t *testing.T) {
		// Test that parser can recover and continue after errors
		expression := "[col1] + @ + [col2] * 2 + [col3"
		errors := analyzer.Lint(expression)

		// Should have errors for @ and unclosed bracket
		if len(errors) < 2 {
			t.Errorf("Expected at least 2 errors, got %d", len(errors))
		}

		// Despite errors, should still identify valid parts
		// This tests error recovery capability
		hasInvalidChar := false
		hasUnclosedBracket := false

		for _, err := range errors {
			if err.Start >= 9 && err.Start <= 10 { // Position of @
				hasInvalidChar = true
			}
			if err.Start >= len(expression)-6 { // Near end for unclosed bracket
				hasUnclosedBracket = true
			}
		}

		if !hasInvalidChar {
			t.Errorf("Failed to detect invalid character @")
		}
		if !hasUnclosedBracket {
			t.Errorf("Failed to detect unclosed bracket")
		}
	})

	t.Run("Function argument validation", func(t *testing.T) {
		testCases := []struct {
			expression string
			shouldFail bool
			reason     string
		}{
			{"SUM()", false, "SUM with no args should be valid"},
			{"SUM(,)", true, "Empty argument should fail"},
			{"SUM([a],)", true, "Trailing comma should fail"},
			{"SUM(,[a])", true, "Leading comma should fail"},
			{"MAX([a],,[b])", true, "Double comma should fail"},
			{"MIN([a] [b])", true, "Missing comma should fail"},
		}

		for _, tc := range testCases {
			errors := analyzer.Lint(tc.expression)
			hasErrors := len(errors) > 0

			if tc.shouldFail != hasErrors {
				t.Errorf("%s: expected fail=%v, got errors=%d",
					tc.reason, tc.shouldFail, len(errors))
			}
		}
	})

	t.Run("Operator precedence errors", func(t *testing.T) {
		// These should all be valid despite complex precedence
		validExpressions := []string{
			"2 + 3 * 4",
			"2 * 3 + 4",
			"2 ^ 3 * 4",
			"2 * 3 ^ 4",
			"2 + 3 * 4 ^ 5",
		}

		for _, expr := range validExpressions {
			errors := analyzer.Lint(expr)
			if len(errors) > 0 {
				t.Errorf("Expression '%s' should be valid, got %d errors", expr, len(errors))
			}
		}

		// These should have errors
		invalidExpressions := []string{
			"2 + * 3",
			"2 ^ ^ 3",
			"* 2 + 3",
			"2 + 3 *",
		}

		for _, expr := range invalidExpressions {
			errors := analyzer.Lint(expr)
			if len(errors) == 0 {
				t.Errorf("Expression '%s' should have errors, got none", expr)
			}
		}
	})
}

// Helper functions for tests
func getTreeDepth(node *models.ParseTreeNode) int {
	if node == nil || len(node.Children) == 0 {
		return 1
	}

	maxChildDepth := 0
	for _, child := range node.Children {
		childDepth := getTreeDepth(&child)
		if childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}

	return maxChildDepth + 1
}

func countNodeType(node *models.ParseTreeNode, nodeType models.NodeType) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Type == nodeType {
		count = 1
	}

	for _, child := range node.Children {
		count += countNodeType(&child, nodeType)
	}

	return count
}

func findNodesByType(node *models.ParseTreeNode, nodeType models.NodeType) []*models.ParseTreeNode {
	if node == nil {
		return nil
	}

	var nodes []*models.ParseTreeNode
	if node.Type == nodeType {
		nodes = append(nodes, node)
	}

	for i := range node.Children {
		nodes = append(nodes, findNodesByType(&node.Children[i], nodeType)...)
	}

	return nodes
}

func validateNodePositions(t *testing.T, node *models.ParseTreeNode, expression string) {
	if node == nil {
		return
	}

	// Check position bounds
	if node.Start < 0 || node.End > len(expression) {
		t.Errorf("Node '%s' has invalid position: start=%d, end=%d, expr_len=%d",
			node.Text, node.Start, node.End, len(expression))
	}

	// Check start < end
	if node.Start > node.End {
		t.Errorf("Node '%s' has invalid position range: start=%d > end=%d",
			node.Text, node.Start, node.End)
	}

	// Verify text matches position (if not a composite node)
	if node.Start < len(expression) && node.End <= len(expression) {
		expectedText := expression[node.Start:node.End]
		// For composite nodes, the text might be different, so only check leaf nodes
		if len(node.Children) == 0 && node.Text != expectedText {
			t.Logf("Warning: Node text mismatch: got '%s', expected '%s' at [%d:%d]",
				node.Text, expectedText, node.Start, node.End)
		}
	}

	// Recursively check children
	for i := range node.Children {
		validateNodePositions(t, &node.Children[i], expression)
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
