package infrastructure

import (
	"github.com/antlr4-go/antlr/v4"

	"antlr-editor/parser/core/models"
)

// BaseErrorListener provides common error handling functionality
type BaseErrorListener struct {
	*antlr.DefaultErrorListener
}

// ExtractErrorInfo extracts error information from syntax error parameters
func (b *BaseErrorListener) ExtractErrorInfo(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) models.ErrorInfo {
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

	return errorInfo
}

// SimpleErrorListener tracks whether any errors occurred
type SimpleErrorListener struct {
	BaseErrorListener
	HasError bool
}

// NewSimpleErrorListener creates a new simple error listener
func NewSimpleErrorListener() *SimpleErrorListener {
	return &SimpleErrorListener{
		BaseErrorListener: BaseErrorListener{DefaultErrorListener: &antlr.DefaultErrorListener{}},
		HasError:          false,
	}
}

// SyntaxError is called when a syntax error is encountered
func (s *SimpleErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	s.HasError = true
}

// CollectingErrorListener collects all error information
type CollectingErrorListener struct {
	BaseErrorListener
	Errors *[]models.ErrorInfo
}

// NewCollectingErrorListener creates a new collecting error listener
func NewCollectingErrorListener(errors *[]models.ErrorInfo) *CollectingErrorListener {
	return &CollectingErrorListener{
		BaseErrorListener: BaseErrorListener{DefaultErrorListener: &antlr.DefaultErrorListener{}},
		Errors:            errors,
	}
}

// SyntaxError is called when a syntax error is encountered
func (c *CollectingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	errorInfo := c.ExtractErrorInfo(recognizer, offendingSymbol, line, column, msg, e)
	*c.Errors = append(*c.Errors, errorInfo)
}
