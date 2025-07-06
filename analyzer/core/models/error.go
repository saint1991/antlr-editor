package models

// ErrorInfo contains detailed information about a parsing error
type ErrorInfo struct {
	Message string `json:"message"` // Error message
	Line    int    `json:"line"`    // Error line (1-based)
	Column  int    `json:"column"`  // Error column (0-based)
	Start   int    `json:"start"`   // Error start position
	End     int    `json:"end"`     // Error end position
}

func (e *ErrorInfo) AsMap() map[string]any {
	return map[string]any{
		"message": e.Message,
		"line":    e.Line,
		"column":  e.Column,
		"start":   e.Start,
		"end":     e.End,
	}
}
