package core

import (
	"encoding/json"
	"testing"

	"antlr-editor/parser/core/models"
)

func TestAnalyzer_SimpleExpression(t *testing.T) {
	analyzer := NewAnalyzer()
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
	analyzer := NewAnalyzer()
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

	analyzer := NewAnalyzer()

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
	analyzer := NewAnalyzer()
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
	analyzer := NewAnalyzer()
	result := analyzer.Analyze("")

	// Empty expression should be valid (no syntax errors)
	if len(result.Errors) != 0 {
		t.Errorf("Expected valid empty expression, got %d errors", len(result.Errors))
	}

	if len(result.Tokens) != 0 {
		t.Errorf("Expected no tokens for empty expression, got %d", len(result.Tokens))
	}
}

func TestAnalyzer_WhitespaceHandling(t *testing.T) {
	analyzer := NewAnalyzer()
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
	analyzer := NewAnalyzer()
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
	analyzer := NewAnalyzer()
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
	analyzer := NewAnalyzer()
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
	analyzer := NewAnalyzer()

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
	analyzer := NewAnalyzer()

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
	analyzer := NewAnalyzer()

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
	analyzer := NewAnalyzer()
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
	analyzer := NewAnalyzer()
	expression := "SUM([price] * [quantity]) > 1000 && [status] == 'active'"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.Analyze(expression)
	}
}

func BenchmarkAnalyzer_ComplexExpression(b *testing.B) {
	analyzer := NewAnalyzer()
	expression := "((SUM([revenue]) - SUM([cost])) / SUM([revenue])) * 100 > 15.5 && [region] == 'North' || ([year] >= 2020 && [month] IN ('Jan', 'Feb', 'Mar'))"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		analyzer.Analyze(expression)
	}
}
