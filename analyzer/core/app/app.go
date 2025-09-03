package app

import (
	"antlr-editor/analyzer/core/app/formatter"
	"antlr-editor/analyzer/core/models"
)

type App struct {
	analyzer  *Analyzer
	formatter *Formatter
}

// NewApp creates a new App instance with analyzer and formatter components
func NewApp() *App {
	return &App{
		analyzer:  newAnalyzer(),
		formatter: newFormatter(),
	}
}

// ParseTree builds a hierarchical parse tree from the expression
func (app *App) ParseTree(expression string) *ParseTreeResult {
	return app.analyzer.ParseTree(expression)
}

// Lint performs comprehensive linting on the expression, checking for syntax errors, invalid tokens, and semantic issues
func (app *App) Lint(expression string) []models.ErrorInfo {
	return app.analyzer.Lint(expression)
}

// Tokenize performs detailed token analysis of the given expression string.
// Returns all tokens from all channels including whitespace and error tokens that don't match any lexer rules.
// The Errors field contains only parse errors, not lexical error tokens (which are included in Tokens).
func (app *App) Tokenize(expression string) *TokenizeResult {
	return app.analyzer.Tokenize(expression)
}

// Validate checks if the given expression string has valid syntax.
// Returns true if the expression is syntactically and semantically valid, false otherwise.
func (app *App) Validate(expression string) bool {
	return app.analyzer.Validate(expression)
}

// Format formats the given expression string using default formatting options
func (app *App) Format(expression string) string {
	return app.formatter.Format(expression)
}

// FormatWithOptions formats the given expression string using specified formatting options
func (app *App) FormatWithOptions(expression string, options *formatter.FormatOptions) string {
	return NewFormatterWithOptions(options).Format(expression)
}
