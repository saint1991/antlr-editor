//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
	"testing"
)

func TestValidateExpression(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       bool
	}{
		{
			name:       "valid simple expression",
			expression: "1 + 2",
			want:       true,
		},
		{
			name:       "valid complex expression",
			expression: "(1 + 2) * 3 / 4",
			want:       true,
		},
		{
			name:       "valid function call",
			expression: "MAX(1, 2, 3)",
			want:       true,
		},
		{
			name:       "valid column reference",
			expression: "[column_a] > 5",
			want:       true,
		},
		{
			name:       "invalid expression - missing operand",
			expression: "1 +",
			want:       false,
		},
		{
			name:       "invalid expression - mismatched parentheses",
			expression: "(1 + 2))",
			want:       false,
		},
		{
			name:       "empty expression",
			expression: "",
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []js.Value{js.ValueOf(tt.expression)}
			result := validate(js.Value{}, args)

			if got := result.(js.Value).Bool(); got != tt.want {
				t.Errorf("validate(%q) = %v, want %v", tt.expression, got, tt.want)
			}
		})
	}
}

func TestTokenizeExpression(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		wantTokens  int
		wantErrors  int
		checkTokens bool
	}{
		{
			name:        "valid simple expression",
			expression:  "1 + 2",
			wantTokens:  6, // 1, whitespace, +, whitespace, 2, EOF
			wantErrors:  0,
			checkTokens: true,
		},
		{
			name:        "valid function call",
			expression:  "SUM(10, 20)",
			wantTokens:  8, // SUM, (, 10, ,, whitespace, 20, ), EOF
			wantErrors:  0,
			checkTokens: true,
		},
		{
			name:        "invalid expression",
			expression:  "1 + + 2",
			wantTokens:  8, // returns token even for invalid expressions
			wantErrors:  1,
			checkTokens: true,
		},
		{
			name:        "empty expression",
			expression:  "",
			wantTokens:  0,
			wantErrors:  1,
			checkTokens: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []js.Value{js.ValueOf(tt.expression)}
			result := tokenize(js.Value{}, args)

			resultMap := result.(js.Value)
			tokens := resultMap.Get("tokens")
			errors := resultMap.Get("errors")

			tokenLength := tokens.Length()
			errorLength := errors.Length()

			if tt.checkTokens && tokenLength != tt.wantTokens {
				t.Errorf("tokenize(%q) returned %d tokens, want %d", tt.expression, tokenLength, tt.wantTokens)
			}

			if errorLength != tt.wantErrors {
				t.Errorf("tokenize(%q) returned %d errors, want %d", tt.expression, errorLength, tt.wantErrors)
			}

			// Check token structure for valid expressions
			if tt.wantErrors == 0 && tokenLength > 0 {
				firstToken := tokens.Index(0)
				if !firstToken.Get("type").Truthy() {
					t.Error("Token missing 'type' field")
				}
				if !firstToken.Get("text").Truthy() {
					t.Error("Token missing 'text' field")
				}
				if !firstToken.Get("start").Truthy() && firstToken.Get("start").Int() != 0 {
					t.Error("Token missing 'start' field")
				}
			}
		})
	}
}

func TestFormatExpression(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       string
	}{
		{
			name:       "simple addition",
			expression: "1+2",
			want:       "1 + 2",
		},
		{
			name:       "already formatted expression",
			expression: "1 + 2",
			want:       "1 + 2",
		},
		{
			name:       "complex expression with parentheses",
			expression: "(1+2)*3",
			want:       "(1 + 2) * 3",
		},
		{
			name:       "nested parentheses",
			expression: "((1+2)*3)/4",
			want:       "((1 + 2) * 3) / 4",
		},
		{
			name:       "function call",
			expression: "MAX(1,2,3)",
			want:       "MAX(1, 2, 3)",
		},
		{
			name:       "nested function calls",
			expression: "SUM(MAX(1,2),MIN(3,4))",
			want:       "SUM(MAX(1, 2), MIN(3, 4))",
		},
		{
			name:       "column reference",
			expression: "[column_a]+[column_b]",
			want:       "[column_a] + [column_b]",
		},
		{
			name:       "comparison operators",
			expression: "[value]>10",
			want:       "[value] > 10",
		},
		{
			name:       "logical operators",
			expression: "[a]>5&&[b]<10",
			want:       "[a] > 5 && [b] < 10",
		},
		{
			name:       "mixed operators",
			expression: "[price]*1.1+[tax]",
			want:       "[price] * 1.1 + [tax]",
		},
		{
			name:       "string literal",
			expression: `"hello"+"world"`,
			want:       `"hello" + "world"`,
		},
		{
			name:       "empty expression",
			expression: "",
			want:       "",
		},
		{
			name:       "invalid expression",
			expression: "1 + + 2",
			want:       "1 + + 2", // formatter returns original string for invalid expressions
		},
		{
			name:       "expression with extra spaces",
			expression: "1    +    2    *    3",
			want:       "1 + 2 * 3",
		},
		{
			name:       "expression with tabs and newlines",
			expression: "1\t+\n2",
			want:       "1 + 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []js.Value{js.ValueOf(tt.expression)}
			result := format(js.Value{}, args)

			if got := result.(js.Value).String(); got != tt.want {
				t.Errorf("format(%q) = %q, want %q", tt.expression, got, tt.want)
			}
		})
	}
}

func TestFormatWithOptions(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		options    map[string]interface{}
		want       string
	}{
		{
			name:       "default options",
			expression: "1+2*3",
			options:    nil,
			want:       "1 + 2 * 3",
		},
		{
			name:       "no spaces around operators",
			expression: "1 + 2 * 3",
			options: map[string]interface{}{
				"spaceAroundOps": false,
			},
			want: "1+2*3",
		},
		{
			name:       "custom indent size",
			expression: "SUM(1, MAX(2, 3))",
			options: map[string]interface{}{
				"indentSize": 4,
			},
			want: "SUM(1, MAX(2, 3))",
		},
		{
			name:       "break long expressions",
			expression: `VERYLONGFUNCTIONNAME("parameter1", "parameter2", "parameter3", "parameter4", "parameter5")`,
			options: map[string]interface{}{
				"breakLongExpressions": true,
				"maxLineLength":        40,
			},
			want: `VERYLONGFUNCTIONNAME(
  "parameter1",
  "parameter2",
  "parameter3",
  "parameter4",
  "parameter5"
)`,
		},
		{
			name:       "complex expression with all options",
			expression: "[column_a]+[column_b]*[column_c]/[column_d]",
			options: map[string]interface{}{
				"spaceAroundOps":       false,
				"breakLongExpressions": false,
				"indentSize":           2,
				"maxLineLength":        80,
			},
			want: "[column_a]+[column_b]*[column_c]/[column_d]",
		},
		{
			name:       "nested function with line breaking",
			expression: `IF([condition], COMPLEXCALCULATION("a", "b", "c"), ANOTHERCALCULATION("d", "e", "f"))`,
			options: map[string]interface{}{
				"breakLongExpressions": true,
				"maxLineLength":        30,
				"indentSize":           2,
			},
			want: `IF(
  [condition],
  COMPLEXCALCULATION(
    "a",
    "b",
    "c"
  ),
  ANOTHERCALCULATION(
    "d",
    "e",
    "f"
  )
)`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var optionsValue js.Value
			if tt.options == nil {
				optionsValue = js.Null()
			} else {
				// Create a JavaScript object from the map
				obj := js.Global().Get("Object").New()
				for key, value := range tt.options {
					obj.Set(key, js.ValueOf(value))
				}
				optionsValue = obj
			}

			args := []js.Value{js.ValueOf(tt.expression), optionsValue}
			result := formatWithOptions(js.Value{}, args)

			if got := result.(js.Value).String(); got != tt.want {
				t.Errorf("formatWithOptions(%q, %v) = %q, want %q", tt.expression, tt.options, got, tt.want)
			}
		})
	}
}

func TestFormatInvalidArguments(t *testing.T) {
	t.Run("format with no arguments", func(t *testing.T) {
		args := []js.Value{}
		result := format(js.Value{}, args)

		if got := result.(js.Value).String(); got != "" {
			t.Errorf("format() with no args = %q, want empty string", got)
		}
	})

	t.Run("format with multiple arguments", func(t *testing.T) {
		args := []js.Value{js.ValueOf("1+2"), js.ValueOf("extra")}
		result := format(js.Value{}, args)

		// Should still work with just the first argument
		if got := result.(js.Value).String(); got != "1 + 2" {
			t.Errorf("format() with extra args = %q, want '1 + 2'", got)
		}
	})

	t.Run("formatWithOptions with no arguments", func(t *testing.T) {
		args := []js.Value{}
		result := formatWithOptions(js.Value{}, args)

		if got := result.(js.Value).String(); got != "" {
			t.Errorf("formatWithOptions() with no args = %q, want empty string", got)
		}
	})

	t.Run("formatWithOptions with one argument", func(t *testing.T) {
		args := []js.Value{js.ValueOf("1+2")}
		result := formatWithOptions(js.Value{}, args)

		if got := result.(js.Value).String(); got != "1 + 2" {
			t.Errorf("formatWithOptions() with one arg = %q, want empty string", got)
		}
	})

	t.Run("formatWithOptions with undefined options", func(t *testing.T) {
		args := []js.Value{js.ValueOf("1+2"), js.Undefined()}
		result := formatWithOptions(js.Value{}, args)

		// Should use default options
		if got := result.(js.Value).String(); got != "1 + 2" {
			t.Errorf("formatWithOptions() with undefined options = %q, want '1 + 2'", got)
		}
	})
}

func TestInvalidArguments(t *testing.T) {
	t.Run("validate with no arguments", func(t *testing.T) {
		args := []js.Value{}
		result := validate(js.Value{}, args)

		if got := result.(js.Value).Bool(); got != false {
			t.Errorf("validate() with no args = %v, want false", got)
		}
	})

	t.Run("tokenize with no arguments", func(t *testing.T) {
		args := []js.Value{}
		result := tokenize(js.Value{}, args)

		resultMap := result.(js.Value)
		errors := resultMap.Get("errors")

		if errors.Length() == 0 {
			t.Error("analyze() with no args should return errors")
		}

		firstError := errors.Index(0)
		if msg := firstError.Get("message").String(); msg != "Invalid arguments" {
			t.Errorf("Expected 'Invalid arguments' error, got %q", msg)
		}
	})
}
