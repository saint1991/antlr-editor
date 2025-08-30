package app

import (
	"sort"

	"github.com/antlr4-go/antlr/v4"

	"antlr-editor/analyzer/core/infrastructure"
	"antlr-editor/analyzer/core/models"
	"antlr-editor/analyzer/gen/parser"
)

// TokenizeResult contains the complete tokenization result
type TokenizeResult struct {
	Tokens []models.TokenInfo `json:"tokens"` // List of tokens
	Errors []models.ErrorInfo `json:"errors"` // List of error information
}

// IsValid returns true if the expression has no errors
func (r *TokenizeResult) IsValid() bool {
	return len(r.Errors) == 0
}

func (r *TokenizeResult) AsMap() map[string]any {
	tokens := make([]any, len(r.Tokens))
	for i, token := range r.Tokens {
		tokens[i] = token.AsMap()
	}

	errors := make([]any, len(r.Errors))
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
func (a *Analyzer) Analyze(expression string) *TokenizeResult {

	errors := make([]models.ErrorInfo, 0)

	if expression == "" {
		// Empty expressions are invalid - add an error
		errors = append(errors, models.ErrorInfo{
			Message: "Empty expression",
			Line:    1,
			Column:  0,
			Start:   0,
			End:     0,
		})
		return &TokenizeResult{
			Tokens: []models.TokenInfo{},
			Errors: errors,
		}
	}

	// First, collect all tokens from the lexer (including whitespace)
	lexer := a.helper.CreateLexer(expression)
	tokens := a.collectTokens(expression, lexer)

	// Check for ERROR_CHAR tokens and add errors if they exist
	for _, token := range tokens {
		if token.Type == models.TokenError {
			// Check if an error already exists for this position
			hasError := false
			for _, err := range errors {
				if token.Start >= err.Start && token.Start < err.End {
					hasError = true
					break
				}
			}
			if !hasError {
				errors = append(errors, models.ErrorInfo{
					Message: "Invalid character sequence: " + token.Text,
					Line:    token.Line,
					Column:  token.Column,
					Start:   token.Start,
					End:     token.End,
				})
			}
		}
	}

	// Create parser context for parsing
	ctx := a.helper.CreateParser(expression)
	errorListener := infrastructure.NewCollectingErrorListener(&errors)
	a.helper.SetupErrorListeners(ctx, errorListener)

	// Parse the expression - start with the root rule
	tree := a.helper.ParseExpression(ctx)

	// Check if all tokens were consumed - if not, this is an error
	if !a.helper.IsAllTokensConsumed(ctx) {
		currentToken := ctx.Parser.GetCurrentToken()
		errors = append(errors, models.ErrorInfo{
			Message: "Unexpected tokens at end of expression",
			Line:    currentToken.GetLine(),
			Column:  currentToken.GetColumn(),
			Start:   currentToken.GetStart(),
			End:     currentToken.GetStop() + 1,
		})
	}

	// Perform semantic validation if syntax parsing succeeded
	if len(errors) == 0 {
		a.performSemanticValidation(tree, &errors)
	}

	// Note: Validity can be determined by checking if len(result.Errors) == 0

	// Mark tokens as invalid if they are in error regions
	a.markErrorTokens(&tokens, errors)

	return &TokenizeResult{
		Tokens: tokens,
		Errors: errors,
	}
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
func (a *Analyzer) collectTokens(expression string, lexer *parser.ExpressionLexer) []models.TokenInfo {
	tokens := make([]models.TokenInfo, 0)

	for {
		token := lexer.NextToken()
		if token.GetTokenType() == antlr.TokenEOF {
			// Add EOF token
			tokens = append(tokens, models.TokenInfo{
				Type:   models.TokenEOF,
				Text:   "",
				Start:  token.GetStart(),
				End:    token.GetStart(),
				Line:   token.GetLine(),
				Column: token.GetColumn(),
			})
			break
		}

		// Skip tokens from HIDDEN channel and ERROR_CHAR channel (they'll be handled separately)
		// channel(HIDDEN) = 1, channel(2) is for ERROR_CHAR
		if token.GetChannel() == antlr.LexerHidden || token.GetChannel() == 2 {
			continue
		}

		// Determine token type based on token type from lexer
		tokenType := a.getTokenType(token.GetTokenType())
		if tokenType == models.TokenColumnReference {
			leftBracket := models.TokenInfo{
				Type:   models.TokenLeftBracket,
				Text:   "[",
				Start:  token.GetStart(),
				End:    token.GetStart() + 1,
				Line:   token.GetLine(),
				Column: token.GetColumn(),
			}
			identifier := models.TokenInfo{
				Type:   models.TokenColumnReference,
				Text:   token.GetText()[1 : len(token.GetText())-1],
				Start:  token.GetStart() + 1,
				End:    token.GetStop(),
				Line:   token.GetLine(),
				Column: token.GetColumn() + 1,
			}
			rightBracket := models.TokenInfo{
				Type:   models.TokenRightBracket,
				Text:   "]",
				Start:  token.GetStop(),
				End:    token.GetStop() + 1,
				Line:   token.GetLine(),
				Column: token.GetColumn() + 1 + len(identifier.Text),
			}
			tokens = append(tokens, leftBracket, identifier, rightBracket)
		} else {
			tokenInfo := models.TokenInfo{
				Type:   tokenType,
				Text:   token.GetText(),
				Start:  token.GetStart(),
				End:    token.GetStop() + 1,
				Line:   token.GetLine(),
				Column: token.GetColumn(),
			}
			tokens = append(tokens, tokenInfo)
		}
	}

	// Also collect whitespace tokens from HIDDEN channel
	tokens = append(tokens, a.whitespaceTokens(expression)...)

	// Also collect ERROR_CHAR tokens from channel 2
	tokens = append(tokens, a.errorCharTokens(expression)...)

	// Sort tokens by position
	sort.SliceStable(tokens, func(i, j int) bool {
		return tokens[i].Start < tokens[j].Start
	})

	// Merge consecutive error tokens
	tokens = a.mergeConsecutiveErrorTokens(tokens)

	return tokens
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
	case parser.ExpressionLexerCOLUMN_REF:
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
	case parser.ExpressionLexerERROR_CHAR:
		return models.TokenError
	default:
		return models.TokenError
	}
}

func (a *Analyzer) collectAntlrTokens(expression string) []antlr.Token {
	// Create a new lexer to collect tokens from all channels
	lexer := a.helper.CreateLexer(expression)

	// Create token stream that includes HIDDEN channel
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	stream.Fill()

	return stream.GetAllTokens()
}

// addWhitespaceTokens adds whitespace tokens from HIDDEN channel
func (a *Analyzer) whitespaceTokens(expression string) []models.TokenInfo {
	whiteSpaceTokens := make([]models.TokenInfo, 0)

	for _, token := range a.collectAntlrTokens(expression) {
		// Check if this is a WS token in HIDDEN channel
		if token.GetChannel() == antlr.LexerHidden {
			whiteSpaceTokens = append(whiteSpaceTokens, models.TokenInfo{
				Type:   models.TokenWhitespace,
				Text:   token.GetText(),
				Start:  token.GetStart(),
				End:    token.GetStop() + 1,
				Line:   token.GetLine(),
				Column: token.GetColumn(),
			})
		}
	}
	return whiteSpaceTokens
}

// errorCharTokens collects ERROR_CHAR tokens from channel 2
func (a *Analyzer) errorCharTokens(expression string) []models.TokenInfo {
	errorTokens := make([]models.TokenInfo, 0)

	for _, token := range a.collectAntlrTokens(expression) {
		// Check if this is an ERROR_CHAR token in channel 2
		if token.GetChannel() == 2 {
			errorTokens = append(errorTokens, models.TokenInfo{
				Type:   models.TokenError,
				Text:   token.GetText(),
				Start:  token.GetStart(),
				End:    token.GetStop() + 1,
				Line:   token.GetLine(),
				Column: token.GetColumn(),
			})
		}
	}
	return errorTokens
}

// mergeConsecutiveErrorTokens merges consecutive error tokens into single tokens
func (a *Analyzer) mergeConsecutiveErrorTokens(tokens []models.TokenInfo) []models.TokenInfo {
	if len(tokens) == 0 {
		return tokens
	}

	merged := make([]models.TokenInfo, 0, len(tokens))
	var currentError *models.TokenInfo

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token.Type == models.TokenError {
			if currentError == nil {
				// Start a new error token
				errorCopy := token
				currentError = &errorCopy
			} else if currentError.End == token.Start {
				// Consecutive error token, merge it
				currentError.Text += token.Text
				currentError.End = token.End
			} else {
				// Non-consecutive error token, save current and start new
				merged = append(merged, *currentError)
				errorCopy := token
				currentError = &errorCopy
			}
		} else {
			// Non-error token
			if currentError != nil {
				merged = append(merged, *currentError)
				currentError = nil
			}
			merged = append(merged, token)
		}
	}

	// Don't forget the last error token if there is one
	if currentError != nil {
		merged = append(merged, *currentError)
	}

	return merged
}

// markErrorTokens marks tokens as invalid if they overlap with error regions
func (a *Analyzer) markErrorTokens(tokens *[]models.TokenInfo, errors []models.ErrorInfo) {
	for i := range *tokens {
		token := &(*tokens)[i]
		for _, err := range errors {
			// Check if token overlaps with error region
			if token.Start < err.End && token.End > err.Start {
				token.Type = models.TokenError
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
	if ctx.COLUMN_REF() != nil && ctx.COLUMN_REF().GetText() != "" {
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
