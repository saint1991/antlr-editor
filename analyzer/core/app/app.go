package app

type App struct {
	analyzer *Analyzer
}

func NewApp() *App {
	return &App{
		analyzer: newAnalyzer(),
	}
}

func (app *App) Validate(expression string) bool {
	return app.analyzer.Validate(expression)
}

func (app *App) Analyze(expression string) *AnalysisResult {
	return app.analyzer.Analyze(expression)
}
