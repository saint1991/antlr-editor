package core

import (
	"github.com/antlr4-go/antlr/v4"

	"antlr-editor/parser/core/models"
	"antlr-editor/parser/gen/parser"
)

// AnalysisResult contains the complete analysis result
type AnalysisResult struct {
	Tokens []models.TokenInfo `json:"tokens"` // List of tokens
	Errors []models.ErrorInfo `json:"errors"` // List of error information
}

// Analyzer provides expression syntax analysis functionality
type Analyzer struct{}

// NewAnalyzer creates a new analyzer instance
func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

// Analyze performs detailed token analysis of the given expression strin
func (a *Analyzer) Analyze(expression string) *AnalysisResult {
	result := &AnalysisResult{
		Tokens: make([]models.TokenInfo, 0),
		Errors: make([]models.ErrorInfo, 0),
	}

	if expression == "" {
		return result
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

	// Add custom error listener to collect detailed error information
	errorListener := &analysisErrorListener{
		errors: &result.Errors,
	}
	p.AddErrorListener(errorListener)
	lexer.AddErrorListener(errorListener)

	// First, collect all tokens from the lexer (including whitespace)
	a.collectTokens(expression, lexer, &result.Tokens)

	// Reset lexer and token stream for parsing
	input = antlr.NewInputStream(expression)
	lexer = parser.NewExpressionLexer(input)
	stream = antlr.NewCommonTokenStream(lexer, 0)
	p = parser.NewExpressionParser(stream)

	// Remove default error listeners again
	p.RemoveErrorListeners()
	lexer.RemoveErrorListeners()

	// Add error listener again
	errorListener = &analysisErrorListener{
		errors: &result.Errors,
	}
	p.AddErrorListener(errorListener)
	lexer.AddErrorListener(errorListener)

	// Parse the expression - start with the root rule
	_ = p.Expression()

	// Note: Validity can be determined by checking if len(result.Errors) == 0

	// Mark tokens as invalid if they are in error regions
	a.markErrorTokens(&result.Tokens, result.Errors)

	return result
}

// collectTokens extracts all tokens from the lexer including whitespace
func (a *Analyzer) collectTokens(expression string, lexer *parser.ExpressionLexer, tokens *[]models.TokenInfo) {
	// Create a new input stream for token collection
	input := antlr.NewInputStream(expression)
	tokenLexer := parser.NewExpressionLexer(input)

	// Remove error listeners to prevent interference
	tokenLexer.RemoveErrorListeners()

	for {
		token := tokenLexer.NextToken()
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

// Custom error listener for detailed error collection
type analysisErrorListener struct {
	*antlr.DefaultErrorListener
	errors *[]models.ErrorInfo
}

func (d *analysisErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	errorInfo := models.ErrorInfo{
		Message: msg,
		Line:    line,
		Column:  column,
		Start:   column,
		End:     column + 1,
	}

	// Try to get more precise position information from the offending symbol
	if token, ok := offendingSymbol.(antlr.Token); ok {
		errorInfo.Start = token.GetStart()
		errorInfo.End = token.GetStop() + 1
	}

	*d.errors = append(*d.errors, errorInfo)
}
