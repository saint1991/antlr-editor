package core

import (
	"github.com/antlr4-go/antlr/v4"

	"antlr-editor/parser/gen/parser"
)

// Validator provides expression syntax validation functionality
type Validator struct{}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	return &Validator{}
}

// Validate checks if the given expression string has valid syntax
// Returns true if the expression is syntactically valid, false otherwise
func (v *Validator) Validate(expression string) bool {
	if expression == "" {
		return false
	}

	// Create input stream from expression string
	input := antlr.NewInputStream(expression)

	// Create lexer
	lexer := parser.NewExpressionLexer(input)

	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Create parser
	p := parser.NewExpressionParser(stream)

	// Remove default error listeners to prevent console output
	p.RemoveErrorListeners()
	lexer.RemoveErrorListeners()

	// Add custom error listener to track if parsing errors occurred
	errorListener := &validationErrorListener{}
	p.AddErrorListener(errorListener)
	lexer.AddErrorListener(errorListener)

	// Parse the expression - start with the root rule
	tree := p.Expression()

	// If parsing failed, return false immediately
	if errorListener.hasError {
		return false
	}

	// Create validation visitor to check semantic validity
	visitor := &validationVisitor{
		BaseExpressionVisitor: &parser.BaseExpressionVisitor{},
	}

	// Visit the parse tree
	result := visitor.Visit(tree)

	// Return validation result
	if result != nil {
		if valid, ok := result.(bool); ok {
			return valid
		}
	}

	return false
}

// Custom error listener to track parsing errors
type validationErrorListener struct {
	*antlr.DefaultErrorListener
	hasError bool
}

func (d *validationErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	d.hasError = true
}

// Custom visitor for semantic validation
type validationVisitor struct {
	*parser.BaseExpressionVisitor
}

func (v *validationVisitor) Visit(tree antlr.ParseTree) any {
	if tree == nil {
		return false
	}
	return tree.Accept(v)
}

func (v *validationVisitor) VisitAndExpr(ctx *parser.AndExprContext) any {
	// Validate both operands
	left := v.Visit(ctx.Expression(0))
	right := v.Visit(ctx.Expression(1))

	if left == nil || right == nil {
		return false
	}

	leftValid, leftOk := left.(bool)
	rightValid, rightOk := right.(bool)

	return leftOk && rightOk && leftValid && rightValid
}

func (v *validationVisitor) VisitOrExpr(ctx *parser.OrExprContext) any {
	// Validate both operands
	left := v.Visit(ctx.Expression(0))
	right := v.Visit(ctx.Expression(1))

	if left == nil || right == nil {
		return false
	}

	leftValid, leftOk := left.(bool)
	rightValid, rightOk := right.(bool)

	return leftOk && rightOk && leftValid && rightValid
}

func (v *validationVisitor) VisitComparisonExpr(ctx *parser.ComparisonExprContext) any {
	// Validate both operands
	left := v.Visit(ctx.Expression(0))
	right := v.Visit(ctx.Expression(1))

	if left == nil || right == nil {
		return false
	}

	leftValid, leftOk := left.(bool)
	rightValid, rightOk := right.(bool)

	return leftOk && rightOk && leftValid && rightValid
}

func (v *validationVisitor) VisitAddSubExpr(ctx *parser.AddSubExprContext) any {
	// Validate both operands
	left := v.Visit(ctx.Expression(0))
	right := v.Visit(ctx.Expression(1))

	if left == nil || right == nil {
		return false
	}

	leftValid, leftOk := left.(bool)
	rightValid, rightOk := right.(bool)

	return leftOk && rightOk && leftValid && rightValid
}

func (v *validationVisitor) VisitMulDivExpr(ctx *parser.MulDivExprContext) any {
	// Validate both operands
	left := v.Visit(ctx.Expression(0))
	right := v.Visit(ctx.Expression(1))

	if left == nil || right == nil {
		return false
	}

	leftValid, leftOk := left.(bool)
	rightValid, rightOk := right.(bool)

	return leftOk && rightOk && leftValid && rightValid
}

func (v *validationVisitor) VisitPowerExpr(ctx *parser.PowerExprContext) any {
	// Validate both operands
	left := v.Visit(ctx.Expression(0))
	right := v.Visit(ctx.Expression(1))

	if left == nil || right == nil {
		return false
	}

	leftValid, leftOk := left.(bool)
	rightValid, rightOk := right.(bool)

	return leftOk && rightOk && leftValid && rightValid
}

func (v *validationVisitor) VisitParenExpr(ctx *parser.ParenExprContext) any {
	// Validate the inner expression
	return v.Visit(ctx.Expression())
}

func (v *validationVisitor) VisitLiteralExpr(ctx *parser.LiteralExprContext) any {
	// Validate the literal
	return v.Visit(ctx.Literal())
}

func (v *validationVisitor) VisitColumnRefExpr(ctx *parser.ColumnRefExprContext) any {
	// Validate the column reference
	return v.Visit(ctx.ColumnReference())
}

func (v *validationVisitor) VisitFunctionCallExpr(ctx *parser.FunctionCallExprContext) any {
	// Validate the function call
	return v.Visit(ctx.FunctionCall())
}

func (v *validationVisitor) VisitLiteral(ctx *parser.LiteralContext) any {
	// Literals are always valid if they were parsed successfully
	return true
}

func (v *validationVisitor) VisitColumnReference(ctx *parser.ColumnReferenceContext) any {
	// Column references are valid if they have a non-empty identifier
	if ctx.IDENTIFIER() != nil && ctx.IDENTIFIER().GetText() != "" {
		return true
	}
	return false
}

func (v *validationVisitor) VisitFunctionCall(ctx *parser.FunctionCallContext) any {
	// Validate function name exists
	if ctx.FUNCTION_NAME() == nil || ctx.FUNCTION_NAME().GetText() == "" {
		return false
	}

	// Validate arguments if present
	if ctx.ArgumentList() != nil {
		return v.Visit(ctx.ArgumentList())
	}

	return true
}

func (v *validationVisitor) VisitArgumentList(ctx *parser.ArgumentListContext) any {
	// Validate all arguments
	for _, expr := range ctx.AllExpression() {
		result := v.Visit(expr)
		if result == nil {
			return false
		}
		if valid, ok := result.(bool); !ok || !valid {
			return false
		}
	}
	return true
}

func (v *validationVisitor) VisitIdentifierExpr(ctx *parser.IdentifierExprContext) any {
	// Identifier expressions are valid if they have a non-empty identifier
	if ctx.IDENTIFIER() != nil && ctx.IDENTIFIER().GetText() != "" {
		return true
	}
	return false
}

func (v *validationVisitor) VisitUnaryMinusExpr(ctx *parser.UnaryMinusExprContext) any {
	// Validate the inner expression
	return v.Visit(ctx.Expression())
}
