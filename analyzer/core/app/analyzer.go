package app

import (
	"sort"

	"github.com/antlr4-go/antlr/v4"

	"antlr-editor/analyzer/core/app/tree"
	"antlr-editor/analyzer/core/infrastructure"
	"antlr-editor/analyzer/core/models"
	"antlr-editor/analyzer/gen/parser"
)

// ParseTreeResult represents the result of parsing with tree structure
type ParseTreeResult struct {
	Tree   *models.ParseTreeNode `json:"tree"`   // Root of the parse tree
	Errors []models.ErrorInfo    `json:"errors"` // List of errors
}

// AsMap converts ParseTreeResult to a map for JSON serialization
func (r *ParseTreeResult) AsMap() map[string]any {
	var treeMap map[string]any
	if r.Tree != nil {
		treeMap = r.Tree.AsMap()
	}

	errors := make([]any, len(r.Errors))
	for i, err := range r.Errors {
		errors[i] = err.AsMap()
	}

	return map[string]any{
		"tree":   treeMap,
		"errors": errors,
	}
}

// TokenizeResult contains the complete tokenization result
type TokenizeResult struct {
	Tokens []models.TokenInfo `json:"tokens"` // List of tokens
	Errors []models.ErrorInfo `json:"errors"` // List of error information
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

// parseExpression parses the input expression string using ANTLR and returns the parse tree context and any parsing errors
func (a *Analyzer) parseExpression(expression string) (parser.IExpressionContext, []models.ErrorInfo) {
	// if expression is empty return nil no error
	if expression == "" {
		return nil, nil
	}

	errors := make([]models.ErrorInfo, 0)

	// Parse the expression
	ctx := a.helper.CreateParser(expression)
	errorListener := infrastructure.NewCollectingErrorListener(&errors)
	a.helper.SetupErrorListeners(ctx, errorListener)

	result := a.helper.ParseExpression(ctx)

	// Check if all tokens were consumed
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

	return result, errors
}

// collectAntlrTokens collects all tokens from the input expression including those in HIDDEN channel (whitespace, comments)
func (a *Analyzer) collectAntlrTokens(expression string) []antlr.Token {
	// Create a new lexer to collect tokens from all channels
	lexer := a.helper.CreateLexer(expression)

	// Create token stream that includes HIDDEN channel
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	stream.Fill()

	return stream.GetAllTokens()
}

// collectWhiteSpaceTokens adds whitespace tokens from HIDDEN channel
func (a *Analyzer) collectWhitesSpaceTokens(expression string) []models.TokenInfo {
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

// collectErrorTokens collects ERROR_CHAR tokens from channel 2
func (a *Analyzer) collectErrorTokens(expression string) []models.TokenInfo {
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
	tokens = append(tokens, a.collectWhitesSpaceTokens(expression)...)

	// Also collect ERROR_CHAR tokens from channel 2
	tokens = append(tokens, a.collectErrorTokens(expression)...)

	// Sort tokens by position
	sort.SliceStable(tokens, func(i, j int) bool {
		return tokens[i].Start < tokens[j].Start
	})

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

// performSemanticValidation performs semantic validation on the parse tree
func (a *Analyzer) performSemanticValidation(_ parser.IExpressionContext) []models.ErrorInfo {
	// TODO after implementing type system. returns nil so far.
	return nil
}

// ParseTree creates a hierarchical parse tree from the expression.
// Returns nil tree for empty expressions.
// Parse errors result in partial trees that exclude unparseable tokens.
func (a *Analyzer) ParseTree(expression string) *ParseTreeResult {
	expressionTree, errors := a.parseExpression(expression)

	// Build the parse tree structure (even with errors to show partial results)
	var parseTree *models.ParseTreeNode

	if expressionTree != nil {
		visitor := tree.NewParseTreeVisitor(expression)
		result := visitor.Visit(expressionTree)
		if node, ok := result.(*models.ParseTreeNode); ok {
			// Wrap the concrete expression type with a root Expression node
			parseTree = &models.ParseTreeNode{
				Type:     models.NodeTypeExpression,
				Text:     expression,
				Start:    0,
				End:      len(expression),
				Children: []models.ParseTreeNode{*node},
			}
		}
	}

	return &ParseTreeResult{
		Tree:   parseTree,
		Errors: errors,
	}
}

// Lint performs comprehensive linting on the expression, checking for syntax errors, invalid tokens, and semantic issues
func (a *Analyzer) Lint(expression string) []models.ErrorInfo {
	tree, errors := a.parseExpression(expression)

	errorTokens := a.collectErrorTokens(expression)
	for _, token := range errorTokens {
		errors = append(errors, models.ErrorInfo{
			Message: "Invalid character sequence: " + token.Text,
			Line:    token.Line,
			Column:  token.Column,
			Start:   token.Start,
			End:     token.End,
		})
	}

	errors = append(errors, a.performSemanticValidation(tree)...)
	return errors
}

// Tokenize performs detailed token analysis of the given expression string.
// Returns all tokens from all channels including whitespace and error tokens that don't match any lexer rules.
// The Errors field contains only parse errors, not lexical error tokens (which are included in Tokens).
func (a *Analyzer) Tokenize(expression string) *TokenizeResult {

	errors := make([]models.ErrorInfo, 0)

	if expression == "" {
		return &TokenizeResult{
			Tokens: []models.TokenInfo{},
			Errors: errors,
		}
	}

	// First, collect all tokens from the lexer (including whitespace)
	lexer := a.helper.CreateLexer(expression)
	tokens := a.collectTokens(expression, lexer)

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
	return len(a.Lint(expression)) == 0
}
