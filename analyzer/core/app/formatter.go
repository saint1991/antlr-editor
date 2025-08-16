package app

import (
	"antlr-editor/analyzer/core/app/formatter"
	"antlr-editor/analyzer/core/infrastructure"
)

// Formatter provides expression formatting functionality
type Formatter struct {
	options *formatter.FormatOptions
	helper  *infrastructure.ParserHelper
}

// newFormatter creates a new formatter instance with default options
func newFormatter() *Formatter {
	return NewFormatterWithOptions(formatter.DefaultFormatOptions())
}

// NewFormatterWithOptions creates a new formatter instance with specified options
func NewFormatterWithOptions(options *formatter.FormatOptions) *Formatter {
	return &Formatter{
		options: options,
		helper:  infrastructure.NewParserHelper(),
	}
}

// Format formats the given expression string according to the formatting rules
func (f *Formatter) Format(expression string) string {
	if expression == "" {
		return ""
	}

	// First, validate the expression using the analyzer
	analyzer := newAnalyzer()
	if !analyzer.Validate(expression) {
		// Return original expression if it's invalid
		return expression
	}

	// Parse the expression
	ctx := f.helper.CreateParser(expression)
	tree := f.helper.ParseExpression(ctx)

	// Create and use the format visitor
	visitor := formatter.NewFormatterVisitor(f.options)

	// Visit the parse tree to generate formatted output
	visitor.Visit(tree)

	return visitor.Finalize()
}
