package formatter

import (
	"strings"
)

type FormatterContext struct {
	builder           *strings.Builder
	indent            int
	column            int
	depth             int
	options           *FormatOptions
	inFunction        bool // Track if we're inside a function call
	functionNestLevel int  // Track function nesting level
}

func newFormatterContext(options *FormatOptions) *FormatterContext {
	return &FormatterContext{
		builder:           &strings.Builder{},
		indent:            0,
		column:            0,
		depth:             0,
		options:           options,
		inFunction:        false,
		functionNestLevel: 0,
	}
}

func (ctx *FormatterContext) enterExpression() {
	ctx.depth++
}

func (ctx *FormatterContext) exitExpression() {
	if ctx.depth > 0 {
		ctx.depth--
	}
}

// write writes a string to the output buffer
func (ctx *FormatterContext) write(s string) {
	ctx.builder.WriteString(s)
	ctx.column += len(s)
}

// writeIndent writes the current indentation
func (ctx *FormatterContext) writeIndent() {
	spaces := strings.Repeat(" ", ctx.indent*ctx.options.IndentSize)
	ctx.write(spaces)
}

// writeNewline writes a newline and applies indentation
func (ctx *FormatterContext) writeNewline() {
	ctx.builder.WriteString("\n")
	ctx.column = 0
	ctx.writeIndent()
}

// writeNewlineWithIndent writes a newline and increases indentation
func (ctx *FormatterContext) writeNewlineWithIndent() {
	ctx.builder.WriteString("\n")
	ctx.column = 0
	ctx.increaseIndent()
	ctx.writeIndent()
}

// decreaseIndent decreases the indentation level
func (ctx *FormatterContext) decreaseIndent() {
	if ctx.indent > 0 {
		ctx.indent--
	}
}

// increaseIndent increases the indentation level
func (ctx *FormatterContext) increaseIndent() {
	ctx.indent++
}

// enterFunction marks that we're entering a function call
func (ctx *FormatterContext) enterFunction() {
	ctx.inFunction = true
	ctx.functionNestLevel++
}

// exitFunction marks that we're exiting a function call
func (ctx *FormatterContext) exitFunction() {
	if ctx.functionNestLevel > 0 {
		ctx.functionNestLevel--
	}
	ctx.inFunction = ctx.functionNestLevel > 0
}

// isNestedFunction returns true if we're inside a nested function call
func (ctx *FormatterContext) isNestedFunction() bool {
	return ctx.functionNestLevel > 1
}

// writeSpaceAroundOperators writes a space if SpaceAroundOps is enabled
func (ctx *FormatterContext) writeSpaceAroundOperators() {
	if ctx.options.SpaceAroundOps {
		ctx.write(" ")
	}
}

// finalize finalizes the formatting and returns the result
func (ctx *FormatterContext) finalize() string {
	return ctx.builder.String()
}
