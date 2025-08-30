package formatter

import (
	"github.com/antlr4-go/antlr/v4"

	"antlr-editor/analyzer/gen/parser"
)

var (
	tokenType2Operator = map[int]string{
		parser.ExpressionLexerADD: "+",
		parser.ExpressionLexerSUB: "-",
		parser.ExpressionLexerMUL: "*",
		parser.ExpressionLexerDIV: "/",
		parser.ExpressionLexerLT:  "<",
		parser.ExpressionLexerLE:  "<=",
		parser.ExpressionLexerGT:  ">",
		parser.ExpressionLexerGE:  ">=",
		parser.ExpressionLexerEQ:  "==",
		parser.ExpressionLexerNEQ: "!=",
		parser.ExpressionLexerAND: "&&",
		parser.ExpressionLexerOR:  "||",
		parser.ExpressionLexerPOW: "^",
	}
)

// FormatVisitor implements the ExpressionVisitor interface for formatting
type FormatVisitor struct {
	*parser.BaseExpressionVisitor
	ctx *FormatterContext
}

func NewFormatterVisitor(options *FormatOptions) *FormatVisitor {
	if options == nil {
		options = DefaultFormatOptions()
	}
	return &FormatVisitor{
		BaseExpressionVisitor: &parser.BaseExpressionVisitor{},
		ctx:                   newFormatterContext(options),
	}
}

func (v *FormatVisitor) Finalize() string {
	return v.ctx.finalize()
}

// Visit visits a parse tree node
func (v *FormatVisitor) Visit(tree antlr.ParseTree) any {
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}

// VisitLiteralExpr formats a literal expression
func (v *FormatVisitor) VisitLiteralExpr(ctx *parser.LiteralExprContext) any {
	return v.Visit(ctx.Literal())
}

// VisitLiteral formats a literal value
func (v *FormatVisitor) VisitLiteral(ctx *parser.LiteralContext) any {
	// Get the literal text and write it as-is
	v.ctx.write(ctx.GetText())
	return nil
}

// VisitColumnRefExpr formats a column reference expression
func (v *FormatVisitor) VisitColumnRefExpr(ctx *parser.ColumnRefExprContext) any {
	return v.Visit(ctx.ColumnReference())
}

// VisitColumnReference formats a column reference
func (v *FormatVisitor) VisitColumnReference(ctx *parser.ColumnReferenceContext) any {
	if ctx.COLUMN_REF() != nil {
		v.ctx.write(ctx.COLUMN_REF().GetText())
	}
	return nil
}

// VisitParenExpr formats a parenthesized expression
func (v *FormatVisitor) VisitParenExpr(ctx *parser.ParenExprContext) any {
	v.ctx.write("(")
	v.Visit(ctx.Expression())
	v.ctx.write(")")
	return nil
}

// VisitUnaryMinusExpr formats a unary minus expression
func (v *FormatVisitor) VisitUnaryMinusExpr(ctx *parser.UnaryMinusExprContext) any {
	v.ctx.write("-")
	v.Visit(ctx.Expression())
	return nil
}

// VisitFunctionCallExpr formats a function call expression
func (v *FormatVisitor) VisitFunctionCallExpr(ctx *parser.FunctionCallExprContext) any {
	return v.Visit(ctx.FunctionCall())
}

// VisitFunctionCall formats a function call
func (v *FormatVisitor) VisitFunctionCall(ctx *parser.FunctionCallContext) any {
	v.ctx.enterFunction()
	defer v.ctx.exitFunction()

	// Function name
	functionName := ""
	if ctx.FUNCTION_NAME() != nil {
		functionName = ctx.FUNCTION_NAME().GetText()
		v.ctx.write(functionName)
	}

	v.ctx.write("(")

	// Check if we need multi-line format for arguments
	if ctx.ArgumentList() != nil {
		argList := ctx.ArgumentList()
		expressions := argList.AllExpression()

		// Estimate total length of arguments with proper spacing
		totalArgLength := 0
		for i, expr := range expressions {
			if i > 0 {
				totalArgLength += 2 // ", "
			}
			totalArgLength += len(expr.GetText())
		}

		// Determine if we should use multi-line format
		// Add the current column position to the total estimated length
		functionCallLength := 1 + totalArgLength                // ")" + args
		currentLineWouldBe := v.ctx.column + functionCallLength // indent + function name + "()" + args

		// Nested functions should not be multi-line if parent is already multi-line
		shouldBreakArgs := v.ctx.options.BreakLongExpressions &&
			currentLineWouldBe > v.ctx.options.MaxLineLength &&
			len(expressions) > 1

		if shouldBreakArgs {
			// Multi-line format
			v.ctx.writeNewlineWithIndent()
			v.visitArgumentListMultiLine(expressions)
			v.ctx.writeNewline()
		} else {
			// Single-line format
			v.Visit(argList)
		}
	}

	v.ctx.write(")")
	return nil
}

// VisitArgumentList formats a function argument list
func (v *FormatVisitor) VisitArgumentList(ctx *parser.ArgumentListContext) any {
	expressions := ctx.AllExpression()

	for i, expr := range expressions {
		if i > 0 {
			v.ctx.write(", ") // Always add space after comma in function arguments
		}
		v.Visit(expr)
	}

	return nil
}

// visitArgumentListMultiLine formats arguments in multi-line style
func (v *FormatVisitor) visitArgumentListMultiLine(expressions []parser.IExpressionContext) {
	for i, expr := range expressions {
		if i > 0 {
			v.ctx.write(",")
			v.ctx.writeNewline()
		}
		v.Visit(expr)
	}
	v.ctx.decreaseIndent()
}

type HasExpressionContext interface {
	antlr.ParserRuleContext
	Expression(i int) parser.IExpressionContext
}

func (v *FormatVisitor) visitBinaryExpression(ctx HasExpressionContext) any {
	v.ctx.enterExpression()
	defer v.ctx.exitExpression()

	left := ctx.Expression(0)
	right := ctx.Expression(1)

	operator := ""
	if operatorNode := ctx.GetChild(1); operatorNode != nil {
		if terminalNode, ok := operatorNode.(antlr.TerminalNode); ok {
			if op, exists := tokenType2Operator[terminalNode.GetSymbol().GetTokenType()]; exists {
				operator = op
			} else {
				operator = terminalNode.GetText() // Fallback to raw text if not found
			}
		}
	}
	if operator == "" {
		panic("Operator not found in binary expression context")
	}

	// Visit left operand
	v.Visit(left)

	// Estimate the length of the remaining expression
	rightText := right.GetText()
	estimatedLength := len(operator) + len(rightText) + 4 // operator + spaces + rough estimate

	// Check if we need to break the line before this operator
	if v.ctx.options.BreakLongExpressions &&
		((v.ctx.column + estimatedLength) > v.ctx.options.MaxLineLength) {
		v.ctx.writeNewlineWithIndent()
		v.ctx.write(operator)
		v.ctx.writeSpaceAroundOperators()
		v.ctx.decreaseIndent()
	} else {
		v.ctx.writeSpaceAroundOperators()
		v.ctx.write(operator)
		v.ctx.writeSpaceAroundOperators()
	}

	// Visit right operand
	v.Visit(right)

	return nil
}

// VisitAndExpr formats a logical AND expression
func (v *FormatVisitor) VisitAndExpr(ctx *parser.AndExprContext) any {
	return v.visitBinaryExpression(ctx)
}

// VisitOrExpr formats a logical OR expression
func (v *FormatVisitor) VisitOrExpr(ctx *parser.OrExprContext) any {
	return v.visitBinaryExpression(ctx)
}

// VisitComparisonExpr formats a comparison expression
func (v *FormatVisitor) VisitComparisonExpr(ctx *parser.ComparisonExprContext) any {
	return v.visitBinaryExpression(ctx)
}

// VisitAddSubExpr formats an addition/subtraction expression
func (v *FormatVisitor) VisitAddSubExpr(ctx *parser.AddSubExprContext) any {
	return v.visitBinaryExpression(ctx)
}

// VisitMulDivExpr formats a multiplication/division expression
func (v *FormatVisitor) VisitMulDivExpr(ctx *parser.MulDivExprContext) any {
	return v.visitBinaryExpression(ctx)
}

// VisitPowerExpr formats a power expression
func (v *FormatVisitor) VisitPowerExpr(ctx *parser.PowerExprContext) any {
	v.ctx.enterExpression()
	defer v.ctx.exitExpression()

	// Visit left operand
	v.Visit(ctx.Expression(0))

	v.ctx.writeSpaceAroundOperators()
	v.ctx.write("^")
	v.ctx.writeSpaceAroundOperators()

	// Visit right operand (power is right-associative)
	v.Visit(ctx.Expression(1))

	return nil
}
