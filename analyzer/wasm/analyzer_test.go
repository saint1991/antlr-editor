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

func TestAnalyzeExpression(t *testing.T) {
	tests := []struct {
		name          string
		expression    string
		wantTokens    int
		wantErrors    int
		checkTokens   bool
	}{
		{
			name:          "valid simple expression",
			expression:    "1 + 2",
			wantTokens:    6, // 1, whitespace, +, whitespace, 2, EOF
			wantErrors:    0,
			checkTokens:   true,
		},
		{
			name:          "valid function call",
			expression:    "SUM(10, 20)",
			wantTokens:    8, // SUM, (, 10, ,, whitespace, 20, ), EOF
			wantErrors:    0,
			checkTokens:   true,
		},
		{
			name:          "invalid expression",
			expression:    "1 + + 2",
			wantTokens:    8, // エラーがあってもトークンは返される
			wantErrors:    1,
			checkTokens:   true,
		},
		{
			name:          "empty expression",
			expression:    "",
			wantTokens:    0,
			wantErrors:    1,
			checkTokens:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := []js.Value{js.ValueOf(tt.expression)}
			result := analyze(js.Value{}, args)
			
			resultMap := result.(js.Value)
			tokens := resultMap.Get("tokens")
			errors := resultMap.Get("errors")
			
			tokenLength := tokens.Length()
			errorLength := errors.Length()
			
			if tt.checkTokens && tokenLength != tt.wantTokens {
				t.Errorf("analyze(%q) returned %d tokens, want %d", tt.expression, tokenLength, tt.wantTokens)
			}
			
			if errorLength != tt.wantErrors {
				t.Errorf("analyze(%q) returned %d errors, want %d", tt.expression, errorLength, tt.wantErrors)
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

func TestInvalidArguments(t *testing.T) {
	t.Run("validate with no arguments", func(t *testing.T) {
		args := []js.Value{}
		result := validate(js.Value{}, args)
		
		if got := result.(js.Value).Bool(); got != false {
			t.Errorf("validate() with no args = %v, want false", got)
		}
	})
	
	t.Run("analyze with no arguments", func(t *testing.T) {
		args := []js.Value{}
		result := analyze(js.Value{}, args)
		
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