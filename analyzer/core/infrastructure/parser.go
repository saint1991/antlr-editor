package infrastructure

import (
	"github.com/antlr4-go/antlr/v4"

	"antlr-editor/analyzer/gen/parser"
)

// ParserContext holds the components needed for parsing
type ParserContext struct {
	Input  *antlr.InputStream
	Lexer  *parser.ExpressionLexer
	Stream *antlr.CommonTokenStream
	Parser *parser.ExpressionParser
}

// ParserHelper provides common parsing utilities
type ParserHelper struct{}

// NewParserHelper creates a new parser helper instance
func NewParserHelper() *ParserHelper {
	return &ParserHelper{}
}

// CreateLexer creates a fresh lexer for token collection
// This is useful when you need to collect all tokens including whitespace
func (h *ParserHelper) CreateLexer(expression string) *parser.ExpressionLexer {
	input := antlr.NewInputStream(expression)
	lexer := parser.NewExpressionLexer(input)
	lexer.RemoveErrorListeners()
	return lexer
}

// CreateParser creates and initializes a parser context with the given expression
func (h *ParserHelper) CreateParser(expression string) *ParserContext {
	// Create input stream from expression string
	input := antlr.NewInputStream(expression)

	// Create lexer
	lexer := parser.NewExpressionLexer(input)

	// Create token stream
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Create parser
	p := parser.NewExpressionParser(stream)

	return &ParserContext{
		Input:  input,
		Lexer:  lexer,
		Stream: stream,
		Parser: p,
	}
}

// SetupErrorListeners removes default error listeners and adds custom ones
func (h *ParserHelper) SetupErrorListeners(ctx *ParserContext, errorListener antlr.ErrorListener) {
	// Remove default error listeners to prevent console output
	ctx.Parser.RemoveErrorListeners()
	ctx.Lexer.RemoveErrorListeners()

	// Add custom error listener
	if errorListener != nil {
		ctx.Parser.AddErrorListener(errorListener)
		ctx.Lexer.AddErrorListener(errorListener)
	}
}

// ParseExpression parses the expression and returns the parse tree
func (h *ParserHelper) ParseExpression(ctx *ParserContext) parser.IExpressionContext {
	return ctx.Parser.Expression()
}

// IsAllTokensConsumed checks if all tokens were consumed during parsing
func (h *ParserHelper) IsAllTokensConsumed(ctx *ParserContext) bool {
	token := ctx.Parser.GetCurrentToken()
	return token.GetTokenType() == antlr.TokenEOF
}
