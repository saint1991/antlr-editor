package formatter

// FormatOptions contains configuration for the expression formatter
type FormatOptions struct {
	// IndentSize specifies the number of spaces per indent level
	IndentSize int

	// MaxLineLength specifies the maximum line length before breaking
	MaxLineLength int

	// SpaceAroundOps adds spaces around operators
	SpaceAroundOps bool

	// BreakLongExpressions automatically breaks long expressions
	BreakLongExpressions bool
}

// DefaultFormatOptions returns the default formatting options
func DefaultFormatOptions() *FormatOptions {
	return &FormatOptions{
		IndentSize:           2,
		MaxLineLength:        40,
		SpaceAroundOps:       true,
		BreakLongExpressions: true,
	}
}

// WithIndentSize returns a copy of options with the specified indent size
func (o *FormatOptions) WithIndentSize(size int) *FormatOptions {
	copy := *o
	copy.IndentSize = size
	return &copy
}

// WithMaxLineLength returns a copy of options with the specified max line length
func (o *FormatOptions) WithMaxLineLength(length int) *FormatOptions {
	copy := *o
	copy.MaxLineLength = length
	return &copy
}

// WithSpaceAroundOps returns a copy of options with the specified space around operators setting
func (o *FormatOptions) WithSpaceAroundOps(enabled bool) *FormatOptions {
	copy := *o
	copy.SpaceAroundOps = enabled
	return &copy
}

// WithBreakLongExpressions returns a copy of options with the specified break long expressions setting
func (o *FormatOptions) WithBreakLongExpressions(enabled bool) *FormatOptions {
	copy := *o
	copy.BreakLongExpressions = enabled
	return &copy
}
