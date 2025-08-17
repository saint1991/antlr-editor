package app

import (
	"antlr-editor/analyzer/core/app/formatter"
)

type App struct {
	analyzer  *Analyzer
	formatter *Formatter
}

func NewApp() *App {
	return &App{
		analyzer:  newAnalyzer(),
		formatter: newFormatter(),
	}
}

func (app *App) Validate(expression string) bool {
	return app.analyzer.Validate(expression)
}

func (app *App) Analyze(expression string) *AnalysisResult {
	return app.analyzer.Analyze(expression)
}

// Format formats the given expression string using default formatting options
func (app *App) Format(expression string) string {
	return app.formatter.Format(expression)
}

// FormatWithOptions formats the given expression string using specified formatting options
func (app *App) FormatWithOptions(expression string, options *formatter.FormatOptions) string {
	return NewFormatterWithOptions(options).Format(expression)
}
