package models

// TokenType represents the type of a token for syntax highlighting
type TokenType string

const (
	// Literals
	TokenString  TokenType = "string"  // String literals
	TokenInteger TokenType = "integer" // Integer literals
	TokenFloat   TokenType = "float"   // Float literals
	TokenBoolean TokenType = "boolean" // Boolean literals

	// Identifiers
	TokenColumnReference TokenType = "columnReference" // Column names like [price]
	TokenFunction        TokenType = "function"        // Function names

	// Operators
	TokenOperator TokenType = "operator" // Arithmetic/comparison/logical operators

	// Delimiters
	TokenComma        TokenType = "comma"        // Commas
	TokenLeftParen    TokenType = "leftParen"    // Left parenthesis (
	TokenRightParen   TokenType = "rightParen"   // Right parenthesis )
	TokenLeftBracket  TokenType = "leftBracket"  // Left bracket [
	TokenRightBracket TokenType = "rightBracket" // Right bracket ]

	// Special
	TokenWhitespace TokenType = "whitespace" // Whitespace characters
	TokenError      TokenType = "error"      // Error tokens
	TokenEOF        TokenType = "eof"        // End of file
)

// TokenInfo contains detailed information about a token
type TokenInfo struct {
	Type   TokenType `json:"type"`   // Token type
	Text   string    `json:"text"`   // Token text
	Start  int       `json:"start"`  // Start position in string (0-based)
	End    int       `json:"end"`    // End position in string (0-based, exclusive)
	Line   int       `json:"line"`   // Line number (1-based)
	Column int       `json:"column"` // Column number (0-based)
}

func (t *TokenInfo) AsMap() map[string]any {
	return map[string]any{
		"type":   string(t.Type),
		"text":   t.Text,
		"start":  t.Start,
		"end":    t.End,
		"line":   t.Line,
		"column": t.Column,
	}
}
