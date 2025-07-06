package app

import (
	"github.com/antlr4-go/antlr/v4"

	"antlr-editor/parser/core/infrastructure"
	"antlr-editor/parser/core/models"
	"antlr-editor/parser/gen/parser"
)

// AnalysisResult contains the complete analysis result
type AnalysisResult struct {
	Tokens []models.TokenInfo `json:"tokens"` // List of tokens
	Errors []models.ErrorInfo `json:"errors"` // List of error information
}

// IsValid returns true if the expression has no errors
func (r *AnalysisResult) IsValid() bool {
	return len(r.Errors) == 0
}

func (r *AnalysisResult) AsMap() map[string]any {
	tokens := make([]map[string]any, len(r.Tokens))
	for i, token := range r.Tokens {
		tokens[i] = token.AsMap()
	}

	errors := make([]map[string]any, len(r.Errors))
	for i, err := range r.Errors {
		errors[i] = err.AsMap()
	}
	return map[string]any{
		"tokens": tokens,
		"errors": errors,
	}
}

// Analyzer provides expression syntax analysis functionality
type Analyzer struct {
	helper *infrastructure.ParserHelper
}

// newAnalyzer creates a new analyzer instance
func newAnalyzer() *Analyzer {
	return &Analyzer{
		helper: infrastructure.NewParserHelper(),
	}
}

// Analyze performs detailed token analysis of the given expression strin
func (a *Analyzer) Analyze(expression string) *AnalysisResult {
	result := &AnalysisResult{
		Tokens: make([]models.TokenInfo, 0),
		Errors: make([]models.ErrorInfo, 0),
	}

	if expression == "" {
		// Empty expressions are invalid - add an error
		result.Errors = append(result.Errors, models.ErrorInfo{
			Message: "Empty expression",
			Line:    1,
			Column:  0,
			Start:   0,
			End:     0,
		})
		return result
	}

	// Create parser context
	ctx := a.helper.CreateParser(expression)

	// Add custom error listener to collect detailed error information
	errorListener := infrastructure.NewCollectingErrorListener(&result.Errors)
	a.helper.SetupErrorListeners(ctx, errorListener)

	// First, collect all tokens from the lexer (including whitespace)
	lexer := a.helper.CreateLexerForTokenCollection(expression)
	a.collectTokens(expression, lexer, &result.Tokens)

	// Create a new parser context for parsing
	ctx = a.helper.CreateParser(expression)
	errorListener = infrastructure.NewCollectingErrorListener(&result.Errors)
	a.helper.SetupErrorListeners(ctx, errorListener)

	// Parse the expression - start with the root rule
	tree := a.helper.ParseExpression(ctx)

	// Check if all tokens were consumed - if not, this is an error
	if !a.helper.IsAllTokensConsumed(ctx) {
		result.Errors = append(result.Errors, models.ErrorInfo{
			Message: "Unexpected tokens at end of expression",
			Line:    1,
			Column:  0,
			Start:   0,
			End:     len(expression),
		})
	}

	// Perform semantic validation if syntax parsing succeeded
	if len(result.Errors) == 0 {
		a.performSemanticValidation(tree, &result.Errors)
	}

	// Note: Validity can be determined by checking if len(result.Errors) == 0

	// Mark tokens as invalid if they are in error regions
	a.markErrorTokens(&result.Tokens, result.Errors)

	return result
}

// Validate checks if the given expression string has valid syntax
// Returns true if the expression is syntactically and semantically valid, false otherwise
func (a *Analyzer) Validate(expression string) bool {
	if expression == "" {
		return false
	}
	result := a.Analyze(expression)
	return result.IsValid()
}

// collectTokens extracts all tokens from the lexer including whitespace
func (a *Analyzer) collectTokens(expression string, lexer *parser.ExpressionLexer, tokens *[]models.TokenInfo) {

	for {
		token := lexer.NextToken()
		if token.GetTokenType() == antlr.TokenEOF {
			// Add EOF token
			*tokens = append(*tokens, models.TokenInfo{
				Type:    models.TokenEOF,
				Text:    "",
				Start:   token.GetStart(),
				End:     token.GetStart(),
				Line:    token.GetLine(),
				Column:  token.GetColumn(),
				IsValid: true,
			})
			break
		}

		tokenInfo := models.TokenInfo{
			Text:    token.GetText(),
			Start:   token.GetStart(),
			End:     token.GetStop() + 1,
			Line:    token.GetLine(),
			Column:  token.GetColumn(),
			IsValid: true,
		}

		// Determine token type based on token type from lexer
		tokenInfo.Type = a.getTokenType(token.GetTokenType())

		*tokens = append(*tokens, tokenInfo)
	}

	// Also collect whitespace tokens by analyzing the original string
	a.addWhitespaceTokens(expression, tokens)
}

// getTokenType maps ANTLR token types to our TokenType enum
func (a *Analyzer) getTokenType(antlrTokenType int) models.TokenType {
	switch antlrTokenType {
	case parser.ExpressionLexerSTRING_LITERAL:
		return models.TokenString
	case parser.ExpressionLexerINTEGER_LITERAL:
		return models.TokenInteger
	case parser.ExpressionLexerFLOAT_LITERAL:
		return models.TokenFloat
	case parser.ExpressionLexerBOOLEAN_LITERAL:
		return models.TokenBoolean
	case parser.ExpressionLexerIDENTIFIER:
		return models.TokenColumnReference
	case parser.ExpressionLexerFUNCTION_NAME:
		return models.TokenFunction
	case parser.ExpressionLexerADD, parser.ExpressionLexerSUB, parser.ExpressionLexerMUL, parser.ExpressionLexerDIV, parser.ExpressionLexerPOW,
		parser.ExpressionLexerLT, parser.ExpressionLexerLE, parser.ExpressionLexerGT, parser.ExpressionLexerGE,
		parser.ExpressionLexerEQ, parser.ExpressionLexerNEQ, parser.ExpressionLexerAND, parser.ExpressionLexerOR:
		return models.TokenOperator
	case parser.ExpressionLexerLPAREN:
		return models.TokenLeftParen
	case parser.ExpressionLexerRPAREN:
		return models.TokenRightParen
	case parser.ExpressionLexerLBRACKET:
		return models.TokenLeftBracket
	case parser.ExpressionLexerRBRACKET:
		return models.TokenRightBracket
	case parser.ExpressionLexerCOMMA:
		return models.TokenComma
	default:
		return models.TokenError
	}
}

// addWhitespaceTokens adds whitespace tokens that were skipped by the lexer
func (a *Analyzer) addWhitespaceTokens(expression string, tokens *[]models.TokenInfo) {
	// Sort existing tokens by start position
	existingTokens := *tokens

	// Create a map of covered positions
	covered := make(map[int]bool)
	for _, token := range existingTokens {
		for i := token.Start; i < token.End; i++ {
			covered[i] = true
		}
	}

	// Find whitespace gaps
	line := 1
	column := 0

	for i, char := range expression {
		if !covered[i] {
			// This position is not covered by any token, check if it's whitespace
			if char == ' ' || char == '\t' || char == '\r' || char == '\n' {
				// Find the end of this whitespace sequence
				start := i
				end := i
				wsText := ""

				for j := i; j < len(expression) && !covered[j]; j++ {
					if expression[j] == ' ' || expression[j] == '\t' || expression[j] == '\r' || expression[j] == '\n' {
						wsText += string(expression[j])
						end = j + 1
					} else {
						break
					}
				}

				if end > start {
					wsToken := models.TokenInfo{
						Type:    models.TokenWhitespace,
						Text:    wsText,
						Start:   start,
						End:     end,
						Line:    line,
						Column:  column,
						IsValid: true,
					}

					// Insert in correct position
					inserted := false
					for idx, existing := range *tokens {
						if existing.Start > start {
							// Insert before this token
							*tokens = append((*tokens)[:idx], append([]models.TokenInfo{wsToken}, (*tokens)[idx:]...)...)
							inserted = true
							break
						}
					}
					if !inserted {
						*tokens = append(*tokens, wsToken)
					}

					// Mark these positions as covered
					for k := start; k < end; k++ {
						covered[k] = true
					}
				}
			}
		}

		// Update line and column tracking
		if char == '\n' {
			line++
			column = 0
		} else {
			column++
		}
	}

	// Sort tokens by start position to ensure proper order
	a.sortTokens(tokens)
}

// sortTokens sorts tokens by their start position
func (a *Analyzer) sortTokens(tokens *[]models.TokenInfo) {
	// Simple bubble sort for small token arrays
	for i := 0; i < len(*tokens)-1; i++ {
		for j := 0; j < len(*tokens)-i-1; j++ {
			if (*tokens)[j].Start > (*tokens)[j+1].Start {
				(*tokens)[j], (*tokens)[j+1] = (*tokens)[j+1], (*tokens)[j]
			}
		}
	}
}

// markErrorTokens marks tokens as invalid if they overlap with error regions
func (a *Analyzer) markErrorTokens(tokens *[]models.TokenInfo, errors []models.ErrorInfo) {
	for i := range *tokens {
		token := &(*tokens)[i]
		for _, err := range errors {
			// Check if token overlaps with error region
			if token.Start < err.End && token.End > err.Start {
				token.IsValid = false
				// Preserve original type but mark as invalid
				// Could also change type to TokenError if preferred
			}
		}
	}
}

// performSemanticValidation performs semantic validation on the parse tree
func (a *Analyzer) performSemanticValidation(tree parser.IExpressionContext, errors *[]models.ErrorInfo) {
	if tree == nil {
		*errors = append(*errors, models.ErrorInfo{
			Message: "Parse tree is nil",
			Line:    1,
			Column:  0,
			Start:   0,
			End:     1,
		})
		return
	}

	// Create validation visitor to check semantic validity
	visitor := &validationVisitor{
		BaseExpressionVisitor: &parser.BaseExpressionVisitor{},
	}

	// Visit the parse tree
	result := visitor.Visit(tree)

	// Check validation result
	if result == nil {
		*errors = append(*errors, models.ErrorInfo{
			Message: "Semantic validation failed",
			Line:    1,
			Column:  0,
			Start:   0,
			End:     1,
		})
		return
	}

	if valid, ok := result.(bool); !ok || !valid {
		*errors = append(*errors, models.ErrorInfo{
			Message: "Semantic validation failed",
			Line:    1,
			Column:  0,
			Start:   0,
			End:     1,
		})
	}
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

func (v *validationVisitor) VisitUnaryMinusExpr(ctx *parser.UnaryMinusExprContext) any {
	// Validate the inner expression
	return v.Visit(ctx.Expression())
}
